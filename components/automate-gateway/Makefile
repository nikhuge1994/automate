include ../../Makefile.common_go
.PHONY: docs run proto unit

BIN=automate-gateway
LINTERARGS:=--skip-files .*\.pb\.swagger\.go

run:
	@go run cmd/automate-gateway/automate-gateway.go serve --config config/config.toml

run-background:
	@go run cmd/automate-gateway/automate-gateway.go serve --config config/config.toml &> /tmp/gateway.log&

build:
	@go build --race -o $(BIN) github.com/chef/automate/components/automate-gateway/cmd/automate-gateway

proto:
	@echo "Please use the hab studio to compile protobuf files to ensure we are all on the same versions."
	@echo "The README has more info on how to use the studio, but to recompile the protobuf definitions please do the following.\n"
	@echo "Change directory to the root of this repo and run:\n"
	@echo "$ hab studio enter"
	@echo "# compile_go_protobuf_component automate-gateway"

unit:
	@go test -v -cover ./...
