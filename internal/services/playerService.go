package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/timolung/brawl-stars-stats-service/internal/config"
)

type PlayerService struct {
	PlayerTag string
}

func NewPlayerService(playerTag string) *PlayerService {
	playerTag = "#" + playerTag
	playerTagEncoded := url.PathEscape(playerTag)
	return &PlayerService{
		PlayerTag: playerTagEncoded,
	}
}

func (bs *PlayerService) GetData() (map[string]interface{}, error) {
	url := strings.Replace(config.Cfg.BattleLogEndpoint, "{playerTag}", bs.PlayerTag, 1)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.BrawlStarsAPIKey)

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
	test, _ := json.Marshal(result)
	fmt.Println(string(test))

	return result, nil
}
