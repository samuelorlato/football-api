package dtos

type MatchFilterRequest struct {
	LeagueCode string  `param:"leagueCode" validate:"required"`
	Team       *string `query:"equipe"`
	Matchday   *int    `query:"rodada"`
}
