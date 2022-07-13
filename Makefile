GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/category/category.proto
	
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
docker:
	docker build -t category:latest .
