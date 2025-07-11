package controllers

import (
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type userController struct {
	getUserByIDUsecase ports2.GetUserByIDUsecase
}

func NewUserController(getUserByIDUsecase ports2.GetUserByIDUsecase) ports.UserController {
	return &userController{
		getUserByIDUsecase,
	}
}

func (u *userController) GetByID(ID string) (*dtos.User, error) {
	user, err := u.getUserByIDUsecase.Execute(ID)
	if err != nil {
		return nil, err
	}

	return dtos.NewUser(user), nil
}
