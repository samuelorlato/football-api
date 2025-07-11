package usecases

import (
	"github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports2 "github.com/samuelorlato/football-api/internal/domain/ports/repositories"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type getFanByEmailUsecase struct {
	fanRepository ports2.FanRepository
}

func NewGetFanByEmailUsecase(fanRepository ports2.FanRepository) ports.GetFanByEmailUsecase {
	return &getFanByEmailUsecase{
		fanRepository,
	}
}

func (g *getFanByEmailUsecase) Execute(email string) (*entities.Fan, error) {
	fan, err := g.fanRepository.FindByEmail(email)
	if err != nil {
		return nil, errs.NewInternalServerError()
	}

	return fan, nil
}
