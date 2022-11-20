package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/vvwind/grpc_simple/protos"
	"google.golang.org/grpc"
)

type simpleServer struct {
	pb.UnimplementedDatabusServiceServer
}

func (s *simpleServer) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	if os.Args[2] == "mul" {
		fmt.Println(in.Prm1, " * ", in.Prm2, " = ", in.Prm1*in.Prm2)
		return &pb.SendResponse{Result: in.Prm1 * in.Prm2}, nil
	}
	if os.Args[2] == "add" {
		fmt.Println(in.Prm1, " + ", in.Prm2, " = ", in.Prm1+in.Prm2)
		return &pb.SendResponse{Result: in.Prm1 + in.Prm2}, nil
	}
	if os.Args[2] == "sub" {
		fmt.Println(in.Prm1, " - ", in.Prm2, " = ", in.Prm1-in.Prm2)
		return &pb.SendResponse{Result: in.Prm1 - in.Prm2}, nil
	}
	if os.Args[2] == "div" {
		if in.Prm2 != 0 {
			fmt.Println(in.Prm1, " / ", in.Prm2, " = ", in.Prm1/in.Prm2)
			return &pb.SendResponse{Result: in.Prm1 / in.Prm2}, nil
		} else {
			return &pb.SendResponse{}, errors.New("Cant divide by 0!")
		}

	}

	return &pb.SendResponse{}, errors.New("Error happened")
}

func main() {
	myport := os.Args[1]
	lis, err := net.Listen("tcp", ":"+myport)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterDatabusServiceServer(s, &simpleServer{})
	log.Printf("server is listening at %v", lis.Addr())

	if err2 := s.Serve(lis); err2 != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
