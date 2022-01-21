package nbatop

import (
	"log"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
)

type Views struct {
	StandingsView   *StandingsView
	TodayView       *TodayView
	TableView       *EmptyTableView
	TeamGameLogView *TeamGameLogView
	BoxScoreView    *BoxScoreView
	PlayerStatsView *PlayerStatsView
}

type State struct {
	CurrentSeason        *api.CurrentNBASeason
	Standings            *api.NBAStandings
	StandingsSpaces      int
	GamesToday           *api.NBAToday
	GamesTodayLength     int
	ActivePlayers        *api.Players
	PersonIDToPlayerName map[string]string
	SidebarWidth         int
	SidebarLength        int
	Today                string
	MaxX, MaxY           int
	GameLogIdxToTeamID   map[int]string
	GameLogIdxToTeamName map[int]string
	TeamIDToTeamName     map[string]string
	FocusedTableView     string
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
			StandingsView:   NewStandingsView(),
			TodayView:       NewTodayView(),
			TableView:       NewTableView(),
			TeamGameLogView: NewTeamGameLogView(),
			BoxScoreView:    NewBoxScoreView(),
			PlayerStatsView: NewPlayerStatsView(),
		},
		State: &State{
			Standings:            &api.NBAStandings{},
			GamesToday:           &api.NBAToday{},
			Today:                time.Now().Format("01-02-2006"),
			PersonIDToPlayerName: make(map[string]string),
			GameLogIdxToTeamID:   make(map[int]string),
			GameLogIdxToTeamName: make(map[int]string),
			TeamIDToTeamName:     make(map[string]string),
			FocusedTableView:     "table",
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

// FormatConferenceStandings returns the formatted standings for the eastern and western conferences
func (nt *NBATop) FormatConferenceStandings() error {
	var westernConference []string
	var easternConference []string
	idx := 0

	for i, team := range nt.State.Standings.League.Standard.Conference.East {
		idx = i
		nt.State.GameLogIdxToTeamID[idx] = team.TeamID
		nt.State.GameLogIdxToTeamName[idx] = team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		nt.State.TeamIDToTeamName[team.TeamID] = team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
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
		idx += 1
		nt.State.GameLogIdxToTeamID[idx] = team.TeamID
		nt.State.GameLogIdxToTeamName[idx] = team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
		nt.State.TeamIDToTeamName[team.TeamID] = team.TeamSitesOnly.TeamName + " " + team.TeamSitesOnly.TeamNickname
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

	nt.Views.StandingsView.WesternConference = westernConference
	nt.Views.StandingsView.EasternConference = easternConference

	return nil
}

// FormatGamesToday returns the formatted game strings for today
func (nt *NBATop) FormatGamesToday() {
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
	nt.Views.TodayView.Games = games
	nt.Views.TodayView.NumLines = len(games) * 3
}

// MapPlayerIDs maps player IDs to player names
func (nt *NBATop) MapPlayerIDs() {
	for _, player := range nt.State.ActivePlayers.League.Players {
		fullName := player.FirstName + " " + player.LastName
		playerID := player.PersonID
		nt.State.PersonIDToPlayerName[playerID] = fullName
	}
}
