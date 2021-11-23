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

// DrawStandings draws the current east and west conference standings in a gocui view
func (nt *NBATop) DrawStandings() error {
	width := nt.State.SidebarWidth
	length := nt.State.SidebarLength
	startY := nt.State.GamesTodayLength
	if v, err := nt.G.SetView("standings", 0, startY, width-1, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		nt.G.SetCurrentView("standings")

		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.Title = "Standings [" + nt.State.Today + "]"
		v.MoveCursor(0, 2, true)
		fmt.Fprint(v, "\tTeam"+strings.Repeat(" ", nt.State.StandingsSpaces)+"W-L")

		w := tabwriter.NewWriter(v, 0, 1, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(v, "\n\t\u001b[33mEastern Conference\u001b[0m")

		nt.Views.Standings.EasternConference = append(nt.Views.Standings.EasternConference, "\t\u001b[33mWestern Conference\u001b[0m")
		standings := append(nt.Views.Standings.EasternConference, nt.Views.Standings.WesternConference...)

		for _, team := range standings {
			if _, err := fmt.Fprintln(w, team); err != nil {
				return err
			}
		}
		w.Flush()
	}
	return nil
}
