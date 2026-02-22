package util

type Event string
type Side string
type Team string

type Reason uint64

const (
	Unknown                       Reason = 0x0
	TargetBombed                  Reason = 0x1
	TerroristsEscaped             Reason = 0x4
	CTsPreventEscape              Reason = 0x5
	EscapingTerroristsNeutralized Reason = 0x6
	BombDefused                   Reason = 0x7
	CTsWin                        Reason = 0x8
	TerroristsWin                 Reason = 0x9
	RoundDraw                     Reason = 0xA
	AllHostageRescued             Reason = 0xB
	TargetSaved                   Reason = 0xC
	HostagesNotRescued            Reason = 0xD
	TerroristsNotEscaped          Reason = 0xE
	GameCommencing                Reason = 0x10
	TerroristsSurrender           Reason = 0x11
	CTsSurrender                  Reason = 0x12
	TerroristsPlanted             Reason = 0x13
	CTsReachedHostage             Reason = 0x14
	SurvivalWin                   Reason = 0x15
	SurvivalDraw                  Reason = 0x16
)

func ReasonToString(r Reason) string {
	switch r {
	case TargetBombed:
		return "Target bombed"
	case TerroristsEscaped:
		return "Terrorists escaped"
	case CTsPreventEscape:
		return "CTs prevented escape"
	case EscapingTerroristsNeutralized:
		return "Escaping terrorists neutralized"
	case BombDefused:
		return "Bomb defused"
	case CTsWin:
		return "CTs win"
	case TerroristsWin:
		return "Terrorists win"
	case RoundDraw:
		return "Round draw"
	case AllHostageRescued:
		return "All hostages rescued"
	case TargetSaved:
		return "Target saved"
	case HostagesNotRescued:
		return "Hostages not rescued"
	case TerroristsNotEscaped:
		return "Terrorists not escaped"
	case GameCommencing:
		return "Game commencing"
	case TerroristsSurrender:
		return "Terrorists surrendered"
	case CTsSurrender:
		return "CTs surrendered"
	case TerroristsPlanted:
		return "Terrorists planted bomb"
	case CTsReachedHostage:
		return "CTs reached hostage"
	case SurvivalWin:
		return "Survival win"
	case SurvivalDraw:
		return "Survival draw"
	default:
		return "Unknown"
	}
}

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

// str tags should never have common start
type MatchZyRes struct {
	Event   Event  `json:"event"`
	MatchID uint64 `json:"matchid" str:"MATCHID"`

	Winner struct {
		Side Side `json:"side,omitempty"`
		Team Team `json:"team,omitempty" str:"WINNER"`
	} `json:"winner,omitempty"`
	MapNumber uint8   `json:"map_number,omitempty" str:"NUMMAP"`
	MapName   string  `json:"map_name,omitempty" str:"PICKEDMAP"`
	Team      Team    `json:"team,omitempty" str:"SELECTOR"`
	Side      string  `json:"side,omitempty" str:"SIDEPICKED"`
	Team1     Team1_t `json:"team1,omitempty"`
	Team2     Team2_t `json:"team2,omitempty"`

	// Some events are missing due to using only common fields
	SeriesStart
	SeriesResult
	RoundEnd
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

type SeriesResult struct {
	Team1SeriesScore uint8  `json:"team1_series_score,omitempty" str:"SERIES1"`
	Team2SeriesScore uint8  `json:"team2_series_score,omitempty" str:"SERIES2"`
	TimeUntilRestore uint16 `json:"time_until_restore,omitempty"`
}

type RoundEnd struct {
	RoundNumber uint8  `json:"round_number,omitempty" str:"ROUND"`
	RoundTime   uint64 `json:"round_time,omitempty" str:"DUROUND"`
	Reason      Reason `json:"reason,omitempty" str:"REASON"`
}
