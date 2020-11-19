package schema

// Record ...
type Record struct {
	Score int    `json:"score"` // LSI
	Tasks []Task `json:"tasks" form:"tasks" binding:"required"`
}

// RecordSchema ...
type RecordSchema struct {
	Key
	Record
}
