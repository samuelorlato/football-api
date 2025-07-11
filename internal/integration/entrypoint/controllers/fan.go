package controllers

import (
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type fanController struct {
	getFanByEmailUsecase ports2.GetFanByEmailUsecase
}

func NewFanController(getFanByEmailUsecase ports2.GetFanByEmailUsecase) ports.FanController {
	return &fanController{
		getFanByEmailUsecase,
	}
}

func (c *fanController) GetByEmail(email string) (*dtos.Fan, error) {
	fan, err := c.getFanByEmailUsecase.Execute(email)
	if err != nil {
		return nil, err
	}

	return dtos.NewFan(fan), nil
}
