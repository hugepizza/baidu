PWD:=${shell pwd}
PROTOC=protoc
GO:=go
PROTO_DIR:=${GOPATH}/src/github.com/funxdata/baidu
RENAME_LIST:=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types

init:
	@which protoc-gen-gogofast || ${GO} get github.com/gogo/protobuf/protoc-gen-gogofast
	@which protoc-gen-gogofaster || ${GO} get github.com/gogo/protobuf/protoc-gen-gogofaster

BAIDU_SVC_PKG?=github.com/funxdata/baidu/service
BAIDU_SVC_DIR:=${GOPATH}/src/${BAIDU_SVC_PKG}
generate-baidu: init
	@${PROTOC} -I./service \
	 -I${PROTO_DIR} \
	 -I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
	 -I${GOPATH}/src/github.com/gogo/protobuf/ \
	 --grpc-gateway_out=logtostderr=true,${RENAME_LIST}:${BAIDU_SVC_DIR} \
	 --gogofaster_out=plugins=grpc,${RENAME_LIST}:${BAIDU_SVC_DIR} \
	 proto/*.proto
	@echo Generate successfully.	

