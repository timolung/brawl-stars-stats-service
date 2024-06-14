package constant

const (
	BattleLogEndpoint   = "https://bsproxy.royaleapi.dev/v1/players/{playerTag}/battlelog"
	ClubMembersEndpoint = "https://bsproxy.royaleapi.dev/v1/clubs/{clubTag}/members"

	// battle log descriptions
	BattleLogTotalGamesDescription        = "Length of game history"
	BattleLogTotalStarPlayerGames         = "Games played containing star player"
	BattleLogLastPlayedDescription        = "Last time a game was played"
	BattleLogTotalTimeDescription         = "When the earliest game was played"
	BattleLogStarPlayerDescription        = "Percent of games as star player"
	BattleLogTotalVictoriesDescription    = "Number of games won"
	BattleLogTotalDefeatsDescription      = "Number of games lost"
	BattleLogTotalTiesDescription         = "Number of games tied"
	BattleLogStarPlayerVictoryDescription = "Percent of games you carried to victory"
	BattleLogStarPlayerDefeatDescription  = "Percent of games you carried but still lost"
	BattleLogVictoryStarPlayerDescription = "Percent of victories when star player ((number of victories as star player) / total star player)"
	BattleLogDefeatStarPlayerDescription  = "Percent of defeats when star player ((number of defeats as star player) / total star player)"

	// showDown vicotry and loss parameters
	SoloShowDownVictoryLowerBound = 3
	SoloShowDownDefeatUpperBound  = 6
	DuoShowdownVictoryLowerBound  = 2
	DuoShowdownDefeatUpperBound   = 4
)
