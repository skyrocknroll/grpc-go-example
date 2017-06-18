package main

import (
	pb "github.com/skyrocknroll/grpc-go-example/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
)

type serverGreet struct {
}

func (s *serverGreet) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println(req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}
func main() {
	//const port int = 50051
	lis, err := net.Listen("tcp", "0.0.0.0:5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor), )
	pb.RegisterGreeterServer(s, &serverGreet{})
	// Register reflection service on gRPC serverGreet.
	reflection.Register(s)
	http.Handle("/metrics", prometheus.Handler())
	go func() { http.ListenAndServe("0.0.0.0:5002", nil) }()
	log.Println("After server")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
