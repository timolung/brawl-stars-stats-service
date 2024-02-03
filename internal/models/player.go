package models

type Event struct {
	ID   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type BattlePlayer struct {
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	Brawler struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Power    int    `json:"power"`
		Trophies int    `json:"trophies"`
	} `json:"brawler"`
}

type Battle struct {
	Mode         string           `json:"mode"`
	Type         string           `json:"type"`
	Rank         int              `json:"rank"`
	TrophyChange int              `json:"trophyChange"`
	Teams        [][]BattlePlayer `json:"teams"`
}

type BattleEvent struct {
	BattleTime string `json:"battleTime"`
	Event      Event  `json:"event,omitempty"`
	Battle     Battle `json:"battle,omitempty"`
}

type PlayerBattleLog struct {
	Items []struct {
		BattleEvent BattleEvent
	} `json:"items"`
	Paging struct {
		Cursors struct {
		} `json:"cursors"`
	} `json:"paging"`
}
