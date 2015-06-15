tai64n.pb.go: tai64n.proto
	protoc --gofast_out=. -I=.:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:$(GOPATH)/src tai64n.proto
	sed -i ''  's/json:\"-\"/json:\"-\" codec:\"-\"/' tai64n.pb.go

