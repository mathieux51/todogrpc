package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/mathieux51/todogrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
)

type todoServer struct{}

// log.Printf("Requested ID: %v", in.Id)

func (s *todoServer) GetTodoByID(ctx context.Context, in *pb.TodoByIDRequest) (*pb.TodoByIDResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}

	return &pb.TodoByIDResponse{Id: 1, Text: "Time to get Schwifty", Completed: false}, nil
}

func (s *todoServer) UploadImage(stream pb.Todo_UploadImageServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("metadata from stream: %v", md)
	}
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.UploadImageResponse{})
			return nil
		}
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("img/test.jpg", r.Data, 0644)
		if err != nil {
			return err
		}
	}

}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", port, err)
	}

	s := grpc.NewServer()

	pb.RegisterTodoServer(s, &todoServer{})

	done := make(chan bool)
	go func() {
		err = s.Serve(l)
	}()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("listening on port %v", port)
	<-done
}
