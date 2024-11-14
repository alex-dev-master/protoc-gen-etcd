.DEFAULT_GOAL := help

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
.PHONY: debug

all: build 		## all commands

build:      		## build
	@printf "\033[32mBuilding...\033[0m\n"
	@go build .

debug: ## go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest
	@protoc -I=. -I=./example/vendor -I=/usr/local/include -I=./example/proto \
       --plugin=protoc-gen-debug=/Users/alexandr/go/bin/protoc-gen-debug \
       --debug_out="./debug:." \
       ./example/proto/*.proto