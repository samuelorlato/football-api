package usecases

import (
	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/external"
	"github.com/samuelorlato/football-api/pkg/errs"
)

type getMatchesUsecase struct {
	footballAPI ports.FootballAPI
}

func NewGetMatchesUsecase(footballAPI ports.FootballAPI) ports2.GetMatchesUseCase {
	return &getMatchesUsecase{
		footballAPI: footballAPI,
	}
}

func (g *getMatchesUsecase) Execute(leagueCode string, team *string, matchday *int) ([]entities.Match, error) {
	// TODO: review business rules

	var matches []entities.Match
	var err error
	if matchday != nil {
		matches, err = g.footballAPI.GetMatchdayFinishedMatches(leagueCode, *matchday)
		if err != nil {
			return nil, errs.NewInternalServerError()
		}
	} else {
		matches, err = g.footballAPI.GetFinishedMatches(leagueCode)
		if err != nil {
			return nil, errs.NewInternalServerError()
		}
	}

	if team != nil {
		var teamMatches []entities.Match
		for _, match := range matches {
			if match.HomeTeam == *team || match.AwayTeam == *team {
				teamMatches = append(teamMatches, match)
			}
		}

		return teamMatches, nil
	}

	return matches, nil
}
