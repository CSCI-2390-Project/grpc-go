module github.com/CSCI-2390-Project/grpc-go

go 1.11

require (
	github.com/CSCI-2390-Project/privacy-go v1.0.6
	github.com/cncf/udpa/go v0.0.0-20201001150855-7e6fe0510fb5
	github.com/envoyproxy/go-control-plane v0.9.7
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.2
	golang.org/x/net v0.0.0-20201029055024-942e2f445f3c
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418
	google.golang.org/genproto v0.0.0-20201029200359-8ce4113da6f7
	google.golang.org/protobuf v1.25.0
)

replace github.com/CSCI-2390-Project/grpc-go => ./
