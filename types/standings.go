package types

type NBAStandings struct {
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
	League struct {
		Standard struct {
			SeasonYear    int `json:"seasonYear"`
			SeasonStageID int `json:"seasonStageId"`
			Teams         []struct {
				TeamID                 string `json:"teamId"`
				Win                    string `json:"win"`
				Loss                   string `json:"loss"`
				WinPct                 string `json:"winPct"`
				WinPctV2               string `json:"winPctV2"`
				LossPct                string `json:"lossPct"`
				LossPctV2              string `json:"lossPctV2"`
				GamesBehind            string `json:"gamesBehind"`
				DivGamesBehind         string `json:"divGamesBehind"`
				ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
				ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
				ConfRank               string `json:"confRank"`
				ConfWin                string `json:"confWin"`
				ConfLoss               string `json:"confLoss"`
				DivWin                 string `json:"divWin"`
				DivLoss                string `json:"divLoss"`
				HomeWin                string `json:"homeWin"`
				HomeLoss               string `json:"homeLoss"`
				AwayWin                string `json:"awayWin"`
				AwayLoss               string `json:"awayLoss"`
				LastTenWin             string `json:"lastTenWin"`
				LastTenLoss            string `json:"lastTenLoss"`
				Streak                 string `json:"streak"`
				DivRank                string `json:"divRank"`
				IsWinStreak            bool   `json:"isWinStreak"`
				TeamSitesOnly          struct {
					TeamKey            string `json:"teamKey"`
					TeamName           string `json:"teamName"`
					TeamCode           string `json:"teamCode"`
					TeamNickname       string `json:"teamNickname"`
					TeamTricode        string `json:"teamTricode"`
					ClinchedConference string `json:"clinchedConference"`
					ClinchedDivision   string `json:"clinchedDivision"`
					ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
					StreakText         string `json:"streakText"`
				} `json:"teamSitesOnly"`
				TieBreakerPts string `json:"tieBreakerPts"`
			} `json:"teams"`
		} `json:"standard"`
		Africa struct {
			SeasonYear    int           `json:"seasonYear"`
			SeasonStageID int           `json:"seasonStageId"`
			Teams         []interface{} `json:"teams"`
		} `json:"africa"`
		Sacramento struct {
			SeasonYear    int `json:"seasonYear"`
			SeasonStageID int `json:"seasonStageId"`
			Teams         []struct {
				TeamID                 string `json:"teamId"`
				Win                    string `json:"win"`
				Loss                   string `json:"loss"`
				WinPct                 string `json:"winPct"`
				WinPctV2               string `json:"winPctV2"`
				LossPct                string `json:"lossPct"`
				LossPctV2              string `json:"lossPctV2"`
				GamesBehind            string `json:"gamesBehind"`
				DivGamesBehind         string `json:"divGamesBehind"`
				ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
				ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
				ConfRank               string `json:"confRank"`
				ConfWin                string `json:"confWin"`
				ConfLoss               string `json:"confLoss"`
				DivWin                 string `json:"divWin"`
				DivLoss                string `json:"divLoss"`
				HomeWin                string `json:"homeWin"`
				HomeLoss               string `json:"homeLoss"`
				AwayWin                string `json:"awayWin"`
				AwayLoss               string `json:"awayLoss"`
				LastTenWin             string `json:"lastTenWin"`
				LastTenLoss            string `json:"lastTenLoss"`
				Streak                 string `json:"streak"`
				DivRank                string `json:"divRank"`
				IsWinStreak            bool   `json:"isWinStreak"`
				TeamSitesOnly          struct {
					TeamKey            string `json:"teamKey"`
					TeamName           string `json:"teamName"`
					TeamCode           string `json:"teamCode"`
					TeamNickname       string `json:"teamNickname"`
					TeamTricode        string `json:"teamTricode"`
					ClinchedConference string `json:"clinchedConference"`
					ClinchedDivision   string `json:"clinchedDivision"`
					ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
					StreakText         string `json:"streakText"`
				} `json:"teamSitesOnly"`
				TieBreakerPts string `json:"tieBreakerPts"`
			} `json:"teams"`
		} `json:"sacramento"`
		Vegas struct {
			SeasonYear    int `json:"seasonYear"`
			SeasonStageID int `json:"seasonStageId"`
			Teams         []struct {
				TeamID                 string `json:"teamId"`
				Win                    string `json:"win"`
				Loss                   string `json:"loss"`
				WinPct                 string `json:"winPct"`
				WinPctV2               string `json:"winPctV2"`
				LossPct                string `json:"lossPct"`
				LossPctV2              string `json:"lossPctV2"`
				GamesBehind            string `json:"gamesBehind"`
				DivGamesBehind         string `json:"divGamesBehind"`
				ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
				ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
				ConfRank               string `json:"confRank"`
				ConfWin                string `json:"confWin"`
				ConfLoss               string `json:"confLoss"`
				DivWin                 string `json:"divWin"`
				DivLoss                string `json:"divLoss"`
				HomeWin                string `json:"homeWin"`
				HomeLoss               string `json:"homeLoss"`
				AwayWin                string `json:"awayWin"`
				AwayLoss               string `json:"awayLoss"`
				LastTenWin             string `json:"lastTenWin"`
				LastTenLoss            string `json:"lastTenLoss"`
				Streak                 string `json:"streak"`
				DivRank                string `json:"divRank"`
				IsWinStreak            bool   `json:"isWinStreak"`
				TeamSitesOnly          struct {
					TeamKey            string `json:"teamKey"`
					TeamName           string `json:"teamName"`
					TeamCode           string `json:"teamCode"`
					TeamNickname       string `json:"teamNickname"`
					TeamTricode        string `json:"teamTricode"`
					ClinchedConference string `json:"clinchedConference"`
					ClinchedDivision   string `json:"clinchedDivision"`
					ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
					StreakText         string `json:"streakText"`
				} `json:"teamSitesOnly"`
				TieBreakerPts string `json:"tieBreakerPts"`
			} `json:"teams"`
		} `json:"vegas"`
		Utah struct {
			SeasonYear    int `json:"seasonYear"`
			SeasonStageID int `json:"seasonStageId"`
			Teams         []struct {
				TeamID                 string `json:"teamId"`
				Win                    string `json:"win"`
				Loss                   string `json:"loss"`
				WinPct                 string `json:"winPct"`
				WinPctV2               string `json:"winPctV2"`
				LossPct                string `json:"lossPct"`
				LossPctV2              string `json:"lossPctV2"`
				GamesBehind            string `json:"gamesBehind"`
				DivGamesBehind         string `json:"divGamesBehind"`
				ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
				ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
				ConfRank               string `json:"confRank"`
				ConfWin                string `json:"confWin"`
				ConfLoss               string `json:"confLoss"`
				DivWin                 string `json:"divWin"`
				DivLoss                string `json:"divLoss"`
				HomeWin                string `json:"homeWin"`
				HomeLoss               string `json:"homeLoss"`
				AwayWin                string `json:"awayWin"`
				AwayLoss               string `json:"awayLoss"`
				LastTenWin             string `json:"lastTenWin"`
				LastTenLoss            string `json:"lastTenLoss"`
				Streak                 string `json:"streak"`
				DivRank                string `json:"divRank"`
				IsWinStreak            bool   `json:"isWinStreak"`
				TeamSitesOnly          struct {
					TeamKey            string `json:"teamKey"`
					TeamName           string `json:"teamName"`
					TeamCode           string `json:"teamCode"`
					TeamNickname       string `json:"teamNickname"`
					TeamTricode        string `json:"teamTricode"`
					ClinchedConference string `json:"clinchedConference"`
					ClinchedDivision   string `json:"clinchedDivision"`
					ClinchedPlayoffs   string `json:"clinchedPlayoffs"`
					StreakText         string `json:"streakText"`
				} `json:"teamSitesOnly"`
				TieBreakerPts string `json:"tieBreakerPts"`
			} `json:"teams"`
		} `json:"utah"`
	} `json:"league"`
}
