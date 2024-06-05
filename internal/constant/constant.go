package constant

const (
	BattleLogEndpoint   = "https://bsproxy.royaleapi.dev/v1/players/{playerTag}/battlelog"
	ClubMembersEndpoint = "https://bsproxy.royaleapi.dev/v1/clubs/{clubTag}/members"

	// battle log descriptions
	BattleLogTotalGamesDescription        = "Total number of games played"
	BattleLogLastPlayedDescription        = "Last time a game was played"
	BattleLogTotalTimeDescription         = "When the earliest game was played"
	BattleLogStarPlayerDescription        = "Percent of games as star player"
	BattleLogTotalVictoriesDescription    = "Number of games won"
	BattleLogTotalDefeatsDescription      = "Number of games lost"
	BattleLogTotalTiesDescription         = "Number of games tied"
	BattleLogStarPlayerVictoryDescription = "Percent of games as star player when won ((number of victories as star player) / victories)"
	BattleLogStarPlayerDefeatDescription  = "Percent of games as star player when lost ((number of defeats as star player) / defeats)"
	BattleLogVictoryStarPlayerDescription = "Percent of victories when star player ((number of victories as star player) / total star player)"
	BattleLogDefeatStarPlayerDescription  = "Percent of defeats when star player ((number of defeats as star player) / total star player)"
)
