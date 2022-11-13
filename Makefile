NODE_MODULES_BIN := $(shell npm bin)

.PHONY: all
all: gameclient gameserver

.PHONY: start
start: all
	@echo "Starting the server..."
	$(NODE_MODULES_BIN)/electron --no-sandbox ./dist/main.js

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: gameclient
gameclient: protoc
	npm exec -- tsc -p gameclient/tsconfig.json

.PHONY: gameserver
gameserver: protoc
	go build -o dist/gameserver ./gameserver

.PHONY: protoc
protoc:
	protoc \
	-I./proto/ \
	--plugin=./node_modules/.bin/protoc-gen-ts_proto \
	--go_out=./gameserver/proto/ \
	--go_opt=paths=source_relative \
	--go-grpc_out=./gameserver/proto/ \
	--go-grpc_opt=paths=source_relative \
	--ts_proto_out=./gameclient/proto \
	--ts_proto_opt=lowerCaseServiceMethods=true \
	--ts_proto_opt=outputServices=grpc-js \
	--ts_proto_opt=fileSuffix=.pb \
	proto/*.proto
