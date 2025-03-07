package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"grpc-users/controller"
	"grpc-users/infra"
	"grpc-users/pb"
	"grpc-users/repository"
	"grpc-users/usecase"
	"log"
	"net"
	"os"
	"path/filepath"
)

var dsn string
var exePath string

type config struct {
	Dsn    string `json:"dsn"`
	Suffix string `json:suffix`
}

type user struct {
	name string
	id   int
}
type server struct {
	pb.UnimplementedUsersServiceServer
}

func (*server) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	/*
		return sample
		res := &pb.ListUserResponse{
			UserList: []*pb.DetailUserResponse{
				{
					Id:    1,
					Name:  "John",
					Email: "john@example.com",
					Age:   20,
				},
				{
					Id:    2,
					Name:  "Alice",
					Email: "alice@example.com",
					Age:   25,
				},
			},
		}
	*/
	exe, _ := os.Executable()
	exePath = filepath.Dir(exe)
	setEnv()
	db := infra.Connect(dsn)
	fmt.Println(db)

	fmt.Println("Hello World")
	//tr := transaction.NewTransaction(db)
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	res, err := uc.ListUser(ctx, req)
	return res, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &server{})

	fmt.Println("server listening at localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setEnv() {
	fname := filepath.Join(exePath, "..", "config", "env.json")
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	dsn = cfg.Dsn
}
