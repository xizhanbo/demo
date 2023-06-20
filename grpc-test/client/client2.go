package main

import (
	pb "back/pbs"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := grpc.Dial(":9998", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	// 初始化grpc客户端
	findClassClient := pb.NewFindClass2Client(client)

	class, err := findClassClient.MyClass2(context.Background(), &pb.Student2{
		Name: "root",
		Age:  3,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(class.ClassName)
}
