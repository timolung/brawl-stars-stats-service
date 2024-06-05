package config

import (
	"flag"
	"os"

	"github.com/timolung/brawl-stars-stats-service/internal/constant"
)

type AppConfig struct {
	BattleLogEndpoint   string
	BrawlStarsAPIKey    string
	ClubMembersEndpoint string
}

var (
	Cfg = AppConfig{
		BattleLogEndpoint:   constant.BattleLogEndpoint,
		BrawlStarsAPIKey:    "test-api-key",
		ClubMembersEndpoint: constant.ClubMembersEndpoint,
	}
)

func Configure() {
	flag.StringVar(&Cfg.BrawlStarsAPIKey, "missing", EnvVarOrString("BRAWL_STARS_API_KEY", Cfg.BrawlStarsAPIKey), "Brawl Stars API Key")
}

func EnvVarOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
