package nbatop

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/utils"
)

type BoxScoreView struct {
	TableView
	gameDate             string
	BoxScore             *api.BoxScore
	PersonIDToPlayerName map[string]string
	TeamIDToTeamName     map[string]string
}

// NewBoxScore creates a new BoxScoreView
func (nt *NBATop) NewBoxScoreView() *BoxScoreView {
	bs := &BoxScoreView{
		TableView: TableView{
			g:            nt.G,
			v:            &gocui.View{},
			name:         "boxscore",
			headerOffset: 6,
			headers:      make([]string, 0),
			rowSet:       make([][]interface{}, 0),
			title:        "Box Score",
			x0:           nt.State.SidebarWidth,
			y0:           0,
			x1:           nt.State.MaxX - 1,
			y1:           nt.State.MaxY - 1,
		},
		gameDate:             "",
		BoxScore:             &api.BoxScore{},
		PersonIDToPlayerName: nt.State.PersonIDToPlayerName,
		TeamIDToTeamName:     nt.State.TeamIDToTeamName,
	}
	bs.TableCreatorWriter = bs
	return bs
}

// UpdateBoxScoreData fetches the given game's data and
// updates the box score data and title accordingly
func (nt *NBATop) UpdateBoxScoreData(gameID, gameDate, matchup string) {
	boxScore := api.GetBoxScore(gameDate, gameID)

	nt.Views.BoxScoreView.BoxScore = boxScore

	nt.Views.BoxScoreView.gameDate = gameDate
	nt.Views.BoxScoreView.title = matchup
}

// Create creates the view that holds a game's box score
func (bs *BoxScoreView) Create() {
	if v, err := bs.g.SetView(bs.name, bs.x0, bs.y0, bs.x1, bs.y1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		v.Title = bs.title
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		bs.v = v
	}
}

// Write writes the current game's box score to the BoxScoreView
func (bs *BoxScoreView) Write() {
	bs.Focus()
	v := bs.v

	// Clear previous game's data, if any
	v.Clear()

	w := tabwriter.NewWriter(v, 1, 1, 1, '\t', tabwriter.AlignRight)

	homeTeamPoints := bs.BoxScore.BasicGameData.HTeam.Score
	awayTeamPoints := bs.BoxScore.BasicGameData.VTeam.Score
	isGameOn := bs.BoxScore.BasicGameData.IsGameActivated

	var gameLine string
	if strings.Contains(bs.title, "@") {
		gameLine = fmt.Sprintf("(%s) %s (%s)", awayTeamPoints, bs.title, homeTeamPoints)
	} else {
		gameLine = fmt.Sprintf("(%s) %s (%s)", homeTeamPoints, bs.title, awayTeamPoints)
	}

	gameDate := bs.gameDate

	if isGameOn {
		quarter := strconv.Itoa(bs.BoxScore.BasicGameData.Period.Current)
		time := bs.BoxScore.BasicGameData.Clock
		if time == "" {
			if bs.BoxScore.BasicGameData.Period.IsHalftime {
				time = "Halftime"
			} else {
				time = "End Q" + quarter
			}
			utils.PrintFigure(w, time, v)
		} else {
			utils.PrintFigure(w, "(Q"+quarter+" - "+time+")", v)
		}
	} else {
		fmt.Fprintln(w, "")
	}
	utils.PrintFigure(w, gameLine, v)
	utils.PrintFigure(w, gameDate, v)
	fmt.Fprintln(w, "")

	v.SetCursor(0, bs.headerOffset)

	// HACK for separating teams, TODO personID to player (not just name)
	awayTeamID := bs.BoxScore.Stats.ActivePlayers[0].TeamID
	awayTeamName := bs.TeamIDToTeamName[awayTeamID]

	utils.PrintName(w, awayTeamName, v)
	printBoxScoreHeaders(w)
	homePrinted := false

	for _, player := range bs.BoxScore.Stats.ActivePlayers {
		// TODO see if we can check sortkeys here and highlight stat leaders

		// players that aren't found in the active players list will not be displayed
		if bs.PersonIDToPlayerName[player.PersonID] == "" {
			continue
		}

		// HACK for separating teams, TODO personID to player (not just name)
		if player.TeamID != awayTeamID && !homePrinted {
			homeTeamName := bs.TeamIDToTeamName[player.TeamID]
			// fmt.Fprintln(w, "")
			// fmt.Fprintf(w, "\u001b[33m%s\u001b[0m\n", homeTeamName)
			utils.PrintName(w, homeTeamName, v)
			printBoxScoreHeaders(w)
			homePrinted = true
		}

		fmt.Fprintf(w, "%s\t", bs.PersonIDToPlayerName[player.PersonID])
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
		utils.BlackPrint(w, "Internal:", true)
		utils.BlackPrint(w, player.PersonID, true)
		utils.BlackPrint(w, player.TeamID, false)
		// fmt.Fprintf(w, "%s\t", player.Dnp)

		fmt.Fprintln(w, "")
	}
	w.Flush()
}

// printBoxScoreHeaders writes the column names for the BoxScoreView,
// since they are not returned as headers from the API
func printBoxScoreHeaders(w *tabwriter.Writer) {
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
	utils.BlackPrint(w, "Internal:", true)
	utils.BlackPrint(w, "PID", true)
	utils.BlackPrint(w, "TID", false)
	// fmt.Fprintf(w, "DNP\t")

	fmt.Fprintln(w, "")
}

// SetBoxScoreKeybinds sets the keybindings for the BoxScoreView
func (nt *NBATop) SetBoxScoreKeybinds() error {
	if err := nt.G.SetKeybinding("boxscore", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", 'H', gocui.ModNone, nt.focusSidebar); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", gocui.KeyEnter, gocui.ModNone, nt.selectPlayer); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", '[', gocui.ModNone, nt.tabLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("boxscore", ']', gocui.ModNone, nt.tabRight); err != nil {
		return err
	}
	return nil
}
