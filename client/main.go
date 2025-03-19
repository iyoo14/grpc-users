package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"grpc-users/pb"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)
	callListUser(client)
}

func callListUser(client pb.UsersServiceClient) {
	md := metadata.New(map[string]string{"authorization": "Bearer test-token"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.ListUser(ctx, &pb.ListUserRequest{
		Order:     1,
		OrderType: 1,
		Limit:     10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range res.GetUserList() {
		// Example: Print the user details
		println(user.GetId(), user.GetName(), user.GetEmail(), user.GetAge())
	}
}
