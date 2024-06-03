package config

import "os"

var (
	Cfg               *Config
	BattleLogEndpoint = "https://bsproxy.royaleapi.dev/v1/players/{playerTag}/battlelog"
)

type Config struct {
	BrawlStarsAPIKey string
	// Other configuration parameters...
}

func LoadConfig() {
	// Load configuration from environment variables or config files
	apiKey := os.Getenv("BRAWL_STARS_API_KEY")

	Cfg = &Config{
		BrawlStarsAPIKey: apiKey,
	}
}
