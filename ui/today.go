package ui

import (
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
func NewTodayView(g *gocui.Gui) *TodayView {
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
		// spaces := strings.Repeat(" ", len(gameInfo)/3)
		startTime := game.StartTimeEastern

		games = append(games, []string{startTime, gameInfo})
	}
	t.Games = games
	t.NumLines = len(games)
}
