package nbatop

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type TodayView struct {
	v        *gocui.View
	name     string
	Games    [][]string
	NumLines int
}

// NewTodayView creates a new TodayView
func NewTodayView() *TodayView {
	return &TodayView{
		v:     &gocui.View{},
		name:  "today",
		Games: [][]string{},
	}
}

// gamesToday returns the formatted games for today
func (t *TodayView) GetGames(gamesToday *api.NBAToday) {
	var games [][]string

	for _, game := range gamesToday.Games {
		homeTeam := game.HTeam
		homeTeamRecord := " (" + homeTeam.Win + "-" + homeTeam.Loss + ")"
		awayTeam := game.VTeam
		awayTeamRecord := "(" + awayTeam.Win + "-" + awayTeam.Loss + ") "

		gameInfo := awayTeamRecord + awayTeam.TriCode + " at " + homeTeam.TriCode + homeTeamRecord
		startTime := game.StartTimeEastern

		games = append(games, []string{startTime, gameInfo})
	}
	t.Games = games
	t.NumLines = len(games) * 3
}

// DrawToday draws today's games in a gocui view
func (nt *NBATop) DrawToday() error {
	t := nt.Views.Today
	x1 := nt.State.SidebarWidth - 1
	y1 := nt.State.GamesTodayLength - 1
	if v, err := nt.G.SetView("today", 0, 0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		t.v = v
		nt.G.SetCurrentView("today")

		// v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.Title = "Games [" + nt.State.Today + "]"

		for _, game := range t.Games {
			startTime := game[0]
			gameInfo := game[1]
			timeSpaces := strings.Repeat(" ", ((x1-2)-len(startTime))/2)
			gameSpaces := strings.Repeat(" ", ((x1-2)-len(gameInfo))/2)
			fmt.Fprintln(v, timeSpaces+startTime)
			fmt.Fprintln(v, gameSpaces+gameInfo+"\n")
		}
	}
	return nil
}
