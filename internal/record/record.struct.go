package record

type Meta struct {
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
	DayOfYear int `json:"dayOfYear"`
	Score float64 `json:"score"`
	Percent float64 `json:"percent"`
	LockTime int64	`json:"lockTime"`
}
