package nbatop

import (
	"log"

	"github.com/jroimartin/gocui"
)

// standingsDown moves the cursor down one row
func standingsDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		if oy+cy+1 > 32 {
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

// selectGame selects the game at the cursor and displays the game's
// box score in the main table view
func (nt *NBATop) selectGame(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()
		_, oy := v.Origin()

		// Top line contains headers
		idx := cy - 1

		offset := 0
		callingView := v.Name()
		if callingView == "teamgamelog" {
			offset = nt.Views.TeamGameLogView.headerOffset
		} else if callingView == "playerstats" {
			offset = nt.Views.PlayerStatsView.headerOffset
		}
		idx -= offset

		// Adjust team index by scroll distance
		if oy > 0 {
			idx += oy
		}

		// Skip top row
		if idx < 0 {
			return nil
		}

		nt.DrawBoxScore(idx, callingView)
	}
	return nil
}

// standingsUp moves the cursor up one row
func standingsUp(g *gocui.Gui, v *gocui.View) error {
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

// cursorTop moves the cursor to the top of the view
func cursorTop(g *gocui.Gui, v *gocui.View) error {
	if v != nil { // _, oy := v.Origin() cx, cy := v.Cursor()
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}
		v.SetCursor(0, 0)
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

// cursorBottom moves the cursor to the bottom of the view
func cursorBottom(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, maxY := v.Size()
		// _, oy := v.Origin()
		// _, cy := v.Cursor()
		if err := v.SetCursor(0, 31); err != nil {
			if err := v.SetOrigin(0, 31-maxY+2); err != nil {
				return err
			}
		}

		v.SetCursor(0, maxY-1)
	}
	return nil
}

// focusTable sets the most recent table view on top, then focuses it
func (nt *NBATop) focusTable(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetViewOnTop(nt.State.FocusedTableView)
	if err != nil {
		return err
	}

	v, err = g.SetCurrentView(nt.State.FocusedTableView)
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

func (nt *NBATop) tabLeft(g *gocui.Gui, v *gocui.View) error {
	v, err := nt.G.ViewByPosition(nt.State.SidebarWidth+1, 1)
	if err != nil {
		log.Panicln(err)
	}

	if nt.State.FocusedTableView == "teamgamelog" {
		nt.FocusPlayerStats()
	} else if nt.State.FocusedTableView == "boxscore" {
		nt.FocusTeamGameLog()

	} else if nt.State.FocusedTableView == "playerstats" {
		nt.FocusBoxScore()
	} else {
		log.Panicln("tabLeft: focused view not found")
	}

	return nil
}

func (nt *NBATop) tabRight(g *gocui.Gui, v *gocui.View) error {
	v, err := nt.G.ViewByPosition(nt.State.SidebarWidth+1, 1)
	if err != nil {
		log.Panicln(err)
	}

	if nt.State.FocusedTableView == "teamgamelog" {
		if nt.Views.BoxScoreView.drawn {
			nt.FocusBoxScore()
		}
	} else if nt.State.FocusedTableView == "boxscore" {
		if nt.Views.PlayerStatsView.drawn {
			nt.FocusPlayerStats()
		}

	} else if nt.State.FocusedTableView == "playerstats" {
		if nt.Views.TeamGameLogView.drawn {
			nt.FocusTeamGameLog()
		}
	} else {
		log.Panicln("tabRight: focused view not found")
	}

	return nil
}
