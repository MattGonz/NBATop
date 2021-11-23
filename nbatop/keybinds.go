package nbatop

import (
	"log"

	"github.com/jroimartin/gocui"
)

func (nt *NBATop) SetKeybindings() error {
	// if err := nt.G.SetKeybinding("standings", 'h', gocui.ModNone, cursorLeft); err != nil {
	// 	return err
	// }
	if err := nt.G.SetKeybinding("standings", 'j', gocui.ModNone, standingsDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'k', gocui.ModNone, standingsUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", gocui.KeyEnter, gocui.ModNone, nt.selectTeam); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'j', gocui.ModNone, todayNext); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'k', gocui.ModNone, todayPrev); err != nil {
		return err
	}
	// if err := nt.G.SetKeybinding("standings", 'l', gocui.ModNone, cursorRight); err != nil {
	// 	return err
	// }
	if err := nt.G.SetKeybinding("standings", 'g', gocui.ModNone, cursorTop); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'G', gocui.ModNone, cursorBottom); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("standings", 'K', gocui.ModNone, focusToday); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("today", 'J', gocui.ModNone, focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	return nil
}
