package nbatop

import (
	"log"

	"github.com/jroimartin/gocui"
)

// SetKeybindings sets the keybindings for all views
func (nt *NBATop) SetKeybindings() error {
	nt.SetTodayKeybinds()
	nt.SetStandingsKeybinds()
	nt.SetGenericTableKeybinds()
	nt.SetTGLKeybinds()
	nt.SetBoxScoreKeybinds()
	nt.SetPlayerStatsKeybinds()

	if err := nt.G.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	return nil
}
