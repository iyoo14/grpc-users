package controller

import (
	"context"
	"grpc-users/usecase"
)

type AccountController interface {
	Authorize(cxt context.Context, accountID int, apiKey string) error
}

type accountController struct {
	u usecase.AccountUsecase
}

func NewAccountController(u usecase.AccountUsecase) AccountController {
	return &accountController{u}
}

func (c *accountController) Authorize(cxt context.Context, accountID int, apiKey string) error {
	_, err := c.u.Authorize(cxt, accountID, apiKey)
	return err
}
