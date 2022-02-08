package nbatop

import (
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

// selectTeam selects the team at the cursor and displays the team's
// game log in the TeamGameLogView
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

		teamID := nt.Views.TeamGameLogView.GameLogIdxToTeamID[idx]
		teamName := nt.Views.TeamGameLogView.GameLogIdxToTeamName[idx]

		nt.UpdateGameLogData(teamID, teamName)

		err := nt.Views.TeamGameLogView.Draw()
		if err != nil {
			return err
		}

		nt.State.FocusedTableView = "teamgamelog"
		nt.UpdateTableTitle()

		if v.Name() == "standings" {
			nt.State.LastSidebarView = "standings"
		}
	}
	return nil
}

// selectGame selects the game at the cursor and displays the game's
// box score in the main table view
func (nt *NBATop) selectGame(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		idx := cy

		// Adjust team index by scroll distance
		if oy > 0 {
			idx += oy
		}

		line := strings.Split(v.BufferLines()[idx], "\t")
		lineLen := len(line)

		// Ignore lines that aren't games
		if lineLen < 4 {
			return nil
		}

		// Ignore header row
		if line[0] == "GAME DATE" {
			return nil
		}

		gameID := strings.TrimSpace(line[lineLen-3])
		gameDate := strings.TrimSpace(line[lineLen-2])
		matchup := strings.TrimSpace(line[lineLen-1])

		nt.UpdateBoxScoreData(gameID, gameDate, matchup)

		err := nt.Views.BoxScoreView.Draw()
		if err != nil {
			return err
		}

		nt.State.FocusedTableView = "boxscore"
		nt.UpdateTableTitle()
	}
	return nil
}

// selectPlayer selects the player at the cursor and displays the player's
// stats in the PlayerStatsView
func (nt *NBATop) selectPlayer(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		idx := cy

		// Adjust player index by scroll distance
		if oy > 0 {
			idx += oy
		}

		line := strings.Split(v.BufferLines()[idx], "\t")
		lineLen := len(line)

		// Ignore lines that aren't player stats
		if lineLen < 5 {
			return nil
		}

		// Ignore header rows
		if line[0] == "NAME" {
			return nil
		}

		var personID, teamID string

		personID = strings.TrimSpace(line[lineLen-3])
		teamID = strings.TrimSpace(line[lineLen-1])

		// Players that did not start will not have a position
		// This adjusts the line index to account for this
		if personID == "Internal:" || teamID == "" {
			personID = strings.TrimSpace(line[lineLen-2])
			teamID = strings.TrimSpace(line[lineLen-1])
		}

		playerName := nt.State.PersonIDToPlayerName[personID]
		nt.UpdatePlayerStatsData(teamID, personID, playerName)

		err := nt.Views.PlayerStatsView.Draw()
		if err != nil {
			return err
		}

		nt.State.FocusedTableView = "playerstats"
		nt.UpdateTableTitle()
	}
	return nil
}

// cursorTop moves the cursor to the top of the view
func cursorTop(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}
		v.SetCursor(0, 0)
	}
	return nil
}

// cursorBottom moves the cursor to the bottom of the view
func cursorBottom(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, maxY := v.Size()
		bottom := len(v.BufferLines()) - 3

		if err := v.SetCursor(0, bottom); err != nil {
			if err := v.SetOrigin(0, bottom-maxY+2); err != nil {
				return err
			}
		}

		v.SetCursor(0, maxY-1)
	}
	return nil
}

// todayPrev moves the cursor up 3 rows
func todayPrev(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-3); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-3); err != nil {
				return err
			}
		}
	}
	return nil
}

// todayNext moves the cursor down 3 rows
func todayNext(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		// TODO this changes as you scroll down, use the attribute instead
		// the problem is that this function takes in a generic view instead of a TodayView
		gameLines := len(v.BufferLines()) - 6

		if oy+cy+1 > gameLines {
			return nil
		}

		if err := v.SetCursor(cx, cy+3); err != nil {
			if err := v.SetOrigin(ox, oy+3); err != nil {
				return err
			}
		}

	}
	return nil
}

