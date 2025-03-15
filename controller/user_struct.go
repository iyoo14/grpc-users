package controller

import (
	"grpc-users/model"
	//"grpc-users/model"
	"grpc-users/pb"
	"grpc-users/repository"
)

func toListUserResponse(users repository.UserListEntity) *pb.ListUserResponse {
	var us []*pb.DetailUserResponse
	for _, user := range users.UserList {
		u := new(pb.DetailUserResponse)
		id, _ := user.ID.(int64)
		u.Id = int32(id)
		u.Name = user.Name.(string)
		u.Email = user.Email.(string)
		age, _ := user.Age.(int64)
		u.Age = int32(age)
		us = append(us, u)
	}
	return &pb.ListUserResponse{UserList: us}
}

func toUserRequestArg(req *pb.ListUserRequest) *model.ListUserRequest {
	return &model.ListUserRequest{
		Order:     string(req.GetOrder()),
		Limit:     req.GetLimit(),
		OrderType: int32(req.GetOrderType()),
	}
}
