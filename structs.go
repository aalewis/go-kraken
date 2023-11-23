package main

import (
	"time"
)

type AllTeams struct {
	Data []struct {
		ID          int    `json:"id"`
		FranchiseID int    `json:"franchiseId"`
		FullName    string `json:"fullName"`
		LeagueID    int    `json:"leagueId"`
		RawTricode  string `json:"rawTricode"`
		TriCode     string `json:"triCode"`
	} `json:"data"`
	Total int `json:"total"`
}

type TeamSchedule struct {
	PreviousSeason int    `json:"previousSeason"`
	CurrentSeason  int    `json:"currentSeason"`
	ClubTimezone   string `json:"clubTimezone"`
	ClubUTCOffset  string `json:"clubUTCOffset"`
	Games          []struct {
		ID       int    `json:"id"`
		Season   int    `json:"season"`
		GameType int    `json:"gameType"`
		GameDate string `json:"gameDate"`
		Venue    struct {
			Default string `json:"default"`
		} `json:"venue"`
		NeutralSite       bool      `json:"neutralSite"`
		StartTimeUTC      time.Time `json:"startTimeUTC"`
		EasternUTCOffset  string    `json:"easternUTCOffset"`
		VenueUTCOffset    string    `json:"venueUTCOffset"`
		VenueTimezone     string    `json:"venueTimezone"`
		GameState         string    `json:"gameState"`
		GameScheduleState string    `json:"gameScheduleState"`
		TvBroadcasts      []any     `json:"tvBroadcasts"`
		GameOutcome       struct {
			LastPeriodType string `json:"lastPeriodType"`
		} `json:"gameOutcome,omitempty"`
	} `json:"games"`
}

type Standings struct {
	Standings []struct {
		ConferenceAbbrev string `json:"conferenceAbbrev"`
		L10Losses        int    `json:"l10Losses"`
		Losses           int    `json:"losses"`
		SeasonID         int    `json:"seasonId"`
		StreakCode       string `json:"streakCode"`
		StreakCount      int    `json:"streakCount"`
		TeamName         struct {
			Default string `json:"default"`
			Fr      string `json:"fr"`
		} `json:"teamName"`
		TeamAbbrev struct {
			Default string `json:"default"`
		} `json:"teamAbbrev"`
		Ties int `json:"ties"`
		Wins int `json:"wins"`
	} `json:"standings"`
}
