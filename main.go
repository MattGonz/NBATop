package main

import (
	"sync"

	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/nbatop"
)

func main() {

	nt := nbatop.NewNBATop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.CurrentSeason = api.GetCurrentSeason()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.ActivePlayers = api.GetPlayers()
		nt.MapPlayerIDs()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.Standings = api.GetStandings()
		nt.FormatConferenceStandings()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nt.State.GamesToday = api.GetGamesToday()
		nt.FormatGamesToday()
	}()
	wg.Wait()

	nt.Run()
}
