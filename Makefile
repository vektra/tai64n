tai64n.pb.go: tai64n.proto
	protoc --gogo_out=. -I=.:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:$(GOPATH)/src *.proto
	sed -i ''  's/json:\"-\"/json:\"-\" codec:\"-\"/' tai64n.pb.go

