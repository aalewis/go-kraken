package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func main() {
	// cpu profiling
	f, _ := os.Create("cpu.pgo")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	if len(os.Args) < 2 {
		log.Fatal("Need arguments")
	}
	var arg string
	for i, j := range os.Args {
		if i == 0 {
			continue
		}
		arg += j + " "
	}
	teamAbbr := getTeamAbbr(arg)
	if teamAbbr == "---" {
		log.Fatal("Didn't find team, sorry")
	}
	nextGame := getNextGame(teamAbbr)
	log.Println(getTeamStats(teamAbbr) + ", " + nextGame)
}

func getURLBytes(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func getTeamAbbr(team string) string {
	bodyBytes, err := getURLBytes("https://api.nhle.com/stats/rest/en/team")
	if err != nil {
		log.Println(err)
		return "---"
	}
	var t AllTeams
	json.Unmarshal(bodyBytes, &t)
	if len(t.Data) > 0 {
		for _, j := range t.Data {
			if strings.Contains(strings.ToLower(strings.ReplaceAll(j.FullName, " ", "")), strings.ToLower(strings.ReplaceAll(team, " ", ""))) {
				return j.TriCode
			}
		}
	}
	return "---"
}

func getNextGame(team string) string {
	bodyBytes, err := getURLBytes("https://api-web.nhle.com/v1/club-schedule-season/" + team + "/now")
	if err != nil {
		log.Println(err)
		return "---"
	}

	now := time.Now()
	var t TeamSchedule
	json.Unmarshal(bodyBytes, &t)
	if len(t.Games) > 0 {
		for _, j := range t.Games {
			if j.StartTimeUTC.After(now) {
				tz, err := time.LoadLocation(t.ClubTimezone)
				if err != nil { // Always check errors even if they should not happen.
					panic(err)
				}
				nextGame := "Next Game: "
				nextGame += j.StartTimeUTC.In(tz).Format(time.RFC850)
				if j.AwayTeam.Abbrev == team {
					nextGame += " against " + getTeamAbbr(j.HomeTeam.PlaceName.Default) + " (Away)"
				} else {
					nextGame += " against " + getTeamAbbr(j.AwayTeam.PlaceName.Default) + " (Home)"
				}
				return nextGame
			}
		}
	}
	return "---"
}

func getTeamStats(team string) string {
	bodyBytes, err := getURLBytes("https://api-web.nhle.com/v1/standings/now")
	if err != nil {
		log.Println(err)
		return "---"
	}

	var stats string
	var t Standings
	json.Unmarshal(bodyBytes, &t)
	if len(t.Standings) > 0 {
		for i, j := range t.Standings {
			if j.TeamAbbrev.Default == team {
				log.Println("teamstats", j.TeamName.Default)
				stats += "Team: " + j.TeamName.Default
				stats += ", Conf: " + j.ConferenceAbbrev
				stats += ", Standing: " + strconv.Itoa(i+1)
				stats += ", Season: " + strconv.Itoa(j.Wins) + "W"
				stats += strconv.Itoa(j.Losses) + "L"
				stats += ", Streak: " + strconv.Itoa(j.StreakCount) + j.StreakCode
			}
		}
	}
	return stats
}
