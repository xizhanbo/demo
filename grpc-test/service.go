package main

import (
	pb "back/pbs"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Teacher struct {
	pb.FindClassServer
}

func (t *Teacher) MyClass(ctx context.Context, stu *pb.Student) (*pb.Class, error) {
	cla := pb.Class{}
	if stu.Name == "root" && stu.Age == 3 {
		cla.ClassName = "Chinese"
		return &cla, nil
	} else {
		return &cla, errors.New("no this student")
	}

}

type Teacher2 struct {
	pb.FindClass2Server
}

func (t *Teacher2) MyClass2(ctx context.Context, stu *pb.Student2) (*pb.Class2, error) {
	cla := pb.Class2{}
	if stu.Name == "root" && stu.Age == 3 {
		cla.ClassName = "Chinese2"
		return &cla, nil
	} else {
		return &cla, errors.New("no this student")
	}

}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9998))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	pb.RegisterFindClassServer(grpcServer, new(Teacher))
	pb.RegisterFindClass2Server(grpcServer, new(Teacher2))
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
