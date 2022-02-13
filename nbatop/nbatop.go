package nbatop

import (
	"log"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mattgonz/nbatop/api"
	"github.com/mattgonz/nbatop/utils"
)

type Views struct {
	StandingsView   *StandingsView
	TodayView       *TodayView
	TableView       *TableView
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
	GamesTodayIdxToGame  map[int][]string
	ActivePlayers        *api.Players
	PersonIDToPlayerName map[string]string
	GameLogIdxToTeamID   map[int]string
	GameLogIdxToTeamName map[int]string
	SidebarWidth         int
	SidebarLength        int
	Today                string
	MaxX, MaxY           int
	TeamIDToTeamName     map[string]string
	FocusedTableView     string
	LastSidebarView      string
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
			StandingsView: NewStandingsView(),
			TodayView:     NewTodayView(),
			// TableView:     NewTableView(),
		},
		State: &State{
			Standings:            &api.NBAStandings{},
			GamesToday:           &api.NBAToday{},
			Today:                time.Now().Format("01-02-2006"),
			GamesTodayIdxToGame:  make(map[int][]string),
			TeamIDToTeamName:     make(map[string]string),
			PersonIDToPlayerName: make(map[string]string),
			GameLogIdxToTeamID:   make(map[int]string),
			GameLogIdxToTeamName: make(map[int]string),
			FocusedTableView:     "table",
			LastSidebarView:      "standings",
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

	nt.State.MaxX, nt.State.MaxY = nt.G.Size()
	nt.State.GamesTodayLength = (utils.Min(nt.State.MaxY/2, nt.Views.TodayView.NumLines+4) / 3 * 3) - 2

	easternConference := nt.Views.StandingsView.EasternConference
	westernConference := nt.Views.StandingsView.WesternConference
	longestEast := utils.Longest(easternConference)
	longestWest := utils.Longest(westernConference)

	nt.State.SidebarLength = utils.Min(len(westernConference)+len(easternConference)+nt.State.GamesTodayLength+3, nt.State.MaxY-1)
	nt.State.SidebarWidth = utils.Min(utils.Max(longestWest, longestEast)+4, nt.State.MaxX-1)
	nt.State.StandingsSpaces = utils.Max(longestWest, longestEast) - 9

	nt.Views.TeamGameLogView = nt.NewTeamGameLogView()
	nt.Views.BoxScoreView = nt.NewBoxScoreView()
	nt.Views.PlayerStatsView = nt.NewPlayerStatsView()

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

	for idx, game := range nt.State.GamesToday.Games {
		var gameInfo string
		var gameTime string

		homeTeam := game.HTeam
		awayTeam := game.VTeam

		homeTeamPoints := game.HTeam.Score
		awayTeamPoints := game.VTeam.Score

		if game.IsGameActivated {
			gameInfo = "(" + awayTeamPoints + ") " + awayTeam.TriCode + " at " + homeTeam.TriCode + " (" + homeTeamPoints + ")"
			quarter := strconv.Itoa(game.Period.Current)
			gameTime = game.Clock
			if gameTime == "" {
				if game.Period.IsHalftime {
					gameTime = "Halftime"
				} else {
					gameTime = "End Q" + quarter
				}
			} else {
				gameTime = "(Q" + quarter + " - " + gameTime + ")"
			}
		} else {
			homeTeamRecord := " (" + homeTeam.Win + "-" + homeTeam.Loss + ")"
			awayTeamRecord := "(" + awayTeam.Win + "-" + awayTeam.Loss + ") "

			gameInfo = awayTeamRecord + awayTeam.TriCode + " at " + homeTeam.TriCode + homeTeamRecord

			gameTimeUTC := game.StartTimeUTC

			// TODO timezones
			loc, err := time.LoadLocation("America/New_York")
			if err != nil {
				log.Panicln(err)
			}

			if gameTimeUTC.Before(time.Now().In(loc)) {
				homeTeamPointsVal, _ := strconv.Atoi(homeTeamPoints)
				awayTeamPointsVal, _ := strconv.Atoi(awayTeamPoints)

				gameTime = "Final"
				if awayTeamPointsVal > homeTeamPointsVal {
					gameInfo = "(" + awayTeamPoints + ") " + awayTeam.TriCode + " def " + homeTeam.TriCode + " (" + homeTeamPoints + ")"
				} else {
					gameInfo = "(" + homeTeamPoints + ") " + homeTeam.TriCode + " def " + awayTeam.TriCode + " (" + awayTeamPoints + ")"
				}
			} else {
				gameTime = gameTimeUTC.In(loc).Format("3:04PM")
			}

		}
		games = append(games, []string{gameTime, gameInfo})

		gameID := game.GameID
		matchup := game.VTeam.TriCode + " @ " + game.HTeam.TriCode

		gameDate, err := time.Parse("20060102", game.StartDateEastern)
		if err != nil {
			log.Panicln(err)
		}
		gameDateStr := gameDate.Format("Jan 02, 2006")

		nt.State.GamesTodayIdxToGame[idx*3] = []string{gameID, gameDateStr, matchup}
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
