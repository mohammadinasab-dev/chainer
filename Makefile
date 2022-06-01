GOPATH:=$(shell go env GOPATH)
.PHONY: proto


# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOPKG := $(.)


# A valid GOPATH is required to use the `go get` command.
# If $GOPATH is not specified, $HOME/go will be used by default
# GOPATH := $(if $(GOPATH),$(GOPATH),~/go)
go-build-api:
	@echo "  >  Building cmd binary..."
	GOBIN=$(GOBIN) go build -o $(GOBIN)/cmd ./cmd


proto:
	mkdir -p pkg/pbs

	protoc -I=./api/protos \
	--go_out=./pkg/pbs \
	--go-grpc_out=require_unimplemented_servers=false:./pkg/pbs \
	api/protos/*.proto
	

twcore-build:
	sudo mkdir -p libs
	docker rm -f twcorebuild || true
	docker create -ti --name twcorebuild trustwallet/wallet-core bash
	sudo docker cp twcorebuild:/wallet-core/build/libTrustWalletCore.a ./libs/libTrustWalletCore.a
	docker rm -f twcorebuild