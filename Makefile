.PHONY: all
all: protoc

.PHONY: protoc
protoc:
	protoc \
	-I./proto/ \
	--plugin=./node_modules/.bin/protoc-gen-ts_proto \
	--go_out=. \
	--go_opt=module=github.com/xremming/mousetoria \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/xremming/mousetoria \
	--ts_proto_out=src/gen \
	--ts_proto_opt=lowerCaseServiceMethods=true \
	--ts_proto_opt=outputServices=grpc-js \
	--ts_proto_opt=fileSuffix=.pb \
	proto/*.proto
