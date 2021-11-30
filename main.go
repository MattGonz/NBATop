package main

import (
	"sync"

	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/nbatop"
)

func main() {

	// TODO: team stats from team ID (endpoint seems to be down or deprecated)
	// TODO: continue to iron out structs / split into multiple files
	// TODO: cache standings / games today / team stats?

	// WORKING: get games today
	// WORKING: get standings
	// WORKING: standings -> gocui
	// WORKING: today     -> gocui
	// WORKING: jump between views
	// WORKING: enter on a team name -> get which team is highlighted

	nt := nbatop.NewNBATop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.Standings = api.GetStandings()
		nt.GetConferenceStandings()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.GamesToday = api.GetGamesToday()
		nt.GetGames()
	}()
	wg.Wait()

	api.GetTeamStats("1610612745")
	// nt.Run()
}
