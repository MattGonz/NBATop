package nbatop

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type TeamGameLogView struct {
	v       *gocui.View
	name    string
	team    string
	headers []string
	rowSet  [][]interface{}
	drawn   bool
}

// NewTeamGameLog creates a new TeamGameLog view
func NewTeamGameLogView() *TeamGameLogView {
	return &TeamGameLogView{
		v:       &gocui.View{},
		name:    "teamgamelog",
		team:    "",
		headers: make([]string, 0),
		rowSet:  make([][]interface{}, 0),
		drawn:   false,
	}
}

// FocusTeamGameLog sets the teamgamelog view on top, focuses it
// and changes the title of the table accordingly
func (nt *NBATop) FocusTeamGameLog() (*gocui.View, error) {
	_, err := nt.G.SetViewOnTop("teamgamelog")
	if err != nil {
		return nil, err
	}

	t, err := nt.G.SetCurrentView("teamgamelog")
	if err != nil {
		return nil, err
	}
	t.Title = "[Game Log]| Box Score | Player Stats "

	return t, nil
}

// WriteGameLog writes the current team's data to the TeamGameLogView
func (nt *NBATop) WriteGameLog() error {
	nt.FocusTeamGameLog()
	v := nt.Views.TeamGameLogView.v

	// Clear previous team's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 2, '\t', tabwriter.AlignRight)

	// Write headers
	for _, header := range nt.Views.TeamGameLogView.headers {
		header = strings.Replace(header, "_", " ", 1)
		fmt.Fprintf(w, "%s\t ", header)
	}

	// Newline after headers to not cut off most recent game
	fmt.Fprintln(w, "")

	// Write data
	for _, row := range nt.Views.TeamGameLogView.rowSet {
		for _, col := range row[2:] {
			fmt.Fprint(w, col)
			fmt.Fprint(w, "\t ")
		}
		fmt.Fprintln(w, "")
	}
	w.Flush()
	return nil
}

// DrawTeamGameLog draws the selected Team IDs recent games in a gocui view
func (nt *NBATop) DrawTeamGameLog(id string) error {
	result := api.GetTeamGameLog(id)
	headers := result.ResultSets[0].Headers[2:] // TeamID and GameID are first 2 headers
	rowSet := result.ResultSets[0].RowSet

	nt.Views.TeamGameLogView.team = id
	nt.Views.TeamGameLogView.headers = headers
	nt.Views.TeamGameLogView.rowSet = rowSet

	// TODO write a generalized "CreateIfNotExists" function
	// Check if the view already exists, create view accordingly
	if nt.Views.TeamGameLogView.drawn {
		nt.WriteGameLog()
	} else {
		if v, err := nt.G.SetView("teamgamelog", nt.State.SidebarWidth, 0, nt.State.MaxX-1, nt.State.MaxY-1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Highlight = true
			nt.Views.TeamGameLogView.drawn = true

			nt.Views.TeamGameLogView.v = v
			nt.WriteGameLog()
		}
	}
	return nil
}

// SetTGLKeybinds sets the keybindings for the TeamGameLogView
func (nt *NBATop) SetTGLKeybinds() error {
	if err := nt.G.SetKeybinding("teamgamelog", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'H', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", gocui.KeyEnter, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	return nil
}
