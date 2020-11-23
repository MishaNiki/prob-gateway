
gen-echo:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true:internal/echo/api/ \
		--swagger_out=allow_merge=true,merge_file_name=./api/echo:. \
		--go_out=plugins=grpc:internal/echo/api/ ./internal/echo/api/*.proto \


gen-hello:	
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true:internal/helloworld/api/ \
		--swagger_out=allow_merge=true,merge_file_name=./api/helloworld:. \
		--go_out=plugins=grpc:internal/helloworld/api/ ./internal/helloworld/api/*.proto

build:
	go build -o ./bin/server ./cmd/server/