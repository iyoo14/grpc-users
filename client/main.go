package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	res, err := client.ListUser(context.Background(), &pb.ListUserRequest{
		Order: "id asc",
		Limit: 10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range res.GetUserList() {
		// Example: Print the user details
		println(user.GetName(), user.GetEmail(), user.GetAge())
	}
}
