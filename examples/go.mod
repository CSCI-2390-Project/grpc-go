module github.com/CSCI-2390-Project/grpc-go/examples

go 1.11

require (
	github.com/CSCI-2390-Project/grpc-go v1.36.0
	github.com/CSCI-2390-Project/privacy-go v1.0.6
	github.com/golang/protobuf v1.4.3
	golang.org/x/net v0.0.0-20201029221708-28c70e62bb1d // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201029200359-8ce4113da6f7
	google.golang.org/protobuf v1.25.0
)

replace github.com/CSCI-2390-Project/grpc-go => ../
