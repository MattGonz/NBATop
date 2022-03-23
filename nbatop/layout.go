package nbatop

import (
	"github.com/jroimartin/gocui"
)

// layout draws the inital layout of the application
func (nt *NBATop) layout(g *gocui.Gui) error {
	var err error

	if nt.State.DrawGamesToday {
		err = nt.DrawToday()
		if err != nil {
			return err
		}
	}
	err = nt.DrawStandings()
	if err != nil {
		return err
	}
	err = nt.DrawGenericTable()
	if err != nil {
		return err
	}

	return nil
}
