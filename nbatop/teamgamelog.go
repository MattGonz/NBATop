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

type TeamGameLogView struct {
	TableView
	teamID               string
	TeamName             string
	GameLogIdxToTeamID   map[int]string
	GameLogIdxToTeamName map[int]string
}

// NewTeamGameLogView creates a new TeamGameLogView
func (nt *NBATop) NewTeamGameLogView() *TeamGameLogView {
	tgl := &TeamGameLogView{
		TableView: TableView{
			g:            nt.G,
			v:            &gocui.View{},
			name:         "teamgamelog",
			headerOffset: 4,
			headers:      make([]string, 0),
			rowSet:       make([][]interface{}, 0),
			title:        " Game Log ",
			x0:           nt.State.SidebarWidth,
			y0:           0,
			x1:           nt.State.MaxX - 1,
			y1:           nt.State.MaxY - 1,
		},
		teamID:               "",
		TeamName:             "",
		GameLogIdxToTeamID:   nt.State.GameLogIdxToTeamID,
		GameLogIdxToTeamName: nt.State.GameLogIdxToTeamName,
	}
	tgl.TableCreatorWriter = tgl
	return tgl
}

// UpdateGameLogData fetches the given team's data and
// updates the team game log data and title accordingly
func (nt *NBATop) UpdateGameLogData(teamID, teamName string) {
	result := api.GetTeamGameLog(teamID)

	nt.Views.TeamGameLogView.headers = result.ResultSets[0].Headers
	nt.Views.TeamGameLogView.rowSet = result.ResultSets[0].RowSet

	nt.Views.TeamGameLogView.teamID = teamID
	nt.Views.TeamGameLogView.TeamName = teamName
	nt.Views.TeamGameLogView.title = teamName
}

// Create creates the view that holds a team's game log
func (tgl *TeamGameLogView) Create() {
	if v, err := tgl.g.SetView(tgl.name, tgl.x0, tgl.y0, tgl.x1, tgl.y1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		v.Title = tgl.title
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		tgl.v = v
	}
}

// Write writes the current team's data to the TeamGameLogView
func (tgl *TeamGameLogView) Write() {
	tgl.Focus()
	v := tgl.v

	// Clear previous team's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "")
	utils.PrintFigure(w, tgl.TeamName, v)
	fmt.Fprintln(w, "")

	v.SetCursor(0, tgl.headerOffset)

	// Write headers
	for _, header := range tgl.headers[2:] {

		// FG_PCT -> FG % etc.
		header = strings.Replace(header, "_", " ", 1)
		header = strings.Replace(header, "PCT", "%", 1)

		fmt.Fprintf(w, "%s\t ", header)
	}
	utils.BlackPrint(w, "Internal:", true)
	utils.BlackPrint(w, tgl.headers[1], true)
	utils.BlackPrint(w, tgl.headers[2], true)
	utils.BlackPrint(w, tgl.headers[3], true)

	// Newline after headers to not cut off most recent game
	fmt.Fprintln(w, "")

	// Write data
	for _, row := range tgl.rowSet {
		gameID := row[1]
		gameDateUnf := row[2]
		matchup := row[3]

		// More compact / readable date (Jan 02, 2006 -> 01-02-2006)
		gameDate, err := time.Parse("Jan 02, 2006", row[2].(string))
		if err != nil {
			log.Panicln(err)
		}
		gameDateFormatted := gameDate.Format("01-02-2006")
		row[2] = gameDateFormatted

		for _, col := range row[2:] {
			colStr := fmt.Sprintf("%v", col)
			col = strings.Replace(colStr, "<nil>", "..", 1)
			fmt.Fprint(w, col)
			fmt.Fprint(w, "\t ")
		}

		// Write black lines that can be read to get game data
		utils.BlackPrint(w, "Internal:", true)
		utils.BlackPrint(w, gameID, true)
		utils.BlackPrint(w, gameDateUnf, true)
		utils.BlackPrint(w, matchup, false)

		fmt.Fprintln(w, "")
	}
	w.Flush()
}

// SetTGLKeybinds sets the keybindings for the TeamGameLogView
func (nt *NBATop) SetTGLKeybinds() error {
	if err := nt.G.SetKeybinding("teamgamelog", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", 'H', gocui.ModNone, nt.focusSidebar); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", gocui.KeyEnter, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", '[', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", ']', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("teamgamelog", gocui.MouseLeft, gocui.ModNone, nt.selectGame); err != nil {
		return err
	}
	return nil
}
