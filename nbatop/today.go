package nbatop

import (
	"fmt"
	"strconv"
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

// focusToday focuses the TodayView
func (nt *NBATop) focusToday(g *gocui.Gui, v *gocui.View) error {
	if !nt.State.DrawGamesToday {
		return nil
	}
	_, err := g.SetCurrentView("today")
	if err != nil {
		return err
	}
	return nil
}

// DrawToday draws today's games in a gocui view
func (nt *NBATop) DrawToday() error {
	today := nt.Views.TodayView
	x1 := nt.State.SidebarWidth - 1
	y1 := nt.State.GamesTodayLength - 1
	if v, err := nt.G.SetView("today", 0, 0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		today.v = v
		nt.G.SetCurrentView("today")

		v.SelFgColor = gocui.ColorGreen
		v.Title = strconv.Itoa(len(nt.Views.TodayView.Games)) + " Games Today [" + nt.State.Today + "]"

		for _, game := range today.Games {
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

// selectGameToday selects a game in the TodayView and displays it in the BoxScoreView
func (nt *NBATop) selectGameToday(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		idx := cy

		// Adjust team index by scroll distance
		if oy > 0 {
			idx += oy
		}

		gameInfo := nt.State.GamesTodayIdxToGame[idx]
		gameID := gameInfo[0]
		gameDate := gameInfo[1]
		matchup := gameInfo[2]

		nt.UpdateBoxScoreData(gameID, gameDate, matchup)

		// Games that haven't started yet will have no active players
		if nt.Views.BoxScoreView.BoxScore.Stats.ActivePlayers == nil {
			return nil
		}

		err := nt.Views.BoxScoreView.Draw()
		if err != nil {
			return err
		}

		nt.State.FocusedTableView = "boxscore"
		nt.UpdateTableTitle()

		nt.State.LastSidebarView = "today"
	}
	return nil
}

// refreshToday refreshes the games in the TodayView
func (nt *NBATop) refreshToday(g *gocui.Gui, v *gocui.View) error {
	x1 := nt.State.SidebarWidth - 1

	v.Clear()

	nt.State.GamesToday = api.GetGamesToday()
	nt.FormatGamesToday()

	today := nt.Views.TodayView

	for _, game := range today.Games {
		startTime := game[0]
		gameInfo := game[1]
		timeSpaces := strings.Repeat(" ", ((x1-2)-len(startTime))/2)
		gameSpaces := strings.Repeat(" ", ((x1-2)-len(gameInfo))/2)
		fmt.Fprintln(v, timeSpaces+startTime)
		fmt.Fprintln(v, gameSpaces+gameInfo+"\n")
	}
	return nil
}

// SetTodayKeybinds sets the keybindings for the TodayView
func (nt *NBATop) SetTodayKeybinds() error {
	if err := nt.G.SetKeybinding("today", 'j', gocui.ModNone, todayNext); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'k', gocui.ModNone, todayPrev); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'J', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'L', gocui.ModNone, nt.focusTable); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", gocui.KeyEnter, gocui.ModNone, nt.selectGameToday); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'r', gocui.ModNone, nt.refreshToday); err != nil {
		return err
	}
	return nil
}
