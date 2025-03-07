package controller

import (
	"grpc-users/model"
	"grpc-users/pb"
	"grpc-users/repository"
)

type ListUserRequest struct {
	Order string
	Limit int
}
type ListUserResponse struct {
	UserList []model.User
}

type DetailUserRequest struct {
	ID int
}

func toListUserResponse(users repository.UserListEntiy) *pb.ListUserResponse {
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
