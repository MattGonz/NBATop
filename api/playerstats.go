package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ==================================== TODO player profile ==================================
// 	url := fmt.Sprintf("https://data.nba.net/prod/v1/%s/players/%s_profile.json", yearStr, personID)

// PlayerProfile represents a player's stats across multiple seasons from the NBA Data API
// type PlayerProfile struct {
// 	Internal struct {
// 		PubDateTime             string `json:"pubDateTime"`
// 		IgorPath                string `json:"igorPath"`
// 		Xslt                    string `json:"xslt"`
// 		XsltForceRecompile      string `json:"xsltForceRecompile"`
// 		XsltInCache             string `json:"xsltInCache"`
// 		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
// 		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
// 		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
// 		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
// 	} `json:"_internal"`
// 	League struct {
// 		Standard struct {
// 			TeamID string `json:"teamId"`
// 			Stats  struct {
// 				Latest struct {
// 					SeasonYear    int    `json:"seasonYear"`
// 					SeasonStageID int    `json:"seasonStageId"`
// 					Ppg           string `json:"ppg"`
// 					Rpg           string `json:"rpg"`
// 					Apg           string `json:"apg"`
// 					Mpg           string `json:"mpg"`
// 					Topg          string `json:"topg"`
// 					Spg           string `json:"spg"`
// 					Bpg           string `json:"bpg"`
// 					Tpp           string `json:"tpp"`
// 					Ftp           string `json:"ftp"`
// 					Fgp           string `json:"fgp"`
// 					Assists       string `json:"assists"`
// 					Blocks        string `json:"blocks"`
// 					Steals        string `json:"steals"`
// 					Turnovers     string `json:"turnovers"`
// 					OffReb        string `json:"offReb"`
// 					DefReb        string `json:"defReb"`
// 					TotReb        string `json:"totReb"`
// 					Fgm           string `json:"fgm"`
// 					Fga           string `json:"fga"`
// 					Tpm           string `json:"tpm"`
// 					Tpa           string `json:"tpa"`
// 					Ftm           string `json:"ftm"`
// 					Fta           string `json:"fta"`
// 					PFouls        string `json:"pFouls"`
// 					Points        string `json:"points"`
// 					GamesPlayed   string `json:"gamesPlayed"`
// 					GamesStarted  string `json:"gamesStarted"`
// 					PlusMinus     string `json:"plusMinus"`
// 					Min           string `json:"min"`
// 					Dd2           string `json:"dd2"`
// 					Td3           string `json:"td3"`
// 				} `json:"latest"`
// 				CareerSummary struct {
// 					Tpp          string `json:"tpp"`
// 					Ftp          string `json:"ftp"`
// 					Fgp          string `json:"fgp"`
// 					Ppg          string `json:"ppg"`
// 					Rpg          string `json:"rpg"`
// 					Apg          string `json:"apg"`
// 					Bpg          string `json:"bpg"`
// 					Mpg          string `json:"mpg"`
// 					Spg          string `json:"spg"`
// 					Assists      string `json:"assists"`
// 					Blocks       string `json:"blocks"`
// 					Steals       string `json:"steals"`
// 					Turnovers    string `json:"turnovers"`
// 					OffReb       string `json:"offReb"`
// 					DefReb       string `json:"defReb"`
// 					TotReb       string `json:"totReb"`
// 					Fgm          string `json:"fgm"`
// 					Fga          string `json:"fga"`
// 					Tpm          string `json:"tpm"`
// 					Tpa          string `json:"tpa"`
// 					Ftm          string `json:"ftm"`
// 					Fta          string `json:"fta"`
// 					PFouls       string `json:"pFouls"`
// 					Points       string `json:"points"`
// 					GamesPlayed  string `json:"gamesPlayed"`
// 					GamesStarted string `json:"gamesStarted"`
// 					PlusMinus    string `json:"plusMinus"`
// 					Min          string `json:"min"`
// 					Dd2          string `json:"dd2"`
// 					Td3          string `json:"td3"`
// 				} `json:"careerSummary"`
// 				RegularSeason struct {
// 					Season []struct {
// 						SeasonYear int `json:"seasonYear"`
// 						Teams      []struct {
// 							TeamID       string `json:"teamId"`
// 							Ppg          string `json:"ppg"`
// 							Rpg          string `json:"rpg"`
// 							Apg          string `json:"apg"`
// 							Mpg          string `json:"mpg"`
// 							Topg         string `json:"topg"`
// 							Spg          string `json:"spg"`
// 							Bpg          string `json:"bpg"`
// 							Tpp          string `json:"tpp"`
// 							Ftp          string `json:"ftp"`
// 							Fgp          string `json:"fgp"`
// 							Assists      string `json:"assists"`
// 							Blocks       string `json:"blocks"`
// 							Steals       string `json:"steals"`
// 							Turnovers    string `json:"turnovers"`
// 							OffReb       string `json:"offReb"`
// 							DefReb       string `json:"defReb"`
// 							TotReb       string `json:"totReb"`
// 							Fgm          string `json:"fgm"`
// 							Fga          string `json:"fga"`
// 							Tpm          string `json:"tpm"`
// 							Tpa          string `json:"tpa"`
// 							Ftm          string `json:"ftm"`
// 							Fta          string `json:"fta"`
// 							PFouls       string `json:"pFouls"`
// 							Points       string `json:"points"`
// 							GamesPlayed  string `json:"gamesPlayed"`
// 							GamesStarted string `json:"gamesStarted"`
// 							PlusMinus    string `json:"plusMinus"`
// 							Min          string `json:"min"`
// 							Dd2          string `json:"dd2"`
// 							Td3          string `json:"td3"`
// 						} `json:"teams"`
// 						Total struct {
// 							Ppg          string `json:"ppg"`
// 							Rpg          string `json:"rpg"`
// 							Apg          string `json:"apg"`
// 							Mpg          string `json:"mpg"`
// 							Topg         string `json:"topg"`
// 							Spg          string `json:"spg"`
// 							Bpg          string `json:"bpg"`
// 							Tpp          string `json:"tpp"`
// 							Ftp          string `json:"ftp"`
// 							Fgp          string `json:"fgp"`
// 							Assists      string `json:"assists"`
// 							Blocks       string `json:"blocks"`
// 							Steals       string `json:"steals"`
// 							Turnovers    string `json:"turnovers"`
// 							OffReb       string `json:"offReb"`
// 							DefReb       string `json:"defReb"`
// 							TotReb       string `json:"totReb"`
// 							Fgm          string `json:"fgm"`
// 							Fga          string `json:"fga"`
// 							Tpm          string `json:"tpm"`
// 							Tpa          string `json:"tpa"`
// 							Ftm          string `json:"ftm"`
// 							Fta          string `json:"fta"`
// 							PFouls       string `json:"pFouls"`
// 							Points       string `json:"points"`
// 							GamesPlayed  string `json:"gamesPlayed"`
// 							GamesStarted string `json:"gamesStarted"`
// 							PlusMinus    string `json:"plusMinus"`
// 							Min          string `json:"min"`
// 							Dd2          string `json:"dd2"`
// 							Td3          string `json:"td3"`
// 						} `json:"total"`
// 					} `json:"season"`
// 				} `json:"regularSeason"`
// 			} `json:"stats"`
// 		} `json:"standard"`
// 	} `json:"league"`
// }

// PlayerGameLog represents a player's game-by-game stats for a given season
type PlayerGameLog struct {
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

// PlayerGameLog fetches and structures a player's game-by-game stats for a given season
func GetPlayerGameLog(displayYear, teamID, personID string) *PlayerGameLog {
	url := fmt.Sprintf("https://stats.nba.com/stats/playergamelog/?playerId=%s&season=%s&seasonType=Regular+Season&leagueId=00&dateFrom=&dateTo=", personID, displayYear)

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

	var result *PlayerGameLog
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panicln(err)
	}

	return result
}
