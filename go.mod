module github.com/CSCI-2390-Project/grpc-go

go 1.11

replace google.golang.org/grpc => github.com/CSCI-2390-Project/grpc-go v1.36.1-0.20201029210533-f22495f2c83c

replace google.golang.org/protobuf => github.com/CSCI-2390-Project/protobuf-go v1.25.1-0.20201029202626-b6c08e03c161

require (
	github.com/CSCI-2390-Project/privacy-go v1.0.8
	github.com/cncf/udpa/go v0.0.0-20201001150855-7e6fe0510fb5
	github.com/envoyproxy/go-control-plane v0.9.7
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.2
	golang.org/x/net v0.0.0-20201029055024-942e2f445f3c
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418
	google.golang.org/genproto v0.0.0-20201028140639-c77dae4b0522
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)
