package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// TeamGameLog represents all of a team's games in a season, from the NBA Data API
type TeamGameLog struct {
	Resource   string `json:"resource"`
	Parameters struct {
		TeamID     int         `json:"TeamID"`
		LeagueID   interface{} `json:"LeagueID"`
		Season     string      `json:"Season"`
		SeasonType string      `json:"SeasonType"`
		DateFrom   interface{} `json:"DateFrom"`
		DateTo     interface{} `json:"DateTo"`
	} `json:"parameters"`
	ResultSets []struct {
		Name    string          `json:"name"`
		Headers []string        `json:"headers"`
		RowSet  [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}

// GetTeamGameLog gets the game log for a given team ID
func GetTeamGameLog(id string) *TeamGameLog {
	url := "http://stats.nba.com/stats/teamgamelog/?TeamID=" + id + "&season=2021-22&seasonType=Regular+Season"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicln(err)
	}

	req.Header = http.Header{
		"Host":       []string{"stats.nba.com"},
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0"},
		"Accept":     []string{"application/json, text/plain, */*"},
		"Referer":    []string{"https://stats.nba.com/"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	var result *TeamGameLog
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panicln(err)
	}
	return result
}
