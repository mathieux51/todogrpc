package main

import (
	"context"
	"io"
	"io/ioutil"
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

	// // GetTodoByID
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// const ID = 1
	// r, err := c.GetTodoByID(ctx, &pb.TodoByIDRequest{Id: ID})
	// if err != nil {
	// 	log.Fatalf("couldn't get todo (ID = %v): %v", ID, err)
	// }
	// log.Println(r)

	// Client streaming

	// Init context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Init client side streaming
	clientStream, err := c.UploadImage(ctx)
	if err != nil {
		log.Fatalf("%v.UploadImage(_) = _, %v", c, err)
	}

	// from image to byte slice
	// We could create chunks here
	data, err := ioutil.ReadFile("img/getschwifty.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// send data
	err = clientStream.Send(&pb.UploadImageRequest{Data: data})
	if err != nil {
		log.Fatal(err)
	}

	// close stream
	reply, err := clientStream.CloseAndRecv()
	if err == io.EOF {
		log.Println("all good")
	}
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", clientStream, err, nil)
	}
	log.Print(reply)

	// server side streaming
	serverStream, err := c.DownloadImage(ctx, &pb.DownloadImageRequest{})
	if err != nil {
		log.Fatalf("%v.DownloadImage(_) = _, %v", c, err)
	}
	for {
		res, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.DownloadImage(_) = _, %v", c, err)
		}
		err = ioutil.WriteFile("img/serversidestreaming.jpg", res.Data, 0644)
		if err != nil {
			log.Fatalf("%v.DownloadImage(_) = _, %v", c, err)
		}
	}
}
