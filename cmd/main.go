package main

import (
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	pb "github.com/mathieux51/todogrpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

type todoServer struct{}

// log.Printf("Requested ID: %v", in.Id)

func (s *todoServer) GetTodoByID(ctx context.Context, req *pb.TodoByIDRequest) (*pb.TodoByIDResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}

	return &pb.TodoByIDResponse{Id: req.Id, Text: "Time to get Schwifty", Completed: false}, nil
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
		err = ioutil.WriteFile("img/clientsidestreaming.jpg", r.Data, 0644)
		if err != nil {
			return err
		}
	}

}

func (s *todoServer) DownloadImage(req *pb.DownloadImageRequest, stream pb.Todo_DownloadImageServer) error {
	data, err := ioutil.ReadFile("img/getschwifty.jpg")
	if err != nil {
		return err
	}
	if err = stream.Send(&pb.DownloadImageResponse{Data: data}); err != nil {
		return err
	}
	return nil
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		port := "50051"

		l, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Fatalf("failed to listen on port %v: %v", port, err)
		}

		s := grpc.NewServer()
		pb.RegisterTodoServer(s, &todoServer{})

		err = s.Serve(l)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("RegisterTodoServer: listening on port %v", port)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// Register gRPC server endpoint
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := pb.RegisterTodoHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
		if err != nil {
			log.Fatal(err)
		}

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		port := "8081"
		log.Printf("RegisterTodoHandlerFromEndpoint: listening on port %v", port)
		http.ListenAndServe(":"+port, mux)
	}()

	wg.Wait()
}
