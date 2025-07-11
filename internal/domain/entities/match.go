package entities

import "time"

type Match struct {
	UTCDate   time.Time
	HomeTeam  string
	AwayTeam  string
	HomeScore *int
	AwayScore *int
}
