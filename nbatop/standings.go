package nbatop

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
)

type StandingsView struct {
	v                 *gocui.View
	name              string
	WesternConference []string
	EasternConference []string
}

// NewStandingsView creates a new standings view
func NewStandingsView() *StandingsView {
	return &StandingsView{
		v:                 &gocui.View{},
		name:              "standings",
		WesternConference: []string{},
		EasternConference: []string{},
	}
}

// focusStandings focuses the standings view and stores the name of the most recently used table view
func (nt *NBATop) focusStandings(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView("standings")

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

// DrawStandings draws the current east and west conference standings in a gocui view
func (nt *NBATop) DrawStandings() error {
	width := nt.State.SidebarWidth
	length := nt.State.SidebarLength
	startY := nt.State.GamesTodayLength

	if v, err := nt.G.SetView("standings", 0, startY, width-1, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nt.Views.StandingsView.v = v
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.Title = "Standings [" + nt.State.Today + "]"

		nt.G.SetCurrentView("standings")
		v.MoveCursor(0, 2, true)

		fmt.Fprint(v, "\tTeam"+strings.Repeat(" ", nt.State.StandingsSpaces)+"W-L")

		w := tabwriter.NewWriter(v, 0, 1, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(v, "\n\t\u001b[33mEastern Conference\u001b[0m")

		nt.Views.StandingsView.EasternConference = append(nt.Views.StandingsView.EasternConference, "\t\u001b[33mWestern Conference\u001b[0m")
		standings := append(nt.Views.StandingsView.EasternConference, nt.Views.StandingsView.WesternConference...)

		for _, team := range standings {
			if _, err := fmt.Fprintln(w, team); err != nil {
				return err
			}
		}
		w.Flush()
	}
	return nil
}

// SetStandingsKeybinds sets the keybindings for the StandingsView
func (nt *NBATop) SetStandingsKeybinds() error {
	if err := nt.G.SetKeybinding("standings", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", gocui.KeyEnter, gocui.ModNone, nt.selectTeam); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'K', gocui.ModNone, focusToday); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'L', gocui.ModNone, nt.focusTable); err != nil {
		return err
	}
	return nil
}
