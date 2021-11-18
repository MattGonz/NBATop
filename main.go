package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/ui"
	"github.com/mattgonz/nbatop/utils"
)

type NBATop struct {
	g         *gocui.Gui
	standings *api.NBAStandings
	today     *api.NBAToday
	views     *ui.Views
}

// standingsDown moves the cursor down one row
func standingsDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		if oy+cy+1 > 32 {
			return nil
		}

		// skip "Western Conference" (performance hit)
		// else if oy+cy+1 == 16 {
		// 	if err := v.SetCursor(cx, cy+2); err != nil {
		// 		return err
		// 	}
		// 	return nil
		// }

		if err := v.SetCursor(cx, cy+1); err != nil {
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

// standingsUp moves the cursor up one row
func standingsUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		// skip "Western Conference" (performance hit)
		// if oy+cy+1 == 18 {
		// 	if err := v.SetCursor(cx, cy-2); err != nil {
		// 		return err
		// 	}
		// 	return nil
		// }

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
		v.SetCursor(0, 2)
	}
	return nil
}

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

func focusToday(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView("today")
	if err != nil {
		return err
	}
	return nil
}

func focusStandings(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView("standings")
	if err != nil {
		return err
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	// if err := g.SetKeybinding("standings", 'h', gocui.ModNone, cursorLeft); err != nil {
	// 	return err
	// }
	if err := g.SetKeybinding("standings", 'j', gocui.ModNone, standingsDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("standings", 'k', gocui.ModNone, standingsUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("today", 'j', gocui.ModNone, todayNext); err != nil {
		return err
	}
	if err := g.SetKeybinding("today", 'k', gocui.ModNone, todayPrev); err != nil {
		return err
	}
	// if err := g.SetKeybinding("standings", 'l', gocui.ModNone, cursorRight); err != nil {
	// 	return err
	// }
	if err := g.SetKeybinding("standings", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := g.SetKeybinding("standings", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := g.SetKeybinding("standings", 'K', gocui.ModNone, focusToday); err != nil {
		return err
	}
	if err := g.SetKeybinding("today", 'J', gocui.ModNone, focusStandings); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	return nil
}

// quit exits the GUI
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// drawStandings draws the current east and west conference standings in a gocui view
func drawStandings(g *gocui.Gui, west, east []string, length, width, startY int, spaces, today string) error {
	if v, err := g.SetView("standings", 0, startY, width-1, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView("standings")

		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.Title = "Standings [" + today + "]"
		v.MoveCursor(0, 2, true)
		fmt.Fprint(v, "\tTeam"+spaces+"W-L")

		w := tabwriter.NewWriter(v, 0, 1, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(v, "\n\t\u001b[33mEastern Conference\u001b[0m")

		east = append(east, "\t\u001b[33mWestern Conference\u001b[0m")
		// east = append(east, "\tTeam\tRecord")
		standings := append(east, west...)

		for _, team := range standings {
			if _, err := fmt.Fprintln(w, team); err != nil {
				return err
			}
		}
		w.Flush()
	}
	return nil
}

// drawToday draws today's games in a gocui view
func drawToday(g *gocui.Gui, games [][]string, length, width int, today string) error {
	if v, err := g.SetView("today", 0, 0, width-1, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView("today")

		// v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		v.Title = "Games [" + today + "]"

		for _, game := range games {
			startTime := game[0]
			gameInfo := game[1]
			timeSpaces := strings.Repeat(" ", ((width-2)-len(startTime))/2)
			gameSpaces := strings.Repeat(" ", ((width-2)-len(gameInfo))/2)
			fmt.Fprintln(v, timeSpaces+startTime)
			fmt.Fprintln(v, gameSpaces+gameInfo+"\n")
		}
	}
	return nil
}

func (nbatop *NBATop) layout(g *gocui.Gui) error {
	today := time.Now().Format("01-02-2006")
	maxX, maxY := nbatop.g.Size()

	// Today
	gamesToday := nbatop.views.Today.Games
	numGames := len(gamesToday)
	todayLength := utils.Min(maxY/2, numGames*3)

	// Standings
	west := nbatop.views.Standings.WesternConference
	east := nbatop.views.Standings.EasternConference
	longestWest := utils.Longest(west)
	longestEast := utils.Longest(east)
	standingsLength := utils.Min(len(west)+len(east)+todayLength, maxY-1)
	standingsWidth := utils.Min(utils.Max(longestWest, longestEast)+4, maxX-1)
	spaces := strings.Repeat(" ", utils.Max(longestWest, longestEast)-8)

	drawToday(g, gamesToday, todayLength, standingsWidth, today)
	drawStandings(g, west, east, standingsLength, standingsWidth, todayLength+1, spaces, today)

	return nil
}

func main() {

	// TODO: iron out structs / split into multiple files
	// TODO: cache standings / games today?
	// TODO: team stats / schedule tabs
	// TODO: enter on a team name to open team schedule / stats

	// WORKING: get games today
	// WORKING: get standings
	// WORKING: standings -> gocui
	// WORKING: today     -> gocui
	// WORKING: jump between views

	// Create GUI
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	nbatop := NBATop{
		g:         g,
		standings: &api.NBAStandings{},
		today:     &api.NBAToday{},
		views: &ui.Views{
			Standings: ui.NewStandingsView(g),
			Today:     ui.NewTodayView(g),
		},
	}
	nbatop.g = g
	nbatop.g.Cursor = true
	nbatop.g.Mouse = true
	nbatop.g.Highlight = true
	nbatop.g.SelFgColor = gocui.ColorGreen

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		nbatop.standings = api.GetStandings()
		nbatop.views.Standings.GetConferenceStandings(nbatop.standings)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		nbatop.today = api.GetGamesToday()
		nbatop.views.Today.GetGames(nbatop.today)
	}()
	wg.Wait()

	nbatop.g.SetManagerFunc(nbatop.layout)

	if err := keybindings(nbatop.g); err != nil {
		log.Panicln(err)
	}

	if err := nbatop.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
