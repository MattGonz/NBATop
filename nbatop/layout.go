package nbatop

import (
	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/utils"
)

// layout calculates the general layout and draws the initial views
func (nt *NBATop) layout(g *gocui.Gui) error {
	nt.State.MaxX, nt.State.MaxY = nt.G.Size()
	nt.State.GamesTodayLength = (utils.Min(nt.State.MaxY/2, nt.Views.TodayView.NumLines+4) / 3 * 3) - 2

	easternConference := nt.Views.StandingsView.EasternConference
	westernConference := nt.Views.StandingsView.WesternConference
	longestEast := utils.Longest(easternConference)
	longestWest := utils.Longest(westernConference)

	nt.State.SidebarLength = utils.Min(len(westernConference)+len(easternConference)+nt.State.GamesTodayLength, nt.State.MaxY-1)
	nt.State.SidebarWidth = utils.Min(utils.Max(longestWest, longestEast)+4, nt.State.MaxX-1)
	nt.State.StandingsSpaces = utils.Max(longestWest, longestEast) - 9

	nt.DrawToday()
	nt.DrawStandings()
	nt.DrawInitialTable()

	return nil
}
