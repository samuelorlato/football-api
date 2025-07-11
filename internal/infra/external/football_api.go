package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/samuelorlato/football-api/internal/domain/entities"
	ports "github.com/samuelorlato/football-api/internal/domain/ports/external"
	"github.com/samuelorlato/football-api/internal/infra/external/dtos"
)

type footballAPI struct {
	httpClient *http.Client
	baseURL    string
	authToken  string
}

func NewFootballAPI(httpClient *http.Client, baseURL string, authToken string) ports.FootballAPI {
	return &footballAPI{
		httpClient: httpClient,
		baseURL:    baseURL,
		authToken:  authToken,
	}
}

func (f *footballAPI) doRequestAndReadBody(method string, URL string, target interface{}) error {
	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", f.authToken)

	res, err := f.httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	return nil
}

func (f *footballAPI) GetMatches(leagueCode string) ([]entities.Match, error) {
	url := fmt.Sprintf("%s/competitions/%s/matches", f.baseURL, leagueCode)
	var matchesRes dtos.MatchesResponse
	err := f.doRequestAndReadBody("GET", url, &matchesRes)
	if err != nil {
		return nil, err
	}

	matches := matchesRes.ToEntities()

	return matches, nil
}

func (f *footballAPI) GetMatchdayMatches(leagueCode string, matchday int) ([]entities.Match, error) {
	url := fmt.Sprintf("%s/competitions/%s/matches?matchday=%d", f.baseURL, leagueCode, matchday)
	var matchesRes dtos.MatchesResponse
	err := f.doRequestAndReadBody("GET", url, &matchesRes)
	if err != nil {
		return nil, err
	}

	matches := matchesRes.ToEntities()

	return matches, nil
}

func (f *footballAPI) GetLeagues() ([]entities.League, error) {
	url := fmt.Sprintf("%s/competitions", f.baseURL)
	var competitionsRes dtos.CompetitionsResponse
	err := f.doRequestAndReadBody("GET", url, &competitionsRes)
	if err != nil {
		return nil, err
	}

	leagues := competitionsRes.ToEntities()

	return leagues, nil
}
