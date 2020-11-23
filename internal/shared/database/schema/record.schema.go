package schema

// Record ...
type Record struct {
	Tasks []Task `json:"tasks" form:"tasks" binding:"required"`
	LockTime int64 `json:"lockTime"`
	Score float64    `json:"score"` // LSI
	InComplete int `json:"inComplete"`
	Complete int `json:"complete"`
	Percent float64 `json:"percent"`
	RecordOpenness int `json:"-"`
	Nickname string `json:"nickname"`
	Duration int64 `json:"duration,omitempty"`
}

// RecordSchema ...
type RecordSchema struct {
	Key
	Record
}
