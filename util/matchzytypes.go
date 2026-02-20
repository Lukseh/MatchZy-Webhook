package util

type Event string

const (
	EventSeriesInit Event = "series_start"
)

type MatchZyRes struct {
	Event Event `json:"event"`
	SeriesInit
}

type Team1 struct {
	Id   string `json:"id"`
	Name string `json:"name" str:"TEAM1"`
}
type Team2 struct {
	Id   string `json:"id"`
	Name string `json:"name" str:"TEAM2"`
}

type SeriesInit struct {
	MatchID int   `json:"matchid" str:"MATCHID"`
	NumMaps int   `json:"num_maps" str:"MAPS"`
	Team1   Team1 `json:"team1"`
	Team2   Team2 `json:"team2"`
}
