module google.golang.org/grpc/examples

go 1.11

require (
	github.com/CSCI-2390-Project/privacy-go v1.99.1-0.20201126161349-6d371f6e2164
	github.com/golang/protobuf v1.4.3
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	google.golang.org/genproto v0.0.0-20201029200359-8ce4113da6f7
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/CSCI-2390-Project/grpc-go v1.36.1-0.20201029210533-f22495f2c83c

replace google.golang.org/protobuf => github.com/CSCI-2390-Project/protobuf-go v1.25.1-0.20201103164251-ba3a6e05f287
