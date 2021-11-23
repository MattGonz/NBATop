package nbatop

import (
	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/utils"
)

func (nt *NBATop) layout(g *gocui.Gui) error {
	nt.State.MaxX, nt.State.MaxY = nt.G.Size()
	nt.State.GamesTodayLength = ((utils.Min(nt.State.MaxY/2, nt.Views.Today.NumLines) + 2) / 3 * 3) - 2

	easternConference := nt.Views.Standings.EasternConference
	westernConference := nt.Views.Standings.WesternConference
	longestEast := utils.Longest(easternConference)
	longestWest := utils.Longest(westernConference)

	nt.State.SidebarLength = utils.Min(len(westernConference)+len(easternConference)+nt.State.GamesTodayLength, nt.State.MaxY-1)
	nt.State.SidebarWidth = utils.Min(utils.Max(longestWest, longestEast)+4, nt.State.MaxX-1)
	nt.State.StandingsSpaces = utils.Max(longestWest, longestEast) - 8

	nt.DrawToday()
	nt.DrawStandings()

	// Placeholder for the main table view
	if v, err := g.SetView("main", nt.State.SidebarWidth, 0, nt.State.MaxX-1, nt.State.MaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "table goes here"
	}

	return nil
}
