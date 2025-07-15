default: install

fmt:
	@go fmt ./...

tidy:
	@go mod tidy

install-tfplugindocs:
	@go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

generate: install-tfplugindocs
	PATH="${PATH}:$(shell go env GOPATH)/bin" go generate ./...

install-tfproviderlint:
	@go install github.com/bflad/tfproviderlint/cmd/tfproviderlint@latest

lint: install-tfproviderlint
	PATH="${PATH}:$(shell go env GOPATH)/bin" tfproviderlint -AT003=false -AT008=false ./...

build: tidy
	@go build -o terraform-provider-stdlib

install:
	@go install .

unit:
	@go test -v ./...

accept:
	TF_ACC=1 go test -v ./...

test: accept
