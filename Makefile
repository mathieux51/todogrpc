
.PHONY: protoc
protoc:
		protoc -I=proto --go_out=plugins=grpc:pb proto/*

.PHONY: start
start: 
		go run cmd/main.go