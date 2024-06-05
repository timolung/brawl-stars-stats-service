package services

import (
	"encoding/json"
	"net/url"
	"sort"
	"sync"

	"github.com/timolung/brawl-stars-stats-service/internal/config"
	"github.com/timolung/brawl-stars-stats-service/internal/models"
	"github.com/timolung/brawl-stars-stats-service/internal/utils"
)

var wg sync.WaitGroup

type ClubMember struct {
	Name           string
	LastPlayed     string
	EarliestPlayed string
	Role           string
}

type ClubMemberTime struct {
}

type ClubService struct {
	ClubTag string
}

func NewClubService(clubTag string) *ClubService {
	clubTag = "#" + clubTag
	clubTagEncoded := url.PathEscape(clubTag)
	return &ClubService{
		ClubTag: clubTagEncoded,
	}
}

func (cs *ClubService) GetClubMembersList() (models.ClubMembersResponse, error) {
	var apiResponse models.ClubMembersResponse
	data, err := utils.MakeAPIRequest(config.Cfg.ClubMembersEndpoint, "{clubTag}", cs.ClubTag)
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

func (cs *ClubService) CalculateClubMemberStats(apiResponse models.ClubMembersResponse) ([]ClubMember, error) {
	if len(apiResponse.Items) == 0 {
		return []ClubMember{}, nil
	}

	clubMembers := []ClubMember{}
	playerService := NewPlayerService("none")

	for _, item := range apiResponse.Items {
		wg.Add(1)

		go func(item models.ClubMemberItem) {
			defer wg.Done()
			playerService.PlayerTag = url.PathEscape(item.Tag)
			battleLog, err := playerService.GetBattleLog()
			if err != nil {
				return
			}

			latestPlayedTimestamp := battleLog.Items[0].BattleTime
			earliestPlayedTimestamp := battleLog.Items[len(battleLog.Items)-1].BattleTime

			clubMember := ClubMember{
				Name:           item.Name,
				Role:           item.Role,
				LastPlayed:     latestPlayedTimestamp,
				EarliestPlayed: utils.CalculateDuration(earliestPlayedTimestamp),
			}

			clubMembers = append(clubMembers, clubMember)
		}(item)
	}
	wg.Wait()

	sort.Slice(clubMembers, func(i, j int) bool {
		return clubMembers[i].LastPlayed < clubMembers[j].LastPlayed
	})

	for i, clubMember := range clubMembers {
		clubMember.LastPlayed = utils.CalculateDuration(clubMember.LastPlayed)
		clubMembers[i] = clubMember
	}

	return clubMembers, nil
}
