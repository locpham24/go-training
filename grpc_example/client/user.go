package main

import (
	"context"
	"fmt"
	pb "github.com/locpham24/go-training/grpc_example/proto"
	"google.golang.org/grpc"
)

func register(client pb.UserClient) (*pb.RegisterRes, error) {
	req := &pb.RegisterReq{
		Username: "admin",
		Email:    "admin@admin",
		Password: "123456",
		Phone:    "00000",
		Address:  "zzzzz",
	}
	return client.UserRegister(context.TODO(), req)
}
func main() {
	// 1. Connect to server at TCP port
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// 2. New client
	client := pb.NewUserClient(conn)
	// 3. Todo
	res, err := register(client)
	// 4
	if err != nil {
		fmt.Println("err:", err)
	}
	// 4. In ket qua
	fmt.Println("Response:", res)
}
