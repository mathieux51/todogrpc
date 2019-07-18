package main

import (
	"log"
	"net"
	"sync"

	pb "github.com/mathieux51/todogrpc/pb"
	todo "github.com/mathieux51/todogrpc/todo"
	"google.golang.org/grpc"
)

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
		pb.RegisterTodoServer(s, &todo.Server{})

		err = s.Serve(l)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("RegisterTodoServer: listening on port %v", port)
	}()

	wg.Add(1)
	go todo.StartTodoServer(&wg)

	wg.Wait()
}
