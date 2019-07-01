package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mathieux51/todogrpc/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	const ID = 1
	r, err := c.GetTodoByID(ctx, &pb.TodoByIDRequest{Id: ID})
	if err != nil {
		log.Fatalf("couldn't get todo (ID = %v): %v", ID, err)
	}
	log.Println(r)
}
