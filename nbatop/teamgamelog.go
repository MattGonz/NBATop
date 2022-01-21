package nbatop

import (
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type TeamGameLogView struct {
	v            *gocui.View
	name         string
	headerOffset int
	teamID       string
	teamName     string
	headers      []string
	rowSet       [][]interface{}
	drawn        bool
}

// NewTeamGameLog creates a new TeamGameLog view
func NewTeamGameLogView() *TeamGameLogView {
	return &TeamGameLogView{
		v:            &gocui.View{},
		name:         "teamgamelog",
		teamID:       "",
		teamName:     "",
		headerOffset: 0,
		headers:      make([]string, 0),
		rowSet:       make([][]interface{}, 0),
		drawn:        false,
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
	t.Title = "[" + nt.Views.TeamGameLogView.teamName + "]| " + nt.Views.BoxScoreView.matchup + " " + nt.Views.BoxScoreView.gameDate + " | " + nt.Views.PlayerStatsView.playerName + " "
	nt.State.FocusedTableView = t.Name()

	return t, nil
}

// WriteGameLog writes the current team's data to the TeamGameLogView
func (nt *NBATop) WriteGameLog() error {
	nt.FocusTeamGameLog()
	v := nt.Views.TeamGameLogView.v

	// Clear previous team's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 1, '\t', tabwriter.AlignRight)

	printFigure(w, nt.Views.TeamGameLogView.teamName)

	nt.Views.TeamGameLogView.headerOffset = len(v.BufferLines()) - 1
	v.SetCursor(0, nt.Views.TeamGameLogView.headerOffset)

	// Write headers
	for _, header := range nt.Views.TeamGameLogView.headers {

		// FG_PCT -> FG % etc.
		header = strings.Replace(header, "_", " ", 1)
		header = strings.Replace(header, "PCT", "%", 1)

		fmt.Fprintf(w, "%s\t ", header)
	}

	// Newline after headers to not cut off most recent game
	fmt.Fprintln(w, "")

	// Write data
	for _, row := range nt.Views.TeamGameLogView.rowSet {

		// More compact / readable date (Jan 02, 2006 -> 01-02-2006)
		gameDate, err := time.Parse("Jan 02, 2006", row[2].(string))
		if err != nil {
			return err
		}
		gameDateFormatted := gameDate.Format("01-02-2006")
		row[2] = gameDateFormatted

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
func (nt *NBATop) DrawTeamGameLog(teamID, teamName string) error {
	result := api.GetTeamGameLog(teamID)
	headers := result.ResultSets[0].Headers[2:] // TeamID and GameID are first 2 headers
	rowSet := result.ResultSets[0].RowSet

	nt.Views.TeamGameLogView.teamID = teamID
	nt.Views.TeamGameLogView.teamName = teamName
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

// selectTeam selects the team at the cursor and displays the team's
// game log in the main table view
func (nt *NBATop) selectTeam(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		// Top 2 lines are headers
		idx := cy - 2

		// Adjust team index by scroll distance
		if oy > 0 {
			idx += oy
		}

		// Skip top row and Western Conference
		if idx < 0 || idx == 15 {
			return nil
		}

		// Adjust for teams after "Western Conference"
		if idx > 15 {
			idx -= 1
		}

		teamID := nt.State.GameLogIdxToTeamID[idx]
		teamName := nt.State.GameLogIdxToTeamName[idx]

		err := nt.DrawTeamGameLog(teamID, teamName)
		if err != nil {
			return err
		}

	}
	return nil
}

func printFigure(w *tabwriter.Writer, content string) {
	// teamNameFigure := figure.NewFigure(content, "standard", true)
	// figure.Write(w, teamNameFigure)
	fmt.Fprintln(w, content)
	fmt.Fprintln(w, "")
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
	if err := nt.G.SetKeybinding("teamgamelog", 'A', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'D', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	return nil
}
