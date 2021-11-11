package main

import (
	"fmt"
	"log"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		ox, oy := v.Origin()

		if oy+cy+1 > 31 {
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

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("standings", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	return nil
}

// quit exits the GUI
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// longest returns the length of the longest string in the given array
func longest(strs []string) int {
	longest := 0
	for _, str := range strs {
		if len(str) > longest {
			longest = len(str)
		}
	}
	return longest
}

// max returns the largest of the two ints.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the smallest of the two ints.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// conferenceStandings returns the formatted standings for the eastern and western conferences
func conferenceStandings() ([]string, []string) {
	var westernConference []string
	var easternConference []string

	var standings = api.Standings()

	for _, team := range standings.League.Standard.Conference.East {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		record := "   " + team.Win + "-" + team.Loss
		easternConference = append(easternConference, "\t"+name+"\t "+record)
	}

	for _, team := range standings.League.Standard.Conference.West {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		record := team.Win + "-" + team.Loss
		westernConference = append(westernConference, "\t"+name+"\t "+record)
	}
	return westernConference, easternConference
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// colLength := maxX / 2
	// rowHeight = maxY / 2

	today := time.Now().Format("01-02-2006")
	west, east := conferenceStandings()
	maxWest := longest(west)
	maxEast := longest(east)
	width := min(max(maxWest, maxEast)+4, maxX-1)
	length := min(len(west)+len(east)+5, maxY-1)

	if v, err := g.SetView("standings", 0, 2, width, length); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if h, err := g.SetView("header", 0, 0, width, 2); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			h.Title = "Standings [" + today + "]"
			fmt.Fprint(h, "\tTeam")
			spaces := strings.Repeat(" ", max(maxWest, maxEast)-8)
			fmt.Fprint(h, spaces)
			fmt.Fprint(h, "W-L")
		}
		g.SetCurrentView("standings")

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.MoveCursor(0, 2, true)

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

func main() {

	// WORKING: games today

	// var today = api.GamesToday()
	// for _, game := range today.Games {
	// 	homeTeam := game.HTeam
	// 	awayTeam := game.VTeam
	// 	startTime := game.StartTimeEastern
	// 	fmt.Println(startTime)
	// 	fmt.Println(homeTeam.TriCode + " vs " + awayTeam.TriCode)
	// 	fmt.Println(homeTeam.Win + "-" + homeTeam.Loss + "    " + awayTeam.Win + "-" + awayTeam.Loss + "\n")
	// }

	// WORKING: standings

	// var standings = api.Standings()
	// for _, team := range standings.League.Standard.Teams {
	// 	fmt.Println(team.TeamSitesOnly.TeamTricode + " " + team.Win + "-" + team.Loss)
	// }

	// TODO standings -> gocui
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
