package config

import (
	"flag"
	"os"

	"github.com/timolung/brawl-stars-stats-service/internal/constant"
)

type AppConfig struct {
	BattleLogEndpoint string
	BrawlStarsAPIKey  string
}

var (
	Cfg = AppConfig{
		BattleLogEndpoint: constant.BattleLogEndpoint,
		BrawlStarsAPIKey:  "missing",
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
