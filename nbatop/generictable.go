package nbatop

import (
	"log"

	"github.com/jroimartin/gocui"
)

// GenericTableView is the empty table view that is drawn
// before any data populates the table
type GenericTableView struct {
	TableView
}

// NewGenericTableView creates a new empty table view
func (nt *NBATop) NewGenericTableView() *GenericTableView {
	tv := &GenericTableView{
		TableView: TableView{
			g:     nt.G,
			v:     &gocui.View{},
			name:  "table",
			title: " Game Log | Box Score | Player Stats ",
			x0:    nt.State.SidebarWidth,
			y0:    0,
			x1:    nt.State.MaxX - 1,
			y1:    nt.State.MaxY - 1,
		},
	}
	tv.TableCreatorWriter = tv
	return tv
}

// Create creates the view that holds the generic table
func (gtv *GenericTableView) Create() {
	if v, err := gtv.g.SetView(gtv.name, gtv.x0, gtv.y0, gtv.x1, gtv.y1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}
		v.Title = gtv.title
		v.Highlight = true
		v.SelFgColor = gocui.ColorGreen
		gtv.v = v
	}

}

// Write writes nothing to the generic table view
// (required by the TableCreatorWriter interface)
func (gtv *GenericTableView) Write() {
}

// DrawGenericTable draws the generic table
func (nt *NBATop) DrawGenericTable() error {
	gtv := nt.NewGenericTableView()

	err := gtv.Draw()
	if err != nil {
		return err
	}
	return nil
}

// SetGenericTableKeybinds sets the keybindings for the empty table view
func (nt *NBATop) SetGenericTableKeybinds() error {
	if err := nt.G.SetKeybinding("table", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("table", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("table", 'H', gocui.ModNone, nt.focusStandings); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("", 'h', gocui.ModNone, nt.cursorLeft); err != nil {
		return err
	}
	if err := nt.G.SetKeybinding("", 'l', gocui.ModNone, nt.cursorRight); err != nil {
		return err
	}
	return nil
}
