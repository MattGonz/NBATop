package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// NBAStandings represents current NBA standings from the NBA Data API
type NBAStandings struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	League struct {
		Standard struct {
			SeasonYear    int `json:"seasonYear"`
			SeasonStageID int `json:"seasonStageId"`
			Conference    struct {
				East []struct {
					TeamID                 string `json:"teamId"`
					Win                    string `json:"win"`
					Loss                   string `json:"loss"`
					WinPct                 string `json:"winPct"`
					WinPctV2               string `json:"winPctV2"`
					LossPct                string `json:"lossPct"`
					LossPctV2              string `json:"lossPctV2"`
					GamesBehind            string `json:"gamesBehind"`
					DivGamesBehind         string `json:"divGamesBehind"`
					ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
					ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
					ConfRank               string `json:"confRank"`
					ConfWin                string `json:"confWin"`
					ConfLoss               string `json:"confLoss"`
					DivWin                 string `json:"divWin"`
					DivLoss                string `json:"divLoss"`
					HomeWin                string `json:"homeWin"`
					HomeLoss               string `json:"homeLoss"`
					AwayWin                string `json:"awayWin"`
					AwayLoss               string `json:"awayLoss"`
					LastTenWin             string `json:"lastTenWin"`
					LastTenLoss            string `json:"lastTenLoss"`
					Streak                 string `json:"streak"`
					DivRank                string `json:"divRank"`
					IsWinStreak            bool   `json:"isWinStreak"`
					TeamSitesOnly          struct {
						TeamKey            string `json:"teamKey"`
						TeamName           string `json:"teamName"`
						TeamCode           string `json:"teamCode"`
						TeamNickname       string `json:"teamNickname"`
						TeamTricode        string `json:"teamTricode"`
						ClinchedConference string `json:"clinchedConference"`
						ClinchedDivision   string `json:"clinchedDivision"`
						ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
						StreakText         string `json:"streakText"`
					} `json:"teamSitesOnly"`
					TieBreakerPts string `json:"tieBreakerPts"`
					SortKey       struct {
						DefaultOrder   int `json:"defaultOrder"`
						Nickname       int `json:"nickname"`
						Win            int `json:"win"`
						Loss           int `json:"loss"`
						WinPct         int `json:"winPct"`
						GamesBehind    int `json:"gamesBehind"`
						ConfWinLoss    int `json:"confWinLoss"`
						DivWinLoss     int `json:"divWinLoss"`
						HomeWinLoss    int `json:"homeWinLoss"`
						AwayWinLoss    int `json:"awayWinLoss"`
						LastTenWinLoss int `json:"lastTenWinLoss"`
						Streak         int `json:"streak"`
					} `json:"sortKey"`
				} `json:"east"`
				West []struct {
					TeamID                 string `json:"teamId"`
					Win                    string `json:"win"`
					Loss                   string `json:"loss"`
					WinPct                 string `json:"winPct"`
					WinPctV2               string `json:"winPctV2"`
					LossPct                string `json:"lossPct"`
					LossPctV2              string `json:"lossPctV2"`
					GamesBehind            string `json:"gamesBehind"`
					DivGamesBehind         string `json:"divGamesBehind"`
					ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
					ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
					ConfRank               string `json:"confRank"`
					ConfWin                string `json:"confWin"`
					ConfLoss               string `json:"confLoss"`
					DivWin                 string `json:"divWin"`
					DivLoss                string `json:"divLoss"`
					HomeWin                string `json:"homeWin"`
					HomeLoss               string `json:"homeLoss"`
					AwayWin                string `json:"awayWin"`
					AwayLoss               string `json:"awayLoss"`
					LastTenWin             string `json:"lastTenWin"`
					LastTenLoss            string `json:"lastTenLoss"`
					Streak                 string `json:"streak"`
					DivRank                string `json:"divRank"`
					IsWinStreak            bool   `json:"isWinStreak"`
					TeamSitesOnly          struct {
						TeamKey            string `json:"teamKey"`
						TeamName           string `json:"teamName"`
						TeamCode           string `json:"teamCode"`
						TeamNickname       string `json:"teamNickname"`
						TeamTricode        string `json:"teamTricode"`
						ClinchedConference string `json:"clinchedConference"`
						ClinchedDivision   string `json:"clinchedDivision"`
						ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
						StreakText         string `json:"streakText"`
					} `json:"teamSitesOnly"`
					TieBreakerPts string `json:"tieBreakerPts"`
					SortKey       struct {
						DefaultOrder   int `json:"defaultOrder"`
						Nickname       int `json:"nickname"`
						Win            int `json:"win"`
						Loss           int `json:"loss"`
						WinPct         int `json:"winPct"`
						GamesBehind    int `json:"gamesBehind"`
						ConfWinLoss    int `json:"confWinLoss"`
						DivWinLoss     int `json:"divWinLoss"`
						HomeWinLoss    int `json:"homeWinLoss"`
						AwayWinLoss    int `json:"awayWinLoss"`
						LastTenWinLoss int `json:"lastTenWinLoss"`
						Streak         int `json:"streak"`
					} `json:"sortKey"`
				} `json:"west"`
			} `json:"conference"`
		} `json:"standard"`
	} `json:"league"`
}

// GetStandings fetches and structures the conference standings data from the NBA Data API
func GetStandings() *NBAStandings {
	url := "https://data.nba.net/prod/v2/current/standings_conference.json"
	response, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln(err)
	}

	var result *NBAStandings
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panicln(err)
	}
	return result
}
