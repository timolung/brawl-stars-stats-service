package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timolung/brawl-stars-stats-service/internal/config"
	"github.com/timolung/brawl-stars-stats-service/internal/constant"
	"github.com/timolung/brawl-stars-stats-service/internal/models"
)

var (
	sampleData = models.BattleLogResponse{
		Items: []models.BattleLogItem{
			{
				Battle: models.Battle{
					Duration: 180,
					Mode:     "Gem Grab",
					Result:   "victory",
					StarPlayer: &models.Player{
						Brawler: models.Brawler{
							ID:       1,
							Name:     "Shelly",
							Power:    9,
							Trophies: 500,
						},
						Name: "TestName",
						Tag:  "#PLAYER123",
					},
					Teams: []models.Team{
						{
							{
								Brawler: models.Brawler{
									ID:       2,
									Name:     "Colt",
									Power:    8,
									Trophies: 600,
								},
								Name: "Player1",
								Tag:  "#DEF456",
							},
						},
					},
					Type: "ranked",
				},
				BattleTime: "20240604T010049.000Z",
				Event: struct {
					ID   int    `json:"id"`
					Map  string `json:"map"`
					Mode string `json:"mode"`
				}{
					ID:   123,
					Map:  "Hard Rock Mine",
					Mode: "Gem Grab",
				},
			},
			{
				Battle: models.Battle{
					Duration:   210,
					Mode:       "Brawl Ball",
					Result:     "defeat",
					StarPlayer: nil,
					Teams: []models.Team{
						{
							{
								Brawler: models.Brawler{
									ID:       8,
									Name:     "El Primo",
									Power:    10,
									Trophies: 700,
								},
								Name: "Player7",
								Tag:  "#VWX234",
							},
						},
					},
					Type: "friendly",
				},
				BattleTime: "20240604T005831.000Z",
				Event: struct {
					ID   int    `json:"id"`
					Map  string `json:"map"`
					Mode string `json:"mode"`
				}{
					ID:   124,
					Map:  "Pinball Dreams",
					Mode: "Brawl Ball",
				},
			},
		},
	}
)

func setupMockConfig() {}

func TestNewPlayerService(t *testing.T) {
	playerTag := "PLAYER123"
	ps := NewPlayerService(playerTag)

	expectedTag := "%23PLAYER123"
	assert.Equal(t, expectedTag, ps.PlayerTag)
}

func TestGetData(t *testing.T) {
	setupMockConfig()

	mockResponse, err := json.Marshal(sampleData)
	assert.NoError(t, err)

	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "Bearer test-api-key", r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Update the mock configuration to use the mock server URL
	config.Cfg.BattleLogEndpoint = mockServer.URL + "/battles/{playerTag}"

	playerTag := "PLAYER123"
	ps := NewPlayerService(playerTag)

	data, err := ps.GetBattleLog()
	assert.NoError(t, err)
	assert.Equal(t, sampleData, data)
}

func TestCalculatePlayerStats(t *testing.T) {
	setupMockConfig()

	playerTag := "PLAYER123"
	ps := NewPlayerService(playerTag)

	data, err := ps.CalculatePlayerStats(sampleData)
	assert.NoError(t, err)

	expectedData := []Stat{
		{Description: constant.BattleLogTotalGamesDescription, Value: 2},
		{Description: constant.BattleLogLastPlayedDescription, Value: "0 days and 23 hours ago"},
		{Description: constant.BattleLogTotalTimeDescription, Value: "0 days and 23 hours ago"},
		{Description: constant.BattleLogStarPlayerDescription, Value: float64(50)},
		{Description: constant.BattleLogTotalVictoriesDescription, Value: 1},
		{Description: constant.BattleLogTotalDefeatsDescription, Value: 1},
		{Description: constant.BattleLogTotalTiesDescription, Value: 0},
		{Description: constant.BattleLogStarPlayerVictoryDescription, Value: float64(100)},
		{Description: constant.BattleLogStarPlayerDefeatDescription, Value: float64(0)},
		{Description: constant.BattleLogVictoryStarPlayerDescription, Value: float64(100)},
		{Description: constant.BattleLogDefeatStarPlayerDescription, Value: float64(0)},
	}
	assert.Equal(t, expectedData, data)
}
