package nbatop

import (
	"fmt"
	"io"
	"log"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type BoxScoreView struct {
	v            *gocui.View
	name         string
	headerOffset int
	gameDate     string
	matchup      string
	boxScore     *api.BoxScore
	drawn        bool
}

// NewBoxScore creates a new box score view
func NewBoxScoreView() *BoxScoreView {
	return &BoxScoreView{
		v:            &gocui.View{},
		name:         "boxscore",
		matchup:      " Box Score ",
		headerOffset: 0,
		boxScore:     &api.BoxScore{},
		drawn:        false,
	}
}

// FocusBoxScore sets the boxscore view on top, focuses it
// and changes the title of the table accordingly
func (nt *NBATop) FocusBoxScore() {
	_, err := nt.G.SetViewOnTop("boxscore")
	if err != nil {
		log.Panicln(err)
	}

	t, err := nt.G.SetCurrentView("boxscore")
	if err != nil {
		log.Panicln(err)
	}

	t.Title = " " + nt.Views.TeamGameLogView.teamName + " |[" + nt.Views.BoxScoreView.matchup + " " + nt.Views.BoxScoreView.gameDate + "]| " + nt.Views.PlayerStatsView.playerName + " "
	nt.State.FocusedTableView = t.Name()
}

// WriteBoxScore writes the current game's box score to the BoxScoreView
func (nt *NBATop) WriteBoxScore() error {
	nt.FocusBoxScore()
	v := nt.Views.BoxScoreView.v

	// Clear previous game's data, if any
	v.Clear()

	// TODO think through window-based padding throughout
	padding := 0
	if nt.State.MaxX < 150 {
		padding = 1
	} else {
		padding = 2
	}

	w := tabwriter.NewWriter(v, 0, 1, padding, '\t', tabwriter.AlignRight)

	homeTeamPoints := nt.Views.BoxScoreView.boxScore.BasicGameData.HTeam.Score
	awayTeamPoints := nt.Views.BoxScoreView.boxScore.BasicGameData.VTeam.Score

	gameLine := fmt.Sprintf("(%s) %s (%s)", awayTeamPoints, nt.Views.BoxScoreView.matchup, homeTeamPoints)
	gameDate := nt.Views.BoxScoreView.gameDate

	printFigure(w, gameLine+"\n"+gameDate)

	nt.Views.BoxScoreView.headerOffset = len(v.BufferLines()) - 1
	v.SetCursor(0, nt.Views.BoxScoreView.headerOffset)

	printBoxScoreHeaders(w)

	// HACK for separating teams, TODO personID to player (not just name)
	// awayTeamID := nt.Views.BoxScoreView.boxScore.Stats.ActivePlayers[0].TeamID
	// awayTeamName := nt.State.TeamIDToTeamName[awayTeamID]
	// fmt.Fprintln(w, awayTeamName)
	// homePrinted := false

	for _, player := range nt.Views.BoxScoreView.boxScore.Stats.ActivePlayers {
		// TODO see if we can check sortkeys here and highlight stat leaders

		// players that aren't found in the active players list will not be displayed
		if nt.State.PersonIDToPlayerName[player.PersonID] == "" {
			continue
		}

		// HACK for separating teams, TODO personID to player (not just name)
		// if player.TeamID != awayTeamID && !homePrinted {
		// 	homeTeamName := nt.State.TeamIDToTeamName[player.TeamID]
		// 	fmt.Fprintln(w, "")
		// 	fmt.Fprintln(w, homeTeamName)
		// 	homePrinted = true
		// }

		fmt.Fprintf(w, "%s\t", nt.State.PersonIDToPlayerName[player.PersonID])
		fmt.Fprintf(w, "%s\t", player.Pos)
		fmt.Fprintf(w, "%s\t", player.Min)
		fmt.Fprintf(w, "%s\t", player.Fgm)
		fmt.Fprintf(w, "%s\t", player.Fga)
		fmt.Fprintf(w, "%s\t", player.Fgp)
		fmt.Fprintf(w, "%s\t", player.Tpm)
		fmt.Fprintf(w, "%s\t", player.Tpa)
		fmt.Fprintf(w, "%s\t", player.Tpp)
		fmt.Fprintf(w, "%s\t", player.Ftm)
		fmt.Fprintf(w, "%s\t", player.Fta)
		fmt.Fprintf(w, "%s\t", player.Ftp)
		fmt.Fprintf(w, "%s\t", player.OffReb)
		fmt.Fprintf(w, "%s\t", player.DefReb)
		fmt.Fprintf(w, "%s\t", player.TotReb)
		fmt.Fprintf(w, "%s\t", player.Assists)
		fmt.Fprintf(w, "%s\t", player.Steals)
		fmt.Fprintf(w, "%s\t", player.Blocks)
		fmt.Fprintf(w, "%s\t", player.Turnovers)
		fmt.Fprintf(w, "%s\t", player.PFouls)
		fmt.Fprintf(w, "%s\t", player.Points)
		fmt.Fprintf(w, "%s\t", player.PlusMinus)
		// fmt.Fprintf(w, "%s\t", player.Dnp)

		fmt.Fprintln(w, "")
	}
	w.Flush()
	return nil
}

// DrawBoxScore draws the box score of the game at the given TeamGameLog game index
func (nt *NBATop) DrawBoxScore(idx int, callingView string) error {
	// TODO game score

	var gameID string
	var gameDate string
	var matchup string

	if callingView == "teamgamelog" {
		gameID = nt.Views.TeamGameLogView.rowSet[idx][1].(string)
		gameDate = nt.Views.TeamGameLogView.rowSet[idx][2].(string)
		matchup = nt.Views.TeamGameLogView.rowSet[idx][3].(string)
	} else if callingView == "playerstats" {
		gameID = nt.Views.PlayerStatsView.rowSet[idx][2].(string)
		gameDate = nt.Views.PlayerStatsView.rowSet[idx][3].(string)
		matchup = nt.Views.PlayerStatsView.rowSet[idx][4].(string)
	}

	boxScore := api.GetBoxScore(gameDate, gameID)
	nt.Views.BoxScoreView.gameDate = gameDate
	nt.Views.BoxScoreView.matchup = matchup
	nt.Views.BoxScoreView.boxScore = boxScore

	// TODO write a generalized "CreateIfNotExists" function
	// Check if the view already exists, create view accordingly
	if nt.Views.BoxScoreView.drawn {
		nt.WriteBoxScore()
	} else {
		if v, err := nt.G.SetView("boxscore", nt.State.SidebarWidth, 0, nt.State.MaxX-1, nt.State.MaxY-1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Highlight = true
			nt.Views.BoxScoreView.drawn = true
			nt.Views.BoxScoreView.v = v
			nt.WriteBoxScore()
		}
	}
	return nil
}

// printBoxScoreHeaders writes the column names for the BoxScoreView
func printBoxScoreHeaders(w io.Writer) {
	fmt.Fprintf(w, "NAME\t")
	fmt.Fprintf(w, "POS\t")
	fmt.Fprintf(w, "MIN\t")
	fmt.Fprintf(w, "FGM\t")
	fmt.Fprintf(w, "FGA\t")
	fmt.Fprintf(w, "FG%%\t")
	fmt.Fprintf(w, "3PM\t")
	fmt.Fprintf(w, "3PA\t")
	fmt.Fprintf(w, "3P%%\t")
	fmt.Fprintf(w, "FTM\t")
	fmt.Fprintf(w, "FTA\t")
	fmt.Fprintf(w, "FT%%\t")
	fmt.Fprintf(w, "OREB\t")
	fmt.Fprintf(w, "DREB\t")
	fmt.Fprintf(w, "REB\t")
	fmt.Fprintf(w, "AST\t")
	fmt.Fprintf(w, "STL\t")
	fmt.Fprintf(w, "BLK\t")
	fmt.Fprintf(w, "TO\t")
	fmt.Fprintf(w, "PF\t")
	fmt.Fprintf(w, "PTS\t")
	fmt.Fprintf(w, "+/-\t")
	// fmt.Fprintf(w, "DNP\t")

	fmt.Fprintln(w, "")
}

// selectPlayer selects the player at the cursor and displays the player's
// stats in the main table view
func (nt *NBATop) selectPlayer(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		// Top line contains headers
		idx := cy - 1

		offset := nt.Views.BoxScoreView.headerOffset
		idx -= offset

		// Adjust player index by scroll distance
		if oy > 0 {
			idx += oy
		}

		// Skip top row
		if idx < 0 {
			return nil
		}

		player := nt.Views.BoxScoreView.boxScore.Stats.ActivePlayers[idx]
		teamID := player.TeamID
		personID := player.PersonID
		playerName := nt.State.PersonIDToPlayerName[personID]
		nt.DrawPlayerGameLog(teamID, personID, playerName)
	}
	return nil
}

// SetBoxScoreKeybinds sets the keybindings for the BoxScoreView
func (nt *NBATop) SetBoxScoreKeybinds() error {
	if err := nt.G.SetKeybinding("boxscore", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'H', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", gocui.KeyEnter, gocui.ModNone, nt.selectPlayer); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'A', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'D', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	return nil
}
