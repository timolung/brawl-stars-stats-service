package models

type Brawler struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}

type Player struct {
	Brawler Brawler `json:"brawler"`
	Name    string  `json:"name"`
	Tag     string  `json:"tag"`
}

type Team []Player

type Battle struct {
	Duration   int     `json:"duration"`
	Mode       string  `json:"mode"`
	Rank       int     `json:"rank"`
	Result     string  `json:"result"`
	StarPlayer *Player `json:"starPlayer"`
	Teams      []Team  `json:"teams"`
	Type       string  `json:"type"`
}

type BattleLogItem struct {
	Battle     Battle `json:"battle"`
	BattleTime string `json:"battleTime"`
	Event      struct {
		ID   int    `json:"id"`
		Map  string `json:"map"`
		Mode string `json:"mode"`
	} `json:"event"`
}

type BattleLogResponse struct {
	Items  []BattleLogItem `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}

type Icon struct {
	Id int `json:"id"`
}

type ClubMemberItem struct {
	Icon      Icon   `json:"battle"`
	Name      string `json:"name"`
	NameColor string `json:"nameColor"`
	Role      string `json:"role"`
	Tag       string `json:"tag"`
	Trophies  int    `json:"trophies"`
}

type ClubMembersResponse struct {
	Items  []ClubMemberItem `json:"items"`
	Paging struct {
		Cursors struct{} `json:"cursors"`
	} `json:"paging"`
}
