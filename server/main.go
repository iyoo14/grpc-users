package main

import (
	"context"
	"encoding/json"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	/*
		exe, _ := os.Executable()
		exePath = filepath.Dir(exe)
		setEnv()
		db := infra.Connect(dsn)
	*/

	db := ctx.Value("db_connect").(*sqlx.DB)
	fmt.Println(db)

	//fmt.Println("Hello World")
	//tr := transaction.NewTransaction(db)
	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)
	res, err := uc.ListUser(ctx, req)
	return res, err
}

type Request struct {
	Data *pb.ListUserRequest
}

func authorizeWithRequest(ctx context.Context) (context.Context, error) {
	// カスタムインターセプターで設定されたリクエストデータを取り出す
	//fmt.Printf("AuthCtx: %+v\n", ctx)
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		fmt.Println("MD error", err)
		return nil, err
	}
	req := ctx.Value("request_data").(*pb.ListUserRequest)
	fmt.Printf("AuthCtx: %+v\n", req.GetAccountId())
	db := ctx.Value("db_connect").(*sqlx.DB)
	fmt.Printf("DB: %+v\n", db)
	ar := repository.NewAccountRepository(db)
	au := usecase.NewAccountUsecase(ar)
	uc := controller.NewAccountController(au)
	err = uc.Authorize(ctx, int(req.GetAccountId()), token)
	if err != nil {
		fmt.Println("token not match", token)
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}
	fmt.Println("token match", token)
	return ctx, nil
}

func dbConnectInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	exe, _ := os.Executable()
	exePath = filepath.Dir(exe)
	setEnv()
	db := infra.Connect(dsn)
	newCtx := context.WithValue(ctx, "db_connect", db)
	return handler(newCtx, req)
}

func requestDataInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// リクエストデータをコンテキストに保存

	//fmt.Printf("requestdate: %+v\n", req)
	//fmt.Printf("request order: %+v\n", req.(*pb.ListUserRequest).GetLimit())
	newCtx := context.WithValue(ctx, "request_data", req.(*pb.ListUserRequest))
	//fmt.Printf("requestdateCtx: %+v\n", newCtx)
	return handler(newCtx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				dbConnectInterceptor,
				requestDataInterceptor,
				// カスタムインターセプターを最初に実行
				grpc_auth.UnaryServerInterceptor(authorizeWithRequest), // grpc_auth を実行
			),
		))
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
