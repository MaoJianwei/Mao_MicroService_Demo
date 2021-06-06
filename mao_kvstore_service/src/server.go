package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "maojianwei.com/microservice.kvstore/grpc.maojianwei.com/api"
	"net"
)

const (
	addr = "[::]:9876"
)

var (
	kvstore = new(map[string]string)
)

type RealBigmaoServer struct {
	pb.UnimplementedBigmaoServer
}

func (s *RealBigmaoServer) GetRequest(ctx context.Context, data *pb.GetData) (*pb.GetResult, error) {

	value := (*kvstore)[data.Key.Key]
	fmt.Printf("Get: key %v, value %v\n", data.Key.Key, value)

	return &pb.GetResult{
		Result: &pb.Result{
			Err: 0,
			Msg: "ok",
		},
		Value: &pb.Value{
			Value: value,
		},
	}, nil
}

func (s *RealBigmaoServer) PutRequest(ctx context.Context, data *pb.PutData) (*pb.PutResult, error) {

	lastValue := (*kvstore)[data.Key.Key]
	(*kvstore)[data.Key.Key] = data.Value.Value
	fmt.Printf("Put: key %v, value %v\n", data.Key.Key, data.Value.Value)

	return &pb.PutResult{
		Result:    &pb.Result{
			Err: 0,
			Msg: "ok",
		},
		LastValue: &pb.Value{
			Value: lastValue,
		},
	}, nil
}

func main() {
	(*kvstore) = map[string]string{}
	(*kvstore)["beijing"] = "118.5"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf("listen fail, %v", err)
	}

	log.Printf("listen ok, %v", listener)

	server := grpc.NewServer()
	log.Printf("server ok, %v", server)

	pb.RegisterBigmaoServer(server, &RealBigmaoServer{})
	if err := server.Serve(listener); err != nil {
		log.Printf("Serve fail, %v", err)
	}
	log.Printf("serve ok, %v", err)
}




























