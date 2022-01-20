package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Players struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	League struct {
		Players []struct {
			FirstName            string `json:"firstName"`
			LastName             string `json:"lastName"`
			TemporaryDisplayName string `json:"temporaryDisplayName,omitempty"`
			PersonID             string `json:"personId"`
			TeamID               string `json:"teamId"`
			Jersey               string `json:"jersey"`
			IsActive             bool   `json:"isActive"`
			Pos                  string `json:"pos"`
			HeightFeet           string `json:"heightFeet"`
			HeightInches         string `json:"heightInches"`
			HeightMeters         string `json:"heightMeters"`
			WeightPounds         string `json:"weightPounds"`
			WeightKilograms      string `json:"weightKilograms"`
			DateOfBirthUTC       string `json:"dateOfBirthUTC"`
			TeamSitesOnly        struct {
				PlayerCode         string `json:"playerCode"`
				PosFull            string `json:"posFull"`
				DisplayAffiliation string `json:"displayAffiliation"`
				FreeAgentCode      string `json:"freeAgentCode"`
			} `json:"teamSitesOnly,omitempty"`
			Teams []struct {
				TeamID      string `json:"teamId"`
				SeasonStart string `json:"seasonStart"`
				SeasonEnd   string `json:"seasonEnd"`
			} `json:"teams"`
			Draft struct {
				TeamID     string `json:"teamId"`
				PickNum    string `json:"pickNum"`
				RoundNum   string `json:"roundNum"`
				SeasonYear string `json:"seasonYear"`
			} `json:"draft"`
			NbaDebutYear    string `json:"nbaDebutYear"`
			YearsPro        string `json:"yearsPro"`
			CollegeName     string `json:"collegeName"`
			LastAffiliation string `json:"lastAffiliation"`
			Country         string `json:"country"`
			IsallStar       bool   `json:"isallStar,omitempty"`
		} `json:"standard"`
		Africa     []interface{} `json:"africa"`
		Sacramento []struct {
			FirstName            string `json:"firstName"`
			LastName             string `json:"lastName"`
			TemporaryDisplayName string `json:"temporaryDisplayName"`
			PersonID             string `json:"personId"`
			TeamID               string `json:"teamId"`
			Jersey               string `json:"jersey"`
			IsActive             bool   `json:"isActive"`
			Pos                  string `json:"pos"`
			HeightFeet           string `json:"heightFeet"`
			HeightInches         string `json:"heightInches"`
			HeightMeters         string `json:"heightMeters"`
			WeightPounds         string `json:"weightPounds"`
			WeightKilograms      string `json:"weightKilograms"`
			DateOfBirthUTC       string `json:"dateOfBirthUTC"`
			TeamSitesOnly        struct {
				PlayerCode         string `json:"playerCode"`
				PosFull            string `json:"posFull"`
				DisplayAffiliation string `json:"displayAffiliation"`
				FreeAgentCode      string `json:"freeAgentCode"`
			} `json:"teamSitesOnly"`
			Teams []struct {
				TeamID      string `json:"teamId"`
				SeasonStart string `json:"seasonStart"`
				SeasonEnd   string `json:"seasonEnd"`
			} `json:"teams"`
			Draft struct {
				TeamID     string `json:"teamId"`
				PickNum    string `json:"pickNum"`
				RoundNum   string `json:"roundNum"`
				SeasonYear string `json:"seasonYear"`
			} `json:"draft"`
			NbaDebutYear    string `json:"nbaDebutYear"`
			YearsPro        string `json:"yearsPro"`
			CollegeName     string `json:"collegeName"`
			LastAffiliation string `json:"lastAffiliation"`
			Country         string `json:"country"`
		} `json:"sacramento"`
		Vegas []struct {
			FirstName            string `json:"firstName"`
			LastName             string `json:"lastName"`
			TemporaryDisplayName string `json:"temporaryDisplayName,omitempty"`
			PersonID             string `json:"personId"`
			TeamID               string `json:"teamId"`
			Jersey               string `json:"jersey"`
			IsActive             bool   `json:"isActive"`
			Pos                  string `json:"pos"`
			HeightFeet           string `json:"heightFeet"`
			HeightInches         string `json:"heightInches"`
			HeightMeters         string `json:"heightMeters"`
			WeightPounds         string `json:"weightPounds"`
			WeightKilograms      string `json:"weightKilograms"`
			DateOfBirthUTC       string `json:"dateOfBirthUTC"`
			TeamSitesOnly        struct {
				PlayerCode         string `json:"playerCode"`
				PosFull            string `json:"posFull"`
				DisplayAffiliation string `json:"displayAffiliation"`
				FreeAgentCode      string `json:"freeAgentCode"`
			} `json:"teamSitesOnly,omitempty"`
			Teams []struct {
				TeamID      string `json:"teamId"`
				SeasonStart string `json:"seasonStart"`
				SeasonEnd   string `json:"seasonEnd"`
			} `json:"teams"`
			Draft struct {
				TeamID     string `json:"teamId"`
				PickNum    string `json:"pickNum"`
				RoundNum   string `json:"roundNum"`
				SeasonYear string `json:"seasonYear"`
			} `json:"draft"`
			NbaDebutYear    string `json:"nbaDebutYear"`
			YearsPro        string `json:"yearsPro"`
			CollegeName     string `json:"collegeName"`
			LastAffiliation string `json:"lastAffiliation"`
			Country         string `json:"country"`
			IsallStar       bool   `json:"isallStar,omitempty"`
		} `json:"vegas"`
		Utah []struct {
			FirstName            string `json:"firstName"`
			LastName             string `json:"lastName"`
			TemporaryDisplayName string `json:"temporaryDisplayName"`
			PersonID             string `json:"personId"`
			TeamID               string `json:"teamId"`
			Jersey               string `json:"jersey"`
			IsActive             bool   `json:"isActive"`
			Pos                  string `json:"pos"`
			HeightFeet           string `json:"heightFeet"`
			HeightInches         string `json:"heightInches"`
			HeightMeters         string `json:"heightMeters"`
			WeightPounds         string `json:"weightPounds"`
			WeightKilograms      string `json:"weightKilograms"`
			DateOfBirthUTC       string `json:"dateOfBirthUTC"`
			TeamSitesOnly        struct {
				PlayerCode         string `json:"playerCode"`
				PosFull            string `json:"posFull"`
				DisplayAffiliation string `json:"displayAffiliation"`
				FreeAgentCode      string `json:"freeAgentCode"`
			} `json:"teamSitesOnly"`
			Teams []interface{} `json:"teams"`
			Draft struct {
				TeamID     string `json:"teamId"`
				PickNum    string `json:"pickNum"`
				RoundNum   string `json:"roundNum"`
				SeasonYear string `json:"seasonYear"`
			} `json:"draft"`
			NbaDebutYear    string `json:"nbaDebutYear"`
			YearsPro        string `json:"yearsPro"`
			CollegeName     string `json:"collegeName"`
			LastAffiliation string `json:"lastAffiliation"`
			Country         string `json:"country"`
		} `json:"utah"`
	} `json:"league"`
}

// GetPlayers fetches and structures this season's active players, fetched from the NBA Data API
func GetPlayers() *Players {
	t := time.Now()
	month := t.Month()
	year := t.Year()

	// this year's season hasn't started, roll back a year
	if month < 9 {
		year -= 1
	}
	yearStr := strconv.Itoa(year)

	url := fmt.Sprintf("https://data.nba.net/prod/v1/%s/players.json", yearStr)
	response, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln(err)
	}

	var activePlayers *Players
	if err := json.Unmarshal(body, &activePlayers); err != nil {
		log.Panicln(err)
	}

	return activePlayers
}
