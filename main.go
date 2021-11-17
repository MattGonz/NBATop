package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/utils"
)

// cursorDown moves the cursor down one row
func cursorDown(g *gocui.Gui, v *gocui.View) error {
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

// cursorUp moves the cursor up one row
func cursorUp(g *gocui.Gui, v *gocui.View) error {
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

// cursorBottom moves the cursor to the bottom of the view
func cursorBottom(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, maxY := v.Size()
		// _, oy := v.Origin()
		// _, cy := v.Cursor()
		if err := v.SetCursor(0, 31); err != nil {
			if err := v.SetOrigin(0, 31-maxY+1); err != nil {
				return err
			}
		}

		v.SetCursor(0, maxY-1)
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	// if err := g.SetKeybinding("standings", 'h', gocui.ModNone, cursorLeft); err != nil {
	// 	return err
	// }
	if err := g.SetKeybinding("standings", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("standings", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	// if err := g.SetKeybinding("standings", 'l', gocui.ModNone, cursorRight); err != nil {
	// 	return err
	// }
	if err := g.SetKeybinding("", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'G', gocui.ModNone, cursorBottom); err != nil {
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

// conferenceStandings returns the formatted standings for the eastern and western conferences
func conferenceStandings() ([]string, []string) {
	var westernConference []string
	var easternConference []string

	var standings = api.Standings()

	for _, team := range standings.League.Standard.Conference.East {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname

		wins, err := strconv.Atoi(team.Win)
		if err != nil {
			log.Panicln(err)
		}

		var record string
		if wins < 10 {
			record = "    " + team.Win + "-" + team.Loss
		} else {
			record = "   " + team.Win + "-" + team.Loss
		}

		easternConference = append(easternConference, "\t"+name+"\t"+record)
	}

	for _, team := range standings.League.Standard.Conference.West {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		wins, err := strconv.Atoi(team.Win)
		if err != nil {
			log.Panicln(err)
		}

		var record string
		if wins < 10 {
			record = " " + team.Win + "-" + team.Loss
		} else {
			record = team.Win + "-" + team.Loss
		}

		westernConference = append(westernConference, "\t"+name+"\t"+record)
	}
	return westernConference, easternConference
}

// gamesToday returns the formatted games for today
func gamesToday() [][]string {
	var games [][]string

	var gamesToday = api.GamesToday()

	for _, game := range gamesToday.Games {
		homeTeam := game.HTeam
		homeTeamRecord := " (" + homeTeam.Win + "-" + homeTeam.Loss + ")"
		awayTeam := game.VTeam
		awayTeamRecord := "(" + awayTeam.Win + "-" + awayTeam.Loss + ") "

		gameInfo := awayTeamRecord + awayTeam.TriCode + " at " + homeTeam.TriCode + homeTeamRecord
		// spaces := strings.Repeat(" ", len(gameInfo)/3)
		startTime := game.StartTimeEastern

		games = append(games, []string{startTime, gameInfo})
	}
	return games
}

// drawStandings draws the current east and west conference standings in a gocui view
func drawStandings(g *gocui.Gui, west, east []string, length, width, startY int, spaces, today string) error {
	if v, err := g.SetView("standings", 0, startY, width-1, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView("standings")

		// v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Title = "Standings [" + today + "]"
		v.MoveCursor(0, 2, true)
		fmt.Fprint(v, "\tTeam"+spaces+"W-L")

		w := tabwriter.NewWriter(v, 0, 1, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(v, "\n\t\u001b[32mEastern Conference\u001b[0m")

		east = append(east, "\t\u001b[32mWestern Conference\u001b[0m")
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
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
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

func layout(g *gocui.Gui) error {
	today := time.Now().Format("01-02-2006")

	west, east := conferenceStandings()
	gamesToday := gamesToday()

	maxX, maxY := g.Size()
	longestWest := utils.Longest(west)
	longestEast := utils.Longest(east)
	numGames := len(gamesToday)
	spaces := strings.Repeat(" ", utils.Max(longestWest, longestEast)-8)
	standingsLength := utils.Max(len(west)+len(east)+4, maxY-1)
	standingsWidth := utils.Min(utils.Max(longestWest, longestEast)+4, maxX-1)

	drawToday(g, gamesToday, numGames*3, standingsWidth, today)
	drawStandings(g, west, east, standingsLength, standingsWidth, numGames*3+1, spaces, today)

	return nil
}

func main() {

	// TODO: structs / split into multiple files
	// TODO: cache standings / games today to increase performance on keypress
	// TODO: team stats / schedule tabs
	// TODO: enter on a team name to open team schedule / stats
	// TODO: jump between views

	// WORKING: get games today
	// WORKING: get standings
	// WORKING: standings -> gocui
	// WORKING: today     -> gocui

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
