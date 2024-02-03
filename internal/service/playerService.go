package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"github.com/timolung/brawl-stars-stats-service/internal/models"
)

type PlayerService struct {
	PlayerBattleLogEndpoint string
	PlayerInfoEndpoint      string
	OAuthToken              string
}

func NewPlayerService() *PlayerService {
	_, filename, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(filename)))

	// Construct the absolute path to config.toml
	configPath := filepath.Join(rootDir, "config.toml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return &PlayerService{
		PlayerBattleLogEndpoint: viper.GetString("api.playerBattleLogEndpoint"),
		PlayerInfoEndpoint:      viper.GetString("api.playerInfoEndpoint"),
		OAuthToken:              viper.GetString("oauth.token"),
	}
}

func (s *PlayerService) GetStatsSummary(playerTag string) (string, error) {
	return "hi", nil

}

func (s *PlayerService) GetBattleLog(playerTag string) (string, error) {
	modifiedURI := strings.Replace(s.PlayerBattleLogEndpoint, "{playerTag}", playerTag, 1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", modifiedURI, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer: "+s.OAuthToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject models.PlayerBattleLog
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println(string(bodyBytes))
	fmt.Println("Status: ", resp.StatusCode)
	fmt.Printf("API Response as struct %+v\n", responseObject)
	return "hi", nil
}
