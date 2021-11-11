package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mattgonz/nbatop/types"
)

// Standings fetches and structures the conference standings data from the NBA Data API
func Standings() types.NBAStandings {
	url := "https://data.nba.net/prod/v2/current/standings_conference.json"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result types.NBAStandings
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result
}
