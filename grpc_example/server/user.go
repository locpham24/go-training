package main

import (
	context "context"
	"fmt"

	"github.com/joho/godotenv"
	pb "github.com/locpham24/go-training/grpc_example/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userService struct {
}

func NewUserService() *userService {
	return &userService{}
}

func (u userService) UserLogin(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	fmt.Println(in)
	res := &pb.LoginRes{
		Ok: true,
		Data: &pb.AccessToken{
			Token: "abc",
		},
	}
	return res, nil
}

func (u userService) UserRegister(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	fmt.Println(req)
	res := &pb.RegisterRes{
		Ok: true,
		Data: &pb.AccessToken{
			Token: "abc",
		},
	}
	return res, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 1. Listen/Open a TPC connect at port
	lis, _ := net.Listen("tcp", ":50051")
	// 2. Tao server tu GRP
	grpcServer := grpc.NewServer()
	// 3. Map service to server
	pb.RegisterUserServer(grpcServer, &userService{})
	// 4. Binding port
	fmt.Println("Start service")
	grpcServer.Serve(lis)
}
