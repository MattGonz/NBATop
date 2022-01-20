package nbatop

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
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
func focusToday(g *gocui.Gui, v *gocui.View) error {
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
	return nil
}
