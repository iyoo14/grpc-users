package controller

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	users, err := c.u.ListUser(cxt, int(req.GetId()), int(req.GetOrder()), int(req.GetOrderType()), int(req.GetLimit()))
	res := toListUserResponse(users)
	return res, err
}
