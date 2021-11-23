package nbatop

import "github.com/jroimartin/gocui"

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
