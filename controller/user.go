package controller

import (
	"context"
	"grpc-users/pb"
	"grpc-users/usecase"
)

type UserController interface {
	ListUser(cxt context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error)
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (c *userController) ListUser(cxt context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := c.u.ListUser(cxt, req.GetOrder(), int(req.GetLimit()))
	res := toListUserResponse(users)
	return res, err
}
