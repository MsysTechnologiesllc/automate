# Testing with mocked clients

Files in this directory are generated using a library called [mockgen]('https://github.com/golang/mock#running-mockgen'), which is part of [gomock]('https://github.com/grpc/grpc-go/blob/master/Documentation/gomock-example.md'). These generated files are used for testing the automate-gateway with mocked clients.

Gomock is used for testing client-side logic. All the GRPC clients in the gateway comes from the `ClientsFactory` [interface](https://github.com/chef/automate/blob/dc2e9abeaf09ad3577e922d0e8afec9f6713cc4e/components/automate-gateway/gateway/clients.go#L81). This interface has explicit definitions for all the micro-services that the gateway depends on, it is here where we dialing in to connect to those clients. Those clients contains the GRPC functions that we we need to mock so we can run our tests against the mocked version rather than the real version.

## How to generate a new mock

### Download mocking tools into your studio

Use the function defined in the common studio to download the tool:
```
[120][default:/hab/cache/src/go/src/github.com/chef/automate:0]# install_gomock
=> Installing Go Tool 'gomock'
github.com/chef/automate/vendor/golang.org/x/net/context
github.com/chef/automate/vendor/github.com/golang/mock/gomock
```

Ensure that it was downloaded correctly:
```
[123][default:/hab/cache/src/go/src/github.com/chef/automate:2]# mockgen
mockgen has two modes of operation: source and reflect.

Source mode generates mock interfaces from a source file.
It is enabled by using the -source flag. Other flags that
may be useful in this mode are -imports and -aux_files. ...
```

Most of these `mockgen` commands requires you to be located inside the Go Workspace, to enter the workspace use the alias `egw` or the command:
```
[1][default:/src:0]# enter_go_workspace
[2][default:/hab/cache/src/go/src/github.com/chef/automate:0]#
```

### Generate Mocked file

Mockgen works by generating a mock from the clients of the **protobuf file**.

The commands that generated `chef_ingest_mock.go` is:

```
[2][default:/hab/cache/src/go/src/github.com/chef/automate:0]# mockgen -destination components/automate-gateway/gateway_mocks/mock_ingest/chef_ingest_mock.go -imports empty=github.com/golang/protobuf/ptypes/empty -source api/interservice/ingest/chef.pb.go ChefIngesterClient
```

Breaking down the command:

    mockgen -destination components/automate-gateway/gateway_mocks/mock_ingest/chef_ingest_mock.go
        - the command plus the path where you want the mocked file
    -imports empty=github.com/golang/protobuf/ptypes/empty
        - in this case I want to ensure that the correct import path is used for the empty package
    -source api/interservice/ingest/chef.pb.go
        - the path to the protobuf file containing the client you wish to mock
    - ChefIngesterClient
        - the name of the client you wish to mock

Note that there's `make` targets for the existing mock files in `components/automate-gateway/gateway_mocks`.
Those can be regenerated by running `make` in that directory.

## How to Run the Tests

To run all of the unit tests in automate-gateway:

```
[0][default:/src:0]# go_component_unit automate-gateway
```
To run just a subset of the test, for example if you are working on an specific endpoint:

```
[1][default:/src:0]# enter_go_workspace
[2][default:/hab/cache/src/go/src/github.com/chef/automate:130]# go test -v -cover github.com/chef/automate/components/automate-gateway/gateway -run TestDataCollectorHandler
```

## Mocking Patterns

The function `newMockGatewayServer()` generates a mock `gateway.Server` instance, it receives a list of mocked services that will be injected to the `ClientsFactory` for its use when calling internal functions, handlers or even when passing the instance to a real object.

**Example:** Create an Automate Gateway where the IngestClient is mocked

```go
package gateway

import (
  "context"
  "errors"
  "testing"

  "github.com/golang/mock/gomock"
  gp "github.com/golang/protobuf/ptypes/empty"
  mock_ingest "github.com/chef/automate/components/automate-gateway/gateway_mocks/mock_ingest"
  ingestReq "github.com/chef/automate/components/ingest-service/events/chef"
)

func TestGatewayWithMockedIngestClient(t *testing.T) {
  // Create a mocked Chef Ingest client
  mockIngest := mock_ingest.NewMockChefIngesterClient(gomock.NewController(t))

  // Assert that we will call the ProcessChefAction() func and return an error
  mockIngest.EXPECT().ProcessChefAction(gomock.Any(), gomock.Any()).DoAndReturn(
    func(_ context.Context, _ *ingestReq.Action) (*gp.Empty, error) {
      return &gp.Empty{}, errors.New("Something happened")
    },
  )

  // Create a new gateway.Server instance with mocked Clients
  subject := newMockGatewayServer(t, mockIngest)

  // Call functions, inspec, pass it to other objects. Play with it!
}
```

You can mock almost any grpc client that the automate-gateway depends on, if you can't find the mock of the micro-service you are working on, you might have to generate them as showed in the previous sections.
