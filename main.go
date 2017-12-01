package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "myapp/protoc"
	_ "myapp/routers"
	"net"
	"time"
)

const (
	address  = "localhost:4959"
	address2 = "localhost:4950"
	port     = ":4959"
	port2    = ":4950"
)

// server is used to implement helloworld.GreeterServer.
type Myappserver struct {
	Ecreply []*pb.EchoReply
	Gtreply []*pb.GettimeReply
}

func (s *Myappserver) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	glog.Infoln("client said:" + in.Command)
	return &pb.EchoReply{Name: in.Command}, nil
}
func (s *Myappserver) Gettime(rect *pb.GettimeRequest, stream pb.Myappserver_GettimeServer) error {
	for {
		timer := time.NewTimer(time.Second * 60)
		time := time.Now()
		Timeecho := pb.GettimeReply{}
		const layout = "Jan 2, 2006 at 3:04pm (MST)"
		Timeecho.Name = time.Format(layout)
		err := stream.Send(&pb.GettimeReply{Name: time.Format(layout)})
		glog.Infoln("Send time:" + time.Format(layout))
		if err != nil {
			glog.Info(err)
			return err
		}
		<-timer.C
	}
}
func start_echo() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Info(err)
	}
	s := grpc.NewServer()
	pb.RegisterMyappserverServer(s, &Myappserver{})
	s.Serve(lis)
}
func start_time() {
	lis, err := net.Listen("tcp", port2)
	if err != nil {
		glog.Info(err)
	}
	s := grpc.NewServer()
	pb.RegisterMyappserverServer(s, &Myappserver{})
	s.Serve(lis)
}
func main() {
	flag.Parse()
	defer glog.Flush()
	go start_echo()
	glog.Info("start echo server")
	go start_time()
	glog.Info("start time server")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
