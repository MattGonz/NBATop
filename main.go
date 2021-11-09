package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

func main() {

	// WORKING: games today

	// var today = api.GamesToday()
	// for _, game := range today.Games {
	// 	homeTeam := game.HTeam
	// 	awayTeam := game.VTeam
	// 	startTime := game.StartTimeEastern
	// 	fmt.Println(startTime)
	// 	fmt.Println(homeTeam.TriCode + " vs " + awayTeam.TriCode)
	// 	fmt.Println(homeTeam.Win + "-" + homeTeam.Loss + "    " + awayTeam.Win + "-" + awayTeam.Loss + "\n")
	// }

	// WORKING: standings

	// var standings = api.Standings()
	// for _, team := range standings.League.Standard.Teams {
	// 	fmt.Println(team.TeamSitesOnly.TeamTricode + " " + team.Win + "-" + team.Loss)
	// }

	// TODO standings -> gocui

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, maxX/2, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		var standingsData []string
		var standings = api.Standings()
		for _, team := range standings.League.Standard.Teams {
			city := team.TeamSitesOnly.TeamName
			name := team.TeamSitesOnly.TeamNickname
			wins := team.Win
			losses := team.Loss
			standingsData = append(standingsData, city+" "+name+" "+wins+"-"+losses)
		}

		today := time.Now().Format("01-02-2006")
		v.Title = "Standings [" + today + "]"
		for _, team := range standingsData {
			fmt.Fprintln(v, team)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
