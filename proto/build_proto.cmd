protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. .\google\api\annotations.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. .\google\api\http.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=./google/api:. ./user.proto
