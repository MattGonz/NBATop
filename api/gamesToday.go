package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mattgonz/nbatop/types"
)

func GamesToday() types.NBAScoreboard {
	today := time.Now().Format("20060102")
	url := fmt.Sprintf("http://data.nba.net/prod/v2/%s/scoreboard.json", today)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result types.NBAScoreboard
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result
}
