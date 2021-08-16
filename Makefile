all: run
dir="./bin"
build:
	go build -o ./bin/server ./server/cmd/main.go

clean:
	go clean -i google.golang.org/grpc/...

rm:
	rm ./bin/*
rm_pb:
	rm ./proto/*.go
deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate ./proto

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/*.proto
test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

run:
	make protoc && make build && ./bin/server

testsubmodule:
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testrace \
	vet \
	vetdeps
