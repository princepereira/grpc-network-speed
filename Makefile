#SUBDIRS := $(wildcard */.)
SUBDIRS := cmd pkg
DOCKER_NAME := $(shell ls Dockerfile.*|cut -f2 -d".")
DOCKER_CLIENT := grpcclient
DOCKER_SERVER := grpcserver
DOCKER_LABEL := latest

export DOCS_PATH=$(PWD)/docs

# Static code analyzers
GOLINT := golint
GOFMT := gofmt
GOIMPORTS := goimports
STATIC_TOOLS := $(GOLINT) $(GOFMT) $(GOIMPORTS)
GO_FILES := `find . -name *.go`
CHANGE_DIR=cd
INTERSDPONGW_ROOT_PATH := ..
TEST_RESULT_DIR=test_results
ETCD=../sdpon-infra/db/kvstore/
KAFKA=../sdpon-infra/messaging/client/
PROTOS=protos/gw-protos
INTERSDPONGW_COMP_TEST_LOGS=/tmp/sdpon_intersdpongw_comp_test.log

exe:
	$(MAKE) -C cmd exe

execlient:
	@go build pkg/client/main.go

exeserver:
	@go build pkg/server/main.go

clean:
	$(MAKE) -C cmd clean

docker: exe
	@echo Building Docker $(DOCKER_NAME)....
	sudo docker build -t $(DOCKER_NAME):$(DOCKER_LABEL) -f Dockerfile.$(DOCKER_NAME) .

client: execlient
	@echo Building Docker client....
	sudo docker build -t $(DOCKER_CLIENT):$(DOCKER_LABEL) -f Dockerfile.grpc .

server: exeserver
	@echo Building Docker server....
	sudo docker build -t $(DOCKER_SERVER):$(DOCKER_LABEL) -f Dockerfile.grpc .

.PHONY: exe clean docker execlient exeserver client server doc start-test run-test end-test comp-test build-protos $(SUBDIRS) static $(STATIC_TOOLS)
