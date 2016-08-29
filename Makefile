tai64n.pb.go: tai64n.proto
	protoc -I=.:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:$(GOPATH)/src --gogoslick_out=plugins=grpc:. tai64n.proto

