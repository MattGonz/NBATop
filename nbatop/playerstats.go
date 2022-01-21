package nbatop

import (
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type PlayerStatsView struct {
	v            *gocui.View
	name         string
	playerName   string
	headerOffset int
	headers      []string
	rowSet       [][]interface{}
	drawn        bool
}

// NewTeamGameLog creates a new TeamGameLog view
func NewPlayerStatsView() *PlayerStatsView {
	return &PlayerStatsView{
		v:            &gocui.View{},
		name:         "playerstats",
		playerName:   "Player Stats",
		headerOffset: 0,
		headers:      make([]string, 0),
		rowSet:       make([][]interface{}, 0),
		drawn:        false,
	}
}

// FocusTeamGameLog sets the teamgamelog view on top, focuses it
// and changes the title of the table accordingly
func (nt *NBATop) FocusPlayerStats() (*gocui.View, error) {
	_, err := nt.G.SetViewOnTop("playerstats")
	if err != nil {
		return nil, err
	}

	p, err := nt.G.SetCurrentView("playerstats")
	if err != nil {
		return nil, err
	}
	p.Title = " " + nt.Views.TeamGameLogView.teamName + " | " + nt.Views.BoxScoreView.matchup + " " + nt.Views.BoxScoreView.gameDate + " |[" + nt.Views.PlayerStatsView.playerName + "]"
	nt.State.FocusedTableView = p.Name()

	return p, nil
}

// WriteGameLog writes the current team's data to the TeamGameLogView
func (nt *NBATop) WritePlayerStats() error {
	nt.FocusPlayerStats()
	v := nt.Views.PlayerStatsView.v

	// Clear previous team's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 1, '\t', tabwriter.AlignRight)

	name := strings.Fields(nt.Views.PlayerStatsView.playerName)
	firstName := name[0]
	lastName := name[1]

	printFigure(w, firstName+" "+lastName)

	lines := len(v.BufferLines()) - 1
	nt.Views.PlayerStatsView.headerOffset = lines
	v.SetCursor(0, lines)

	// Write headers
	for _, header := range nt.Views.PlayerStatsView.headers {

		// FG_PCT -> FG % etc.
		header = strings.Replace(header, "_", " ", 1)
		header = strings.Replace(header, "PCT", "%", 1)

		fmt.Fprintf(w, "%s\t ", header)
	}

	// Newline after headers to not cut off most recent game
	fmt.Fprintln(w, "")

	// Write data
	for _, row := range nt.Views.PlayerStatsView.rowSet {

		// More compact / readable date (Jan 02, 2006 -> 01-02-2006)
		gameDate, err := time.Parse("Jan 02, 2006", row[3].(string))
		if err != nil {
			return err
		}
		gameDateFormatted := gameDate.Format("01-02-2006")
		row[3] = gameDateFormatted

		for _, col := range row[3:] {
			fmt.Fprint(w, col)
			fmt.Fprint(w, "\t ")
		}
		fmt.Fprintln(w, "")
	}
	w.Flush()
	return nil
}

// DrawPlayerGameLog draws a given Person IDs recent games in a gocui view
func (nt *NBATop) DrawPlayerGameLog(teamID, personID, playerName string) error {
	displayYear := nt.State.CurrentSeason.TeamSitesOnly.DisplayYear
	result := api.GetPlayerGameLog(displayYear, teamID, personID)
	headers := result.ResultSets[0].Headers[3:]
	rowSet := result.ResultSets[0].RowSet

	nt.Views.PlayerStatsView.playerName = playerName
	nt.Views.PlayerStatsView.headers = headers
	nt.Views.PlayerStatsView.rowSet = rowSet

	// TODO write a generalized "CreateIfNotExists" function
	// Check if the view already exists, create view accordingly
	if nt.Views.PlayerStatsView.drawn {
		nt.WritePlayerStats()
	} else {
		if v, err := nt.G.SetView("playerstats", nt.State.SidebarWidth, 0, nt.State.MaxX-1, nt.State.MaxY-1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Highlight = true
			nt.Views.PlayerStatsView.drawn = true

			nt.Views.PlayerStatsView.v = v
			nt.WritePlayerStats()
		}
	}
	return nil
}

// SetPlayerStatsKeybinds sets the keybindings for the PlayerStatsView
func (nt *NBATop) SetPlayerStatsKeybinds() error {
	if err := nt.G.SetKeybinding("playerstats", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'H', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", gocui.KeyEnter, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'A', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'D', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	return nil
}
