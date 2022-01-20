package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// BoxScore represents a game's box score from the NBA Data API
type BoxScore struct {
	Internal struct {
		PubDateTime string `json:"pubDateTime"`
		Xslt        string `json:"xslt"`
		EventName   string `json:"eventName"`
	} `json:"_internal"`
	BasicGameData struct {
		SeasonStageID int    `json:"seasonStageId"`
		SeasonYear    string `json:"seasonYear"`
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
		EndTimeUTC            time.Time `json:"endTimeUTC"`
		StartDateEastern      string    `json:"startDateEastern"`
		Clock                 string    `json:"clock"`
		IsBuzzerBeater        bool      `json:"isBuzzerBeater"`
		IsPreviewArticleAvail bool      `json:"isPreviewArticleAvail"`
		IsRecapArticleAvail   bool      `json:"isRecapArticleAvail"`
		Tickets               struct {
			MobileApp  string `json:"mobileApp"`
			DesktopWeb string `json:"desktopWeb"`
			MobileWeb  string `json:"mobileWeb"`
		} `json:"tickets"`
		HasGameBookPdf bool `json:"hasGameBookPdf"`
		IsStartTimeTBD bool `json:"isStartTimeTBD"`
		Nugget         struct {
			Text string `json:"text"`
		} `json:"nugget"`
		Attendance   string `json:"attendance"`
		GameDuration struct {
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
			TeamID     string `json:"teamId"`
			TriCode    string `json:"triCode"`
			Win        string `json:"win"`
			Loss       string `json:"loss"`
			SeriesWin  string `json:"seriesWin"`
			SeriesLoss string `json:"seriesLoss"`
			Score      string `json:"score"`
			Linescore  []struct {
				Score string `json:"score"`
			} `json:"linescore"`
		} `json:"vTeam"`
		HTeam struct {
			TeamID     string `json:"teamId"`
			TriCode    string `json:"triCode"`
			Win        string `json:"win"`
			Loss       string `json:"loss"`
			SeriesWin  string `json:"seriesWin"`
			SeriesLoss string `json:"seriesLoss"`
			Score      string `json:"score"`
			Linescore  []struct {
				Score string `json:"score"`
			} `json:"linescore"`
		} `json:"hTeam"`
		Watch struct {
			Broadcast struct {
				Broadcasters struct {
					National []interface{} `json:"national"`
					Canadian []interface{} `json:"canadian"`
					VTeam    []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"vTeam"`
					HTeam []struct {
						ShortName string `json:"shortName"`
						LongName  string `json:"longName"`
					} `json:"hTeam"`
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
		Officials struct {
			Formatted []struct {
				FirstNameLastName string `json:"firstNameLastName"`
			} `json:"formatted"`
		} `json:"officials"`
	} `json:"basicGameData"`
	PreviousMatchup struct {
		GameID   string `json:"gameId"`
		GameDate string `json:"gameDate"`
	} `json:"previousMatchup"`
	Stats struct {
		TimesTied   string `json:"timesTied"`
		LeadChanges string `json:"leadChanges"`
		VTeam       struct {
			FastBreakPoints    string `json:"fastBreakPoints"`
			PointsInPaint      string `json:"pointsInPaint"`
			BiggestLead        string `json:"biggestLead"`
			SecondChancePoints string `json:"secondChancePoints"`
			PointsOffTurnovers string `json:"pointsOffTurnovers"`
			LongestRun         string `json:"longestRun"`
			Totals             struct {
				Points    string `json:"points"`
				Fgm       string `json:"fgm"`
				Fga       string `json:"fga"`
				Fgp       string `json:"fgp"`
				Ftm       string `json:"ftm"`
				Fta       string `json:"fta"`
				Ftp       string `json:"ftp"`
				Tpm       string `json:"tpm"`
				Tpa       string `json:"tpa"`
				Tpp       string `json:"tpp"`
				OffReb    string `json:"offReb"`
				DefReb    string `json:"defReb"`
				TotReb    string `json:"totReb"`
				Assists   string `json:"assists"`
				PFouls    string `json:"pFouls"`
				Steals    string `json:"steals"`
				Turnovers string `json:"turnovers"`
				Blocks    string `json:"blocks"`
				PlusMinus string `json:"plusMinus"`
				Min       string `json:"min"`
			} `json:"totals"`
			Leaders struct {
				Points struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"points"`
				Rebounds struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"rebounds"`
				Assists struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"assists"`
			} `json:"leaders"`
		} `json:"vTeam"`
		HTeam struct {
			FastBreakPoints    string `json:"fastBreakPoints"`
			PointsInPaint      string `json:"pointsInPaint"`
			BiggestLead        string `json:"biggestLead"`
			SecondChancePoints string `json:"secondChancePoints"`
			PointsOffTurnovers string `json:"pointsOffTurnovers"`
			LongestRun         string `json:"longestRun"`
			Totals             struct {
				Points    string `json:"points"`
				Fgm       string `json:"fgm"`
				Fga       string `json:"fga"`
				Fgp       string `json:"fgp"`
				Ftm       string `json:"ftm"`
				Fta       string `json:"fta"`
				Ftp       string `json:"ftp"`
				Tpm       string `json:"tpm"`
				Tpa       string `json:"tpa"`
				Tpp       string `json:"tpp"`
				OffReb    string `json:"offReb"`
				DefReb    string `json:"defReb"`
				TotReb    string `json:"totReb"`
				Assists   string `json:"assists"`
				PFouls    string `json:"pFouls"`
				Steals    string `json:"steals"`
				Turnovers string `json:"turnovers"`
				Blocks    string `json:"blocks"`
				PlusMinus string `json:"plusMinus"`
				Min       string `json:"min"`
			} `json:"totals"`
			Leaders struct {
				Points struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"points"`
				Rebounds struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"rebounds"`
				Assists struct {
					Value   string `json:"value"`
					Players []struct {
						PersonID string `json:"personId"`
					} `json:"players"`
				} `json:"assists"`
			} `json:"leaders"`
		} `json:"hTeam"`
		ActivePlayers []struct {
			PersonID  string `json:"personId"`
			TeamID    string `json:"teamId"`
			IsOnCourt bool   `json:"isOnCourt"`
			Points    string `json:"points"`
			Pos       string `json:"pos"`
			Min       string `json:"min"`
			Fgm       string `json:"fgm"`
			Fga       string `json:"fga"`
			Fgp       string `json:"fgp"`
			Ftm       string `json:"ftm"`
			Fta       string `json:"fta"`
			Ftp       string `json:"ftp"`
			Tpm       string `json:"tpm"`
			Tpa       string `json:"tpa"`
			Tpp       string `json:"tpp"`
			OffReb    string `json:"offReb"`
			DefReb    string `json:"defReb"`
			TotReb    string `json:"totReb"`
			Assists   string `json:"assists"`
			PFouls    string `json:"pFouls"`
			Steals    string `json:"steals"`
			Turnovers string `json:"turnovers"`
			Blocks    string `json:"blocks"`
			PlusMinus string `json:"plusMinus"`
			Dnp       string `json:"dnp"`
			SortKey   struct {
				Name      int `json:"name"`
				Pos       int `json:"pos"`
				Points    int `json:"points"`
				Min       int `json:"min"`
				Fgm       int `json:"fgm"`
				Fga       int `json:"fga"`
				Fgp       int `json:"fgp"`
				Ftm       int `json:"ftm"`
				Fta       int `json:"fta"`
				Ftp       int `json:"ftp"`
				Tpm       int `json:"tpm"`
				Tpa       int `json:"tpa"`
				Tpp       int `json:"tpp"`
				OffReb    int `json:"offReb"`
				DefReb    int `json:"defReb"`
				TotReb    int `json:"totReb"`
				Assists   int `json:"assists"`
				PFouls    int `json:"pFouls"`
				Steals    int `json:"steals"`
				Turnovers int `json:"turnovers"`
				Blocks    int `json:"blocks"`
				PlusMinus int `json:"plusMinus"`
			} `json:"sortKey"`
		} `json:"activePlayers"`
	} `json:"stats"`
}

// GetBoxScore fetches and structures the box score of a given game date and ID,
// from the NBA Data API
func GetBoxScore(gameDate, gameID string) *BoxScore {
	time, err := time.Parse("JAN 2, 2006", gameDate)
	if err != nil {
		log.Panicln(err)
	}
	gameDateFormatted := time.Format("20060102")

	url := fmt.Sprintf("https://data.nba.net/prod/v1/%s/%s_boxscore.json", gameDateFormatted, gameID)

	req, err := http.NewRequest("GET", url, nil)
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

	var result *BoxScore
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panicln(err)
	}
	return result
}
