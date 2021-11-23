package ui

import (
	"github.com/jroimartin/gocui"
)

type TodayView struct {
	v        *gocui.View
	name     string
	Games    [][]string
	NumLines int
}

// NewTodayView creates a new TodayView
func NewTodayView(g *gocui.Gui) *TodayView {
	return &TodayView{
		v:     &gocui.View{},
		name:  "today",
		Games: [][]string{},
	}
}
