package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "maojianwei.com/microservice.kvstore/grpc.maojianwei.com/api"
	"time"

	"log"
)

const (
	serverAddr = "[3::1]:9876"
)

//os.Args[1], len(os.Args)
func main() {
	//WithInsecure, WithTimeout, WithBlock
	connect, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("connect fail, %v", err)
		return
	}
	defer connect.Close()
	log.Printf("client ok, %v, %v", connect, err)

	client := pb.NewBigmaoClient(connect)

	//var max int64 = 0
	//var min int64 = 922337203685477

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	keys := []string {"aaa", "bbb", "ccc"}
	values := []string {"1080", "7181", "5511"}
	updateValues := []string {"9999", "8888", "6666"}


	for i := 0; i < 3; i++ {
		t1 := time.Now()
		resp, err := client.GetRequest(ctx, &pb.GetData{
			Key:   &pb.Key{Key: keys[i]},
		})
		t2 := time.Now()
		diff := t2.Sub(t1).Nanoseconds()

		fmt.Printf("resp: %v\nerr: %v\ndiff: %v\n", resp, err, diff)
	}

	fmt.Println("=====================================================")

	for i := 0; i < 3; i++ {
		t1 := time.Now()
		resp, err := client.PutRequest(ctx, &pb.PutData{
			Key:   &pb.Key{Key:  keys[i]},
			Value: &pb.Value{Value:  values[i]},
		})
		t2 := time.Now()
		diff := t2.Sub(t1).Nanoseconds()

		fmt.Printf("resp: %v\nerr: %v\ndiff: %v\n", resp, err, diff)
	}

	fmt.Println("=====================================================")

	for i := 0; i < 3; i++ {
		t1 := time.Now()
		resp, err := client.GetRequest(ctx, &pb.GetData{
			Key:   &pb.Key{Key: keys[i]},
		})
		t2 := time.Now()
		diff := t2.Sub(t1).Nanoseconds()

		fmt.Printf("resp: %v\nerr: %v\ndiff: %v\n", resp, err, diff)
	}

	fmt.Println("=====================================================")

	for i := 0; i < 3; i++ {
		t1 := time.Now()
		resp, err := client.PutRequest(ctx, &pb.PutData{
			Key:   &pb.Key{Key:  keys[i]},
			Value: &pb.Value{Value:  updateValues[i]},
		})
		t2 := time.Now()
		diff := t2.Sub(t1).Nanoseconds()

		fmt.Printf("resp: %v\nerr: %v\ndiff: %v\n", resp, err, diff)
	}

	fmt.Println("=====================================================")

	for i := 0; i < 3; i++ {
		t1 := time.Now()
		resp, err := client.GetRequest(ctx, &pb.GetData{
			Key:   &pb.Key{Key: keys[i]},
		})
		t2 := time.Now()
		diff := t2.Sub(t1).Nanoseconds()

		fmt.Printf("resp: %v\nerr: %v\ndiff: %v\n", resp, err, diff)
	}
	//for {
	//	ctx, cancelFunc = context.WithTimeout(context.Background(), 10 * time.Second)
	//	defer cancelFunc()
	//
	//
	//	t1 = time.Now()
	//	resp, err = client.QingdaoRequest(ctx, &pb.MaoRequestData{RStr: "qingdao radar"})
	//	t2 = time.Now()
	//
	//	diff = t2.Sub(t1).Nanoseconds()
	//
	//	if max < diff {
	//		max = diff
	//	}
	//
	//	if min > diff {
	//		min = diff
	//	}
	//
	//	if err != nil {
	//		log.Printf("got data fail, %v", err)
	//		continue
	//	}
	//	if resp.Count % 10000 == 0 {
	//		log.Printf("data: %d, max: %d, min: %d, cur: %d, t1: %d, t2: %d,  %.20f, %.20f", resp.Count, max, min, diff, t1.Nanosecond(), t2.Nanosecond(), resp.Location.Latitude, resp.Location.Longitude)
	//	}
	//}
	//log.Printf("got data ok, %v, %v", resp, err)
	//log.Printf("data: %d, %.20f, %.20f", resp.Count, resp.Location.Latitude, resp.Location.Longitude)
}
