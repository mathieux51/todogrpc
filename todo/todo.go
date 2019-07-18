package todo

import (
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log"
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

// Server ...
type Server struct{}

// GetTodoByID ...
func (s *Server) GetTodoByID(ctx context.Context, req *pb.TodoByIDRequest) (*pb.TodoByIDResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}

	return &pb.TodoByIDResponse{Id: req.Id, Text: "Time to get Schwifty", Completed: false}, nil
}

// UploadImage ...
func (s *Server) UploadImage(stream pb.Todo_UploadImageServer) error {
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

// DownloadImage ...
func (s *Server) DownloadImage(req *pb.DownloadImageRequest, stream pb.Todo_DownloadImageServer) error {
	data, err := ioutil.ReadFile("img/getschwifty.jpg")
	if err != nil {
		return err
	}
	if err = stream.Send(&pb.DownloadImageResponse{Data: data}); err != nil {
		return err
	}
	return nil
}

// StartTodoServer ...
func StartTodoServer(wg *sync.WaitGroup) {
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
}
