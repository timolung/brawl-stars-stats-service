package services

import (
	"encoding/json"
	"net/url"

	"github.com/timolung/brawl-stars-stats-service/internal/config"
	"github.com/timolung/brawl-stars-stats-service/internal/constant"
	"github.com/timolung/brawl-stars-stats-service/internal/models"
	"github.com/timolung/brawl-stars-stats-service/internal/utils"
)

type Stat struct {
	Description string      `json:"description"`
	Value       interface{} `json:"value"`
}

type PlayerService struct {
	PlayerTag          string
	PlayerTagUnencoded string
}

func NewPlayerService(playerTag string) *PlayerService {
	playerTag = "#" + playerTag
	playerTagEncoded := url.PathEscape(playerTag)
	return &PlayerService{
		PlayerTag:          playerTagEncoded,
		PlayerTagUnencoded: playerTag,
	}
}

func (bs *PlayerService) GetBattleLog() (models.BattleLogResponse, error) {
	var apiResponse models.BattleLogResponse
	data, err := utils.MakeAPIRequest(config.Cfg.BattleLogEndpoint, "{playerTag}", bs.PlayerTag)
	if err != nil {
		return apiResponse, err
	}

	// parse API Response
	jsonData, _ := json.Marshal(data)
	err = json.Unmarshal(jsonData, &apiResponse)
	if err != nil {
		return apiResponse, err
	}

	return apiResponse, nil
}

func (bs *PlayerService) CalculatePlayerStats(apiResponse models.BattleLogResponse) ([]Stat, error) {
	if len(apiResponse.Items) == 0 {
		return []Stat{
			{Description: "No Game History", Value: "No Data"},
		}, nil
	}

	// Initialize counters
	starPlayerCount := 0
	starPlayerVictoryCount := 0
	starPlayerDefeatCount := 0
	victoryCount := 0
	defeatCount := 0
	tiedCount := 0

	// Iterate over the items
	for _, item := range apiResponse.Items {
		battle := item.Battle

		// Check if the player was the star player

		// Check if the player's team won or lost the battle
		if battle.Result == "victory" {
			if battle.StarPlayer != nil && battle.StarPlayer.Tag == bs.PlayerTagUnencoded {
				starPlayerCount++
				starPlayerVictoryCount++
			}
			victoryCount++
		} else if battle.Result == "defeat" {
			if battle.StarPlayer != nil && battle.StarPlayer.Tag == bs.PlayerTagUnencoded {
				starPlayerCount++
				starPlayerDefeatCount++
			}
			defeatCount++
		} else {
			tiedCount++
		}
	}

	// Calculate percentages
	totalGames := len(apiResponse.Items)
	lastPlayed := utils.CalculateDuration(apiResponse.Items[0].BattleTime)
	earliestPlayed := utils.CalculateDuration(apiResponse.Items[totalGames-1].BattleTime)

	starPlayerPercent := utils.RoundToNearestTwoDecimals(float64(starPlayerCount) / float64(totalGames) * 100)
	if totalGames == 0 {
		starPlayerPercent = 0
	}

	starPlayerVictoryPercent := utils.RoundToNearestTwoDecimals(float64(starPlayerVictoryCount) / float64(victoryCount) * 100)
	if victoryCount == 0 {
		starPlayerVictoryPercent = 0
	}

	starPlayerDefeatPercent := utils.RoundToNearestTwoDecimals(float64(starPlayerDefeatCount) / float64(defeatCount) * 100)
	if defeatCount == 0 {
		starPlayerDefeatPercent = 0
	}

	victoryStarPlayerPercent := utils.RoundToNearestTwoDecimals(float64(starPlayerVictoryCount) / float64(starPlayerCount) * 100)
	if victoryCount == 0 {
		victoryStarPlayerPercent = 0
	}

	defeatStarPlayerPercent := utils.RoundToNearestTwoDecimals(float64(starPlayerDefeatCount) / float64(starPlayerCount) * 100)
	if defeatCount == 0 {
		defeatStarPlayerPercent = 0
	}

	stats := []Stat{
		{Description: constant.BattleLogTotalGamesDescription, Value: totalGames},
		{Description: constant.BattleLogLastPlayedDescription, Value: lastPlayed},
		{Description: constant.BattleLogTotalTimeDescription, Value: earliestPlayed},
		{Description: constant.BattleLogStarPlayerDescription, Value: starPlayerPercent},
		{Description: constant.BattleLogTotalVictoriesDescription, Value: victoryCount},
		{Description: constant.BattleLogTotalDefeatsDescription, Value: defeatCount},
		{Description: constant.BattleLogTotalTiesDescription, Value: tiedCount},
		{Description: constant.BattleLogStarPlayerVictoryDescription, Value: starPlayerVictoryPercent},
		{Description: constant.BattleLogStarPlayerDefeatDescription, Value: starPlayerDefeatPercent},
		{Description: constant.BattleLogVictoryStarPlayerDescription, Value: victoryStarPlayerPercent},
		{Description: constant.BattleLogDefeatStarPlayerDescription, Value: defeatStarPlayerPercent},
	}

	return stats, nil
}
