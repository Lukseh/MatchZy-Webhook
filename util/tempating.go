package util

type Templates struct {
	Events  map[Event]string  `json:"events"`
	Reasons map[string]string `json:"reasons"`
}
