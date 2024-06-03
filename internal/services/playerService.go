package services

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/timolung/brawl-stars-stats-service/internal/config"
)

type PlayerService struct {
	APIKey         string
	PlayerEndpoint string
	PlayerTag      string
}

func NewPlayerService(apiKey string, playerEndpoint string, playerTag string) *BrawlStarsService {
	return &BrawlStarsService{
		APIKey:         apiKey,
		PlayerEndpoint: playerEndpoint,
		PlayerTag:      playerTag,
	}
}

func (bs *BrawlStarsService) GetData() (map[string]interface{}, error) {
	url := strings.Replace(config.BattleLogEndpoint, "{playerTag}", bs.PlayerEndpoint, 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+bs.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
