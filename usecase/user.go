package usecase

import (
	"context"
	"grpc-users/repository"
)

type UserUsecase interface {
	ListUser(ctx context.Context, id int, order int, orderType int, limit int) (repository.UserListEntity, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) ListUser(ctx context.Context, id int, order int, orderType int, limit int) (repository.UserListEntity, error) {
	return u.ur.ListUser(ctx, id, order, orderType, limit)
}
