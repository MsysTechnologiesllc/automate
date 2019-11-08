// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	"context"

	request "github.com/chef/automate/components/automate-gateway/api/infra_proxy/request"
	response "github.com/chef/automate/components/automate-gateway/api/infra_proxy/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the InfraProxyServer interface (at compile time)
var _ InfraProxyServer = &InfraProxyServerMock{}

// NewInfraProxyServerMock gives you a fresh instance of InfraProxyServerMock.
func NewInfraProxyServerMock() *InfraProxyServerMock {
	return &InfraProxyServerMock{validateRequests: true}
}

// NewInfraProxyServerMockWithoutValidation gives you a fresh instance of
// InfraProxyServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewInfraProxyServerMockWithoutValidation() *InfraProxyServerMock {
	return &InfraProxyServerMock{}
}

// InfraProxyServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type InfraProxyServerMock struct {
	validateRequests bool
	GetServersFunc   func(context.Context, *request.GetServers) (*response.GetServers, error)
	GetOrgsFunc      func(context.Context, *request.GetOrgs) (*response.GetOrgs, error)
}

func (m *InfraProxyServerMock) GetServers(ctx context.Context, req *request.GetServers) (*response.GetServers, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetServersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetServers' not implemented")
}

func (m *InfraProxyServerMock) GetOrgs(ctx context.Context, req *request.GetOrgs) (*response.GetOrgs, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetOrgsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetOrgs' not implemented")
}

// Reset resets all overridden functions
func (m *InfraProxyServerMock) Reset() {
	m.GetServersFunc = nil
	m.GetOrgsFunc = nil
}
