package util

type Event string
type Side string
type Team string

type Teams map[Team]Team_t

var TeamDictionary = make(Teams)

const (
	EventSeriesStart Event = "series_start"
	EventMapResult   Event = "map_result"
)
const (
	CT   Side = "ct"
	TT   Side = "t"
	SPEC Side = "spec"
)

const (
	Team1 Team = "team1"
	Team2 Team = "team2"
	Spec  Team = "spec"
)

type MatchZyRes struct {
	Event   Event  `json:"event"`
	MatchID uint64 `json:"matchid" str:"MATCHID"`
	SeriesStart
	MapResult
}
type Team_t struct {
	Team1    Team1_t
	Team2    Team2_t
	TeamSpec Spec_t
}

type PlayerStats struct {
	/*
			kills
		integer >= 0
		The number of kills the player had.

		deaths
		integer >= 0
		The number of deaths the player had.

		assists
		integer >= 0
		The number of assists the player had.

		flash_assists
		integer >= 0
		The number of flashbang assists the player had.

		team_kills
		integer >= 0
		The number of team-kills the player had.

		suicides
		integer >= 0
		The number of suicides the player had.

		damage
		integer >= 0
		The total amount of damage the player dealt.

		utility_damage
		integer >= 0
		The total amount of damage the player dealt via utility.

		enemies_flashed
		integer >= 0
		The number of enemies flashed by the player.

		friendlies_flashed
		integer >= 0
		The number of teammates flashed by the player.

		knife_kills
		integer >= 0
		The number kills the player had with a knife.

		headshot_kills
		integer >= 0
		The number kills the player had that were headshots.

		rounds_played
		integer >= 0
		The number of rounds the player started.

		bomb_defuses
		integer >= 0
		The number of times the player defused the bomb.

		bomb_plants
		integer >= 0
		The number of times the player planted the bomb.

		1k
		integer >= 0
		The number of rounds where the player killed 1 opponent.

		2k
		integer >= 0
		The number of rounds where the player killed 2 opponents.

		3k
		integer >= 0
		The number of rounds where the player killed 3 opponents.

		4k
		integer >= 0
		The number of rounds where the player killed 4 opponents.

		5k
		integer >= 0
		The number of rounds where the player killed 5 opponents.

		1v1
		integer >= 0
		The number of 1v1s the player won.

		1v2
		integer >= 0
		The number of 1v2s the player won.

		1v3
		integer >= 0
		The number of 1v3s the player won.

		1v4
		integer >= 0
		The number of 1v4s the player won.

		1v5
		integer >= 0
		The number of 1v5s the player won.

		first_kills_t
		integer >= 0
		The number of rounds where the player had the first kill in the round while playing the T side.

		first_kills_ct
		integer >= 0
		The number of rounds where the player had the first kill in the round while playing the CT side.

		first_deaths_t
		integer >= 0
		The number of rounds where the player was the first to die in the round while playing the T side.

		first_deaths_ct
		integer >= 0
		The number of rounds where the player was the first to die in the round while playing the CT side.

		trade_kills
		integer >= 0
		The number of times the player got a kill in a trade.

		kast
		integer >= 0
		The number of rounds where the player (k)illed a player, had an (a)ssist, (s)urvived or was (t)raded.

		score
		integer >= 0
		The in-game "score" of the player.

		mvp
		integer >= 0
		The number of times the player was elected the round MVP.
	*/
}

type Player struct {
	Steamid string      `json:"steamid" str:"STEAMID"`
	Name    string      `json:"name" str:"NICKNAME"`
	Stats   PlayerStats `json:"stats"`
}

type Team1_t struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty" str:"TEAM1"`
	SeriesScore  uint8    `json:"series_score,omitempty" str:"MAPS1"`
	Score        uint8    `json:"score,omitempty" str:"SCORE1"`
	ScoreCT      uint8    `json:"score_ct,omitempty" str:"SCORE_CT1"`
	ScoreTT      uint8    `json:"score_t,omitempty" str:"SCORE_TT1"`
	Players      []Player `json:"players,omitempty" str:"PLAYERS1"`
	Side         string   `json:"side,omitempty" str:"SIDE1"`
	StartingSide string   `json:"starting_side,omitempty" str:"STARTSIDE1"`
}
type Team2_t struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty" str:"TEAM2"`
	SeriesScore  uint8    `json:"series_score,omitempty" str:"MAPS2"`
	Score        uint8    `json:"score,omitempty" str:"SCORE2"`
	ScoreCT      uint8    `json:"score_ct,omitempty" str:"SCORE_CT2"`
	ScoreTT      uint8    `json:"score_t,omitempty" str:"SCORE_TT2"`
	Players      []Player `json:"players,omitempty" str:"PLAYERS2"`
	Side         string   `json:"side,omitempty" str:"SIDE2"`
	StartingSide string   `json:"starting_side,omitempty" str:"STARTSIDE2"`
}
type Spec_t struct {
}

type SeriesStart struct {
	NumMaps uint8   `json:"num_maps,omitempty" str:"MAPS"`
	Team1   Team1_t `json:"team1,omitempty"`
	Team2   Team2_t `json:"team2,omitempty"`
}

type MapResult struct {
	MapNumber uint8 `json:"map_number,omitempty" str:"MAP"`
	Winner    struct {
		Side Side `json:"side"`
		Team Team `json:"team"`
	} `json:"winner,omitempty"`
}
