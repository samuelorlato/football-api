package controllers

import (
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type authorizationController struct {
	registerUsecase ports2.RegisterUsecase
	loginUsecase    ports2.LoginUsecase
}

func AuthorizationController(registerUsecase ports2.RegisterUsecase, loginUsecase ports2.LoginUsecase) ports.AuthorizationController {
	return &authorizationController{
		registerUsecase,
		loginUsecase,
	}
}

func (a *authorizationController) Register(registerRequest dtos.RegisterRequest) error {
	err := a.registerUsecase.Execute(registerRequest.ToEntity())
	if err != nil {
		return err
	}

	return nil
}

func (a *authorizationController) Login(loginRequest dtos.LoginRequest) (*dtos.Token, error) {
	token, err := a.loginUsecase.Execute(loginRequest.User, loginRequest.Password)
	if err != nil {
		return nil, err
	}

	return dtos.NewToken(token.Token), nil
}
