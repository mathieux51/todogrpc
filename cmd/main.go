package main

import (
	"context"
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

	log.Printf("Listening on port %v", port)
	<-done
}
