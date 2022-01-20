package nbatop

import (
	"github.com/jroimartin/gocui"
)

type EmptyTableView struct {
	v    *gocui.View
	name string
}

// NewTableView creates a new table view
func NewTableView() *EmptyTableView {
	return &EmptyTableView{
		v:    &gocui.View{},
		name: "table",
	}
}

// DrawTable draws the initial table view
func (nt *NBATop) DrawInitialTable() error {
	x0 := nt.State.SidebarWidth
	if v, err := nt.G.SetView("table", x0, 0, nt.State.MaxX-1, nt.State.MaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Game Log | Box Score | Player Stats "
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		nt.Views.TableView.v = v
	}
	return nil
}

// SetBoxScoreKeybinds sets the keybindings for the empty table view
func (nt *NBATop) SetTableKeybinds() error {
	if err := nt.G.SetKeybinding("table", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("table", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("table", 'H', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	return nil
}
