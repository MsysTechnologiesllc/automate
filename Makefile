.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Go Modules vendoring is ridiculous. It doesn't vendor files that aren't
# part of packages. See https://github.com/golang/go/issues/26366
#
# Some of our tests and proto generators require files that go doesn't
# believe in vendoring so we copy them from the module cache to the vendor
# directory ourselves when we revendor.
grpcGatewayModPath=$(shell go list -f {{.Dir}} -m "github.com/grpc-ecosystem/grpc-gateway")
grpcGatewayVendorPath=./vendor/github.com/grpc-ecosystem/grpc-gateway/

.PHONY: revendor
revendor: ## revendor dependencies in vendor/, update .bldr.toml
	@go mod tidy
	@go mod vendor
	@go mod verify

	@mkdir -p $(grpcGatewayVendorPath)
	@cp -rf --no-preserve=mode $(grpcGatewayModPath)/* $(grpcGatewayVendorPath)
	go run tools/bldr-config-gen/main.go
