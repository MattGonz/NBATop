package main

import (
	"sync"

	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/nbatop"
)

func main() {

	// TODO: continue to iron out structs / split into multiple files
	// TODO: cache standings / games today?
	// TODO: team stats / schedule tabs
	// TODO: enter on a team name to open team schedule / stats

	// WORKING: get games today
	// WORKING: get standings
	// WORKING: standings -> gocui
	// WORKING: today     -> gocui
	// WORKING: jump between views

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

	nt.Run()

}
