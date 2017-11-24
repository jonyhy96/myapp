package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/protoc"
	"log"
	"os"
)

const (
	address        = "localhost:4959"
	address2       = "localhost:4950"
	defaultcommand = "hello"
)

func get_echo() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyappserverClient(conn)
	name := defaultcommand
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Echo(context.Background(), &pb.EchoRequest{Command: name})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo :%v", r.Name)
}
func get_time() {
	conn, err := grpc.Dial(address2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyappserverClient(conn)
	gtcli := pb.GettimeRequest{Command: ""}
	//gtime := pb.GettimeReply{Name: ""}
	timestream, _ := c.Gettime(context.Background(), &gtcli)
	for true {
		gtime, err := timestream.Recv()
		if err != nil {
			fmt.Println("receive error")
		}
		if gtime.Name != "" {
			fmt.Println(gtime.Name)
		}
	}
}
func main() {
	fmt.Println("start get echo")
	go get_echo()
	fmt.Println("strt to get time")
	get_time()
}
