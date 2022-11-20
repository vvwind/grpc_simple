package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/vvwind/grpc_simple/protos"
	"google.golang.org/grpc"
)

func main() {
	myurl := os.Args[1]
	conn, err := grpc.Dial(myurl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect! %v", err)
	}
	defer conn.Close()
	c := pb.NewDatabusServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s1, errs1 := strconv.ParseFloat(os.Args[2], 32)
	if errs1 != nil {
		log.Fatalf("errror converting! %v", errs1)
	}
	s2, errs2 := strconv.ParseFloat(os.Args[3], 32)
	if errs2 != nil {
		log.Fatalf("errror converting! %v", errs2)
	}
	r, errr := c.Send(ctx, &pb.SendRequest{Prm1: float32(s1), Prm2: float32(s2)})

	if errr != nil {
		log.Fatalf("cant send request %v", errr)
	}
	log.Println("Result: ", r.Result)
}
