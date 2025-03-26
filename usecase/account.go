package usecase

import (
	"context"
	"grpc-users/repository"
)

type AccountUsecase interface {
	Authorize(ctx context.Context, accountId int, apiKey string) (repository.AccountEntity, error)
}

type accountUsecase struct {
	ar repository.AccountRepository
}

func NewAccountUsecase(ar repository.AccountRepository) AccountUsecase {
	return &accountUsecase{ar: ar}
}

func (u *accountUsecase) Authorize(ctx context.Context, accountId int, apiKey string) (repository.AccountEntity, error) {
	return u.ar.GetApiKey(ctx, accountId, apiKey)
}
