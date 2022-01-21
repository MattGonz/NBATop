package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// NBAToday represents today's game schedule from the NBA Data API
type NBAToday struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		RouteName               string `json:"routeName"`
		RouteValue              string `json:"routeValue"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	NumGames int `json:"numGames"`
	Games    []struct {
		SeasonStageID int    `json:"seasonStageId"`
		SeasonYear    string `json:"seasonYear"`
		LeagueName    string `json:"leagueName"`
		GameID        string `json:"gameId"`
		Arena         struct {
			Name       string `json:"name"`
			IsDomestic bool   `json:"isDomestic"`
			City       string `json:"city"`
			StateAbbr  string `json:"stateAbbr"`
			Country    string `json:"country"`
		} `json:"arena"`
		IsGameActivated       bool      `json:"isGameActivated"`
		StatusNum             int       `json:"statusNum"`
		ExtendedStatusNum     int       `json:"extendedStatusNum"`
		StartTimeEastern      string    `json:"startTimeEastern"`
		StartTimeUTC          time.Time `json:"startTimeUTC"`
		StartDateEastern      string    `json:"startDateEastern"`
		HomeStartDate         string    `json:"homeStartDate"`
		HomeStartTime         string    `json:"homeStartTime"`
		VisitorStartDate      string    `json:"visitorStartDate"`
		VisitorStartTime      string    `json:"visitorStartTime"`
		GameURLCode           string    `json:"gameUrlCode"`
		Clock                 string    `json:"clock"`
		IsBuzzerBeater        bool      `json:"isBuzzerBeater"`
		IsPreviewArticleAvail bool      `json:"isPreviewArticleAvail"`
		IsRecapArticleAvail   bool      `json:"isRecapArticleAvail"`
		Nugget                struct {
			Text string `json:"text"`
		} `json:"nugget"`
		Attendance string `json:"attendance"`
		Tickets    struct {
			MobileApp    string `json:"mobileApp"`
			DesktopWeb   string `json:"desktopWeb"`
			MobileWeb    string `json:"mobileWeb"`
			LeagGameInfo string `json:"leagGameInfo"`
			LeagTix      string `json:"leagTix"`
		} `json:"tickets"`
		HasGameBookPdf bool `json:"hasGameBookPdf"`
		IsStartTimeTBD bool `json:"isStartTimeTBD"`
		IsNeutralVenue bool `json:"isNeutralVenue"`
		GameDuration   struct {
			Hours   string `json:"hours"`
			Minutes string `json:"minutes"`
		} `json:"gameDuration"`
		Period struct {
			Current       int  `json:"current"`
			Type          int  `json:"type"`
			MaxRegular    int  `json:"maxRegular"`
			IsHalftime    bool `json:"isHalftime"`
			IsEndOfPeriod bool `json:"isEndOfPeriod"`
		} `json:"period"`
		VTeam struct {
			TeamID     string        `json:"teamId"`
			TriCode    string        `json:"triCode"`
			Win        string        `json:"win"`
			Loss       string        `json:"loss"`
			SeriesWin  string        `json:"seriesWin"`
			SeriesLoss string        `json:"seriesLoss"`
			Score      string        `json:"score"`
			Linescore  []interface{} `json:"linescore"`
		} `json:"vTeam"`
		HTeam struct {
			TeamID     string        `json:"teamId"`
			TriCode    string        `json:"triCode"`
			Win        string        `json:"win"`
			Loss       string        `json:"loss"`
			SeriesWin  string        `json:"seriesWin"`
			SeriesLoss string        `json:"seriesLoss"`
			Score      string        `json:"score"`
			Linescore  []interface{} `json:"linescore"`
		} `json:"hTeam"`
		Watch struct {
			Broadcast struct {
				Broadcasters struct {
					National []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"national"`
					Canadian []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"canadian"`
					VTeam           []interface{} `json:"vTeam"`
					HTeam           []interface{} `json:"hTeam"`
					SpanishHTeam    []interface{} `json:"spanish_hTeam"`
					SpanishVTeam    []interface{} `json:"spanish_vTeam"`
					SpanishNational []interface{} `json:"spanish_national"`
				} `json:"broadcasters"`
				Video struct {
					RegionalBlackoutCodes string `json:"regionalBlackoutCodes"`
					CanPurchase           bool   `json:"canPurchase"`
					IsLeaguePass          bool   `json:"isLeaguePass"`
					IsNationalBlackout    bool   `json:"isNationalBlackout"`
					IsTNTOT               bool   `json:"isTNTOT"`
					IsVR                  bool   `json:"isVR"`
					TntotIsOnAir          bool   `json:"tntotIsOnAir"`
					IsNextVR              bool   `json:"isNextVR"`
					IsNBAOnTNTVR          bool   `json:"isNBAOnTNTVR"`
					IsMagicLeap           bool   `json:"isMagicLeap"`
					IsOculusVenues        bool   `json:"isOculusVenues"`
					Streams               []struct {
						StreamType            string `json:"streamType"`
						IsOnAir               bool   `json:"isOnAir"`
						DoesArchiveExist      bool   `json:"doesArchiveExist"`
						IsArchiveAvailToWatch bool   `json:"isArchiveAvailToWatch"`
						StreamID              string `json:"streamId"`
						Duration              int    `json:"duration"`
					} `json:"streams"`
					DeepLink []struct {
						Broadcaster         string `json:"broadcaster"`
						RegionalMarketCodes string `json:"regionalMarketCodes"`
						IosApp              string `json:"iosApp"`
						AndroidApp          string `json:"androidApp"`
						DesktopWeb          string `json:"desktopWeb"`
						MobileWeb           string `json:"mobileWeb"`
					} `json:"deepLink"`
				} `json:"video"`
				Audio struct {
					National struct {
						Streams []struct {
							Language string `json:"language"`
							IsOnAir  bool   `json:"isOnAir"`
							StreamID string `json:"streamId"`
						} `json:"streams"`
						Broadcasters []interface{} `json:"broadcasters"`
					} `json:"national"`
					VTeam struct {
						Streams []struct {
							Language string `json:"language"`
							IsOnAir  bool   `json:"isOnAir"`
							StreamID string `json:"streamId"`
						} `json:"streams"`
						Broadcasters []struct {
							ShortName string `json:"shortName"`
							LongName  string `json:"longName"`
						} `json:"broadcasters"`
					} `json:"vTeam"`
					HTeam struct {
						Streams []struct {
							Language string `json:"language"`
							IsOnAir  bool   `json:"isOnAir"`
							StreamID string `json:"streamId"`
						} `json:"streams"`
						Broadcasters []struct {
							ShortName string `json:"shortName"`
							LongName  string `json:"longName"`
						} `json:"broadcasters"`
					} `json:"hTeam"`
				} `json:"audio"`
			} `json:"broadcast"`
		} `json:"watch"`
	} `json:"games"`
}

// CurrentNBASeason respresents the current NBA season
type CurrentNBASeason struct {
	Internal struct {
		PubDateTime             string `json:"pubDateTime"`
		IgorPath                string `json:"igorPath"`
		Xslt                    string `json:"xslt"`
		XsltForceRecompile      string `json:"xsltForceRecompile"`
		XsltInCache             string `json:"xsltInCache"`
		XsltCompileTimeMillis   string `json:"xsltCompileTimeMillis"`
		XsltTransformTimeMillis string `json:"xsltTransformTimeMillis"`
		ConsolidatedDomKey      string `json:"consolidatedDomKey"`
		EndToEndTimeMillis      string `json:"endToEndTimeMillis"`
	} `json:"_internal"`
	TeamSitesOnly struct {
		SeasonStage    int    `json:"seasonStage"`
		SeasonYear     int    `json:"seasonYear"`
		RosterYear     int    `json:"rosterYear"`
		StatsStage     int    `json:"statsStage"`
		StatsYear      int    `json:"statsYear"`
		DisplayYear    string `json:"displayYear"`
		LastPlayByPlay string `json:"lastPlayByPlay"`
		AllPlayByPlay  string `json:"allPlayByPlay"`
		PlayerMatchup  string `json:"playerMatchup"`
		Series         string `json:"series"`
	} `json:"teamSitesOnly"`
	SeasonScheduleYear int  `json:"seasonScheduleYear"`
	ShowPlayoffsClinch bool `json:"showPlayoffsClinch"`
	Links              struct {
		AnchorDate                  string `json:"anchorDate"`
		CurrentDate                 string `json:"currentDate"`
		Calendar                    string `json:"calendar"`
		TodayScoreboard             string `json:"todayScoreboard"`
		CurrentScoreboard           string `json:"currentScoreboard"`
		Teams                       string `json:"teams"`
		Scoreboard                  string `json:"scoreboard"`
		LeagueRosterPlayers         string `json:"leagueRosterPlayers"`
		AllstarRoster               string `json:"allstarRoster"`
		LeagueRosterCoaches         string `json:"leagueRosterCoaches"`
		LeagueSchedule              string `json:"leagueSchedule"`
		LeagueConfStandings         string `json:"leagueConfStandings"`
		LeagueDivStandings          string `json:"leagueDivStandings"`
		LeagueUngroupedStandings    string `json:"leagueUngroupedStandings"`
		LeagueMiniStandings         string `json:"leagueMiniStandings"`
		LeagueTeamStatsLeaders      string `json:"leagueTeamStatsLeaders"`
		LeagueLastFiveGameTeamStats string `json:"leagueLastFiveGameTeamStats"`
		PreviewArticle              string `json:"previewArticle"`
		RecapArticle                string `json:"recapArticle"`
		GameBookPdf                 string `json:"gameBookPdf"`
		Boxscore                    string `json:"boxscore"`
		MiniBoxscore                string `json:"miniBoxscore"`
		Pbp                         string `json:"pbp"`
		LeadTracker                 string `json:"leadTracker"`
		PlayerGameLog               string `json:"playerGameLog"`
		PlayerProfile               string `json:"playerProfile"`
		PlayerUberStats             string `json:"playerUberStats"`
		TeamSchedule                string `json:"teamSchedule"`
		TeamsConfig                 string `json:"teamsConfig"`
		TeamRoster                  string `json:"teamRoster"`
		TeamsConfigYear             string `json:"teamsConfigYear"`
		TeamScheduleYear            string `json:"teamScheduleYear"`
		TeamLeaders                 string `json:"teamLeaders"`
		TeamScheduleYear2           string `json:"teamScheduleYear2"`
		TeamLeaders2                string `json:"teamLeaders2"`
		TeamICS                     string `json:"teamICS"`
		TeamICS2                    string `json:"teamICS2"`
		PlayoffsBracket             string `json:"playoffsBracket"`
		PlayoffSeriesLeaders        string `json:"playoffSeriesLeaders"`
		UniversalLinkMapping        string `json:"universalLinkMapping"`
		TicketLink                  string `json:"ticketLink"`
	} `json:"links"`
}

// GetGamesToday fetches and structures today's schedule, fetched from the NBA Data API
func GetGamesToday() *NBAToday {
	today := time.Now().Format("20060102")
	url := fmt.Sprintf("http://data.nba.net/prod/v2/%s/scoreboard.json", today)
	response, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln(err)
	}

	var result *NBAToday
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panicln(err)
	}
	return result
}

// GetCurrentSeason fetches and structures information about the current season,
// fetched from the NBA Data API
func GetCurrentSeason() *CurrentNBASeason {
	todayURL := "http://data.nba.net/prod/v3/today.json"

	req, err := http.NewRequest("GET", todayURL, nil)
	if err != nil {
		log.Panicln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	var today *CurrentNBASeason
	if err := json.Unmarshal(body, &today); err != nil {
		log.Panicln(err)
	}
	return today
}
