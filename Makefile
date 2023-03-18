default: install

tidy:
	@go mod tidy

install-tfplugindocs:
	@go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

generate: install-tfplugindocs
	export PATH="${PATH}:$(shell go env GOPATH)/bin"
	@go generate ./...

build: tidy
	@go build -o terraform-provider-stdlib

install:
	go install .

unit:
	go test -v

accept:
	TF_ACC=1 go test -v ./stdlib

test: unit accept
