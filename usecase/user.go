package usecase

import (
	"context"
	"grpc-users/repository"
)

type UserUsecase interface {
	ListUser(ctx context.Context, order string, limit int) (repository.UserListEntiy, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{ur: ur}
}

func (u *userUsecase) ListUser(ctx context.Context, order string, limit int) (repository.UserListEntiy, error) {
	return u.ur.ListUser(ctx, order, limit)
}
