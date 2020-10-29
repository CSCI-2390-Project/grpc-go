all: vet test testrace

build: deps
	go build github.com/CSCI-2390-Project/grpc-go/...

clean:
	go clean -i github.com/CSCI-2390-Project/grpc-go/...

deps:
	go get -d -v github.com/CSCI-2390-Project/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/CSCI-2390-Project/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/CSCI-2390-Project/grpc-go/...

testsubmodule: testdeps
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m github.com/CSCI-2390-Project/grpc-go/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m github.com/CSCI-2390-Project/grpc-go/security/authorization/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/CSCI-2390-Project/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/CSCI-2390-Project/grpc-go/...

testdeps:
	go get -d -v -t github.com/CSCI-2390-Project/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/CSCI-2390-Project/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/CSCI-2390-Project/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/CSCI-2390-Project/grpc-go/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