// focusTable sets the most recent table view on top, then focuses it
func (nt *NBATop) focusTable(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetViewOnTop(nt.State.FocusedTableView)
	if err != nil {
		return err
	}

	nt.State.LastSidebarView = v.Name()

	v, err = g.SetCurrentView(nt.State.FocusedTableView)
	if err != nil {
		return err
	}

	return nil
}

// focusSidebar focuses most recent sidebar view
func (nt *NBATop) focusSidebar(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(nt.State.LastSidebarView)

	if err != nil {
		return err
	}
	return nil
}

// cursorUp moves the cursor up one row
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

// cursorDown moves the cursor down one row
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		bufferLines := len(v.BufferLines()) - 2

		if oy+cy+1 > bufferLines {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

// cursorDown moves the cursor down one row
func (nt *NBATop) cursorLeft(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()

		if ox == 0 {
			return nil
		}

		if err := v.SetOrigin(ox-3, oy); err != nil {
			if err := v.SetOrigin(ox-2, oy); err != nil {
				if err := v.SetOrigin(ox-1, oy); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// cursorDown moves the cursor down one row
func (nt *NBATop) cursorRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		var headerIdx int
		var longestHeader string
		name := v.Name()
		viewLengthX, _ := v.Size()

		if name != "teamgamelog" && name != "boxscore" && name != "playerstats" {
			return nil
		}

		if name == "teamgamelog" {
			headerIdx = nt.Views.TeamGameLogView.headerOffset
			longestHeader = v.BufferLines()[headerIdx]
		} else if name == "boxscore" {
			headerIdx = nt.Views.BoxScoreView.headerOffset
			firstHeaders := v.BufferLines()[headerIdx]
			secondHeaders := v.BufferLines()[nt.Views.BoxScoreView.SecondHeadersIdx]
			if len(firstHeaders) > len(secondHeaders) {
				longestHeader = firstHeaders
			} else if len(secondHeaders) > len(firstHeaders) {
				longestHeader = secondHeaders
			}
		} else if name == "playerstats" {
			headerIdx = nt.Views.PlayerStatsView.headerOffset
			longestHeader = v.BufferLines()[headerIdx]
		}

		ox, oy := v.Origin()

		maxX := strings.Index(longestHeader, "Internal:")

		if ox+viewLengthX+2 > maxX {
			return nil
		}

		if err := v.SetOrigin(ox+3, oy); err != nil {
			return err
		}
	}
	return nil
}

// tabLeft focuses the table view to the left of the current tab
func (nt *NBATop) tabLeft(g *gocui.Gui, v *gocui.View) error {
	var focused string

	v, err := nt.G.ViewByPosition(nt.State.SidebarWidth+1, 1)
	if err != nil {
		log.Panicln(err)
	}
	focused = v.Name()

	if focused == "teamgamelog" {
		nt.Views.PlayerStatsView.Focus()
	} else if focused == "boxscore" {
		nt.Views.TeamGameLogView.Focus()
	} else if focused == "playerstats" {
		nt.Views.BoxScoreView.Focus()
	} else {
		log.Panicln("tabLeft: focused view not found")
	}

	nt.UpdateTableTitle()

	return nil
}

// tabRight focuses the table view to the right of the current tab
func (nt *NBATop) tabRight(g *gocui.Gui, v *gocui.View) error {
	var focused string

	v, err := nt.G.ViewByPosition(nt.State.SidebarWidth+1, 1)
	if err != nil {
		log.Panicln(err)
	}
	focused = v.Name()

	if focused == "teamgamelog" {
		nt.Views.BoxScoreView.Focus()
	} else if focused == "boxscore" {
		nt.Views.PlayerStatsView.Focus()
	} else if focused == "playerstats" {
		nt.Views.TeamGameLogView.Focus()
	} else {
		log.Panicln("tabRight: focused view not found")
	}

	nt.UpdateTableTitle()

	return nil
}
