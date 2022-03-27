package nbatop

import (
	"fmt"
	"log"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/utils"
)

type PlayerStatsView struct {
	TableView
	playerName  string
	displayYear string
}

// NewPlayerStatsView creates a new PlayerStatsView
func (nt *NBATop) NewPlayerStatsView() *PlayerStatsView {
	ps := &PlayerStatsView{
		TableView: TableView{
			g:            nt.G,
			v:            &gocui.View{},
			name:         "playerstats",
			headerOffset: 4,
			headers:      make([]string, 0),
			rowSet:       make([][]interface{}, 0),
			title:        "Player Stats",
			x0:           nt.State.SidebarWidth,
			y0:           0,
			x1:           nt.State.MaxX - 1,
			y1:           nt.State.MaxY - 1,
		},
		playerName:  "Player Stats",
		displayYear: nt.State.CurrentSeason.TeamSitesOnly.DisplayYear,
	}
	ps.TableCreatorWriter = ps
	return ps
}

// UpdatePlayerStatsData fetches the given player's data and
// updates the player stats data and title accordingly
func (nt *NBATop) UpdatePlayerStatsData(teamID, personID, playerName string) {
	displayYear := nt.Views.PlayerStatsView.displayYear
	result := api.GetPlayerGameLog(displayYear, teamID, personID)

	headers := result.ResultSets[0].Headers
	rowSet := result.ResultSets[0].RowSet

	nt.Views.PlayerStatsView.playerName = playerName
	nt.Views.PlayerStatsView.title = playerName
	nt.Views.PlayerStatsView.headers = headers
	nt.Views.PlayerStatsView.rowSet = rowSet
	nt.Views.PlayerStatsView.title = playerName
}

// Create creates the view that holds a player's stats
func (ps *PlayerStatsView) Create() {
	if v, err := ps.g.SetView(ps.name, ps.x0, ps.y0, ps.x1, ps.y1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		v.Title = ps.title
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		ps.v = v
	}
}

// Write writes the current player's data to the PlayerStatsView
func (ps *PlayerStatsView) Write() {
	ps.Focus()
	v := ps.v

	// Clear previous player's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 1, '\t', tabwriter.AlignRight)

	name := strings.Fields(ps.playerName)
	firstName := name[0]
	lastName := name[1]

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	utils.PrintFigure(w, firstName+" "+lastName, v)
	fmt.Fprintln(w, "")

	v.SetCursor(0, ps.headerOffset)

	// Write headers
	for _, header := range ps.headers[3:26] {

		// FG_PCT -> FG % etc.
		header = strings.Replace(header, "_", " ", 1)
		header = strings.Replace(header, "PCT", "%", 1)
		header = strings.Replace(header, "PLUS MINUS", "+/-", 1)

		fmt.Fprintf(w, "%s\t ", header)
	}
	utils.BlackPrint(w, "Internal:", true)
	utils.BlackPrint(w, ps.headers[2], true)
	utils.BlackPrint(w, ps.headers[3], true)
	utils.BlackPrint(w, ps.headers[4], true)

	// Newline after headers to not cut off most recent game
	fmt.Fprintln(w, "")

	// Write data
	for _, row := range ps.rowSet {
		gameID := row[2]
		gameDateUnf := row[3]
		matchup := row[4]

		row = row[:26] // remove VIDEO_AVAILABLE column

		// More compact / readable date (Jan 02, 2006 -> 01-02-2006)
		gameDate, err := time.Parse("Jan 02, 2006", row[3].(string))
		if err != nil {
			log.Panicln(err)
		}
		gameDateFormatted := gameDate.Format("01-02-2006")
		row[3] = gameDateFormatted

		for _, col := range row[3:] {
			colStr := fmt.Sprintf("%v", col)
			col = strings.Replace(colStr, "<nil>", "..", 1)
			fmt.Fprint(w, col)
			fmt.Fprint(w, "\t ")
		}
		utils.BlackPrint(w, "Internal:", true)
		utils.BlackPrint(w, gameID, true)
		utils.BlackPrint(w, gameDateUnf, true)
		utils.BlackPrint(w, matchup, false)
		fmt.Fprintln(w, "")
	}
	w.Flush()
}

// SetPlayerStatsKeybinds sets the keybindings for the PlayerStatsView
func (nt *NBATop) SetPlayerStatsKeybinds() error {
	if err := nt.G.SetKeybinding("playerstats", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", 'H', gocui.ModNone, nt.focusSidebar); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", gocui.KeyEnter, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", '[', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", ']', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("playerstats", gocui.MouseLeft, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	return nil
}
