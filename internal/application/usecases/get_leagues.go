package usecases

import (
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/external"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type getLeaguesUsecase struct {
	footballAPI ports.FootballAPI
}

func NewGetLeaguesUsecase(footballAPI ports.FootballAPI) ports2.GetLeaguesUsecase {
	return &getLeaguesUsecase{
		footballAPI: footballAPI,
	}
}

func (g *getLeaguesUsecase) Execute() ([]entities.League, error) {
	leagues, err := g.footballAPI.GetLeagues()
	if err != nil {
		return nil, errs.NewInternalServerError()
	}

	return leagues, nil
}
