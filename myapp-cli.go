package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func getecho(str string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := str
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("echo successed")
	return r.Message
}
func main() {
	var str string
	flag.StringVar(&str, "e", "", "write string to use.  defaults to empty.")
	flag.Parse()
	str = getecho(str)
	log.Printf("echo: %s", str)
}
