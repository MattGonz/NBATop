package ui

import (
	"strconv"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type StandingsView struct {
	v                 *gocui.View
	name              string
	WesternConference []string
	EasternConference []string
}

// NewStandingsView creates a new standings view
func NewStandingsView(g *gocui.Gui) *StandingsView {
	return &StandingsView{
		v:                 &gocui.View{},
		name:              "standings",
		WesternConference: []string{},
		EasternConference: []string{},
	}
}

// conferenceStandings returns the formatted standings for the eastern and western conferences
func (s *StandingsView) GetConferenceStandings(standings *api.NBAStandings) error {
	var westernConference []string
	var easternConference []string

	for _, team := range standings.League.Standard.Conference.East {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname

		wins, err := strconv.Atoi(team.Win)
		if err != nil {
			return err
		}

		var record string
		if wins < 10 {
			record = "    " + team.Win + "-" + team.Loss
		} else {
			record = "   " + team.Win + "-" + team.Loss
		}

		easternConference = append(easternConference, "\t"+name+"\t"+record)
	}

	for _, team := range standings.League.Standard.Conference.West {
		name := team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		wins, err := strconv.Atoi(team.Win)
		if err != nil {
			return err
		}

		var record string
		if wins < 10 {
			record = " " + team.Win + "-" + team.Loss
		} else {
			record = team.Win + "-" + team.Loss
		}

		westernConference = append(westernConference, "\t"+name+"\t"+record)
	}
	s.WesternConference = westernConference
	s.EasternConference = easternConference
	return nil
}
