package nbatop

import (
	"log"
	"text/tabwriter"

	"github.com/jroimartin/gocui"
)

type TableCreatorWriter interface {
	Write()
	WriteContents()
	Create()
	CreateView()
}

type TableView struct {
	TableCreatorWriter
	g              *gocui.Gui
	v              *gocui.View
	t              *tabwriter.Writer
	name           string
	headerOffset   int
	headers        []string
	rowSet         [][]interface{}
	title          string
	x0, y0, x1, y1 int
}

// Focus sets a table view on top and focuses it
func (tv *TableView) Focus() {
	_, err := tv.g.SetViewOnTop(tv.name)
	if err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}

	_, err = tv.g.SetCurrentView(tv.name)
	if err != nil && err != gocui.ErrUnknownView {
		log.Panicln(err)
	}
}

// WriteContents writes the contents of the table to the table view
func (tv *TableView) WriteContents() {
	tv.Write()
}

// CreateView creates a table view
func (tv *TableView) CreateView() {
	tv.Create()
}

// Draw creates a table view if it doesn't exist and
// writes the table's contents
func (tv *TableView) Draw() error {
	if view, _ := tv.g.View(tv.name); view != nil {
		tv.WriteContents()
	} else {
		tv.CreateView()
		tv.WriteContents()
	}
	return nil
}

// UpdateTableTitle updates the title of the table view
// to show which tab is currently selected
func (nt *NBATop) UpdateTableTitle() {
	bracketed := func(title string) string { return "[" + title + "]" }
	spaced := func(title string) string { return " " + title + " " }

	focusedView, err := nt.G.ViewByPosition(nt.State.SidebarWidth+1, 1)
	if err != nil {
		return
	}
	focusedViewName := focusedView.Name()

	tglTitle := nt.Views.TeamGameLogView.title
	bsTitle := nt.Views.BoxScoreView.title
	psTitle := nt.Views.PlayerStatsView.title

	if focusedViewName == "teamgamelog" {
		focusedView.Title = bracketed(tglTitle) + spaced(bsTitle) + spaced(psTitle)
	}
	if focusedViewName == "boxscore" {
		focusedView.Title = spaced(tglTitle) + bracketed(bsTitle) + spaced(psTitle)
	}
	if focusedViewName == "playerstats" {
		focusedView.Title = spaced(tglTitle) + spaced(bsTitle) + bracketed(psTitle)
	}
}
