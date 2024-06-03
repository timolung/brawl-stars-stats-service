package services

type BrawlStarsService struct {
	APIKey         string
	PlayerEndpoint string
	PlayerTag      string
}

func NewBrawlStarsService(apiKey string, playerEndpoint string, playerTag string) *BrawlStarsService {
	return &BrawlStarsService{
		APIKey:         apiKey,
		PlayerEndpoint: playerEndpoint,
		PlayerTag:      playerTag,
	}
}
