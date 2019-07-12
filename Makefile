
.PHONY: protoc
protoc:
		protoc -I=proto \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:pb \
		--go_out=plugins=grpc:pb proto/todo.proto

.PHONY: start
start: 
		go run cmd/main.go