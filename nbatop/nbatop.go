package nbatop

import (
	"log"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type Views struct {
	Standings *StandingsView
	Today     *TodayView
}

type State struct {
	Standings        *api.NBAStandings
	StandingsSpaces  int
	GamesToday       *api.NBAToday
	GamesTodayLength int
	SidebarWidth     int
	SidebarLength    int
	Today            string
	MaxX, MaxY       int
}

type NBATop struct {
	G     *gocui.Gui
	Views *Views
	State *State
}

// NewNBATop creates a new instance of NBATop
func NewNBATop() *NBATop {
	nbatop := NBATop{
		Views: &Views{
			Standings: NewStandingsView(),
			Today:     NewTodayView(),
		},
		State: &State{
			Standings:  &api.NBAStandings{},
			GamesToday: &api.NBAToday{},
			Today:      time.Now().Format("01-02-2006"),
		},
	}

	return &nbatop
}

// quit exits the GUI
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Run starts the GUI
func (nt *NBATop) Run() {

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	nt.G = g
	g.Cursor = true
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(nt.layout)

	if err := nt.SetKeybindings(); err != nil {
		log.Panicln(err)
	}

	if err := nt.G.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// GetConferenceStandings returns the formatted standings for the eastern and western conferences
func (nt *NBATop) GetConferenceStandings() error {
	var westernConference []string
	var easternConference []string

	for _, team := range nt.State.Standings.League.Standard.Conference.East {
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

	for _, team := range nt.State.Standings.League.Standard.Conference.West {
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

	nt.Views.Standings.WesternConference = westernConference
	nt.Views.Standings.EasternConference = easternConference

	return nil
}

// GetGames gets the formatted game strings for today
func (nt *NBATop) GetGames() {
	var games [][]string

	for _, game := range nt.State.GamesToday.Games {
		homeTeam := game.HTeam
		homeTeamRecord := " (" + homeTeam.Win + "-" + homeTeam.Loss + ")"
		awayTeam := game.VTeam
		awayTeamRecord := "(" + awayTeam.Win + "-" + awayTeam.Loss + ") "

		gameInfo := awayTeamRecord + awayTeam.TriCode + " at " + homeTeam.TriCode + homeTeamRecord
		// spaces := strings.Repeat(" ", len(gameInfo)/3)
		startTime := game.StartTimeEastern

		games = append(games, []string{startTime, gameInfo})
	}
	nt.Views.Today.Games = games
	nt.Views.Today.NumLines = len(games) * 3
}
