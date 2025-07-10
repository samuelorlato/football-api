package controllers

import (
	"fmt"

	ports2 "github.com/samuelorlato/football-api/internal/application/ports"
	"github.com/samuelorlato/football-api/internal/integration/entrypoint/dtos"
	"github.com/samuelorlato/football-api/internal/integration/ports"
)

type footballController struct {
	getLeaguesUsecase ports2.GetLeaguesUsecase
	getMatchesUsecase ports2.GetMatchesUseCase
}

func NewFootballController(getLeaguesUsecase ports2.GetLeaguesUsecase, getMatchesUsecase ports2.GetMatchesUseCase) ports.FootballController {
	return &footballController{
		getLeaguesUsecase: getLeaguesUsecase,
		getMatchesUsecase: getMatchesUsecase,
	}
}

func (f *footballController) GetMatches(leagueCode string, team *string, matchday *int) (*dtos.Matches, error) {
	matchesEntities, err := f.getMatchesUsecase.Execute(leagueCode, team, matchday)
	if err != nil {
		return nil, err
	}

	matches := make([]dtos.Match, len(matchesEntities))
	for i, matchEntity := range matchesEntities {
		score := fmt.Sprintf("%dx%d", matchEntity.HomeScore, matchEntity.AwayScore)
		matches[i] = dtos.NewMatch(matchEntity, score)
	}

	if matchday != nil {
		return &dtos.Matches{
			Matchday: *matchday,
			Matches:  matches,
		}, nil
	}

	// TODO: review business rules
	return nil, nil
}

func (f *footballController) GetLeagues() ([]dtos.League, error) {
	leaguesEntities, err := f.getLeaguesUsecase.Execute()
	if err != nil {
		return nil, err
	}

	leagues := make([]dtos.League, len(leaguesEntities))
	for i, leagueEntity := range leaguesEntities {
		leagues[i] = dtos.NewLeague(leagueEntity)
	}

	return leagues, nil
}
