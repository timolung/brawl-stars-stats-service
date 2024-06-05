package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/timolung/brawl-stars-stats-service/internal/config"
)

func MakeAPIRequest(endpoint, param string, replaceParam string) (map[string]interface{}, error) {
	url := strings.Replace(endpoint, param, replaceParam, 1)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.Cfg.BrawlStarsAPIKey)

	log.Printf("Begin Making External API Request")

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

	log.Printf("Successfully Made External API Request")

	return result, nil
}
