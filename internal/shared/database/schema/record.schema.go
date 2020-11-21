package schema

// Record ...
type Record struct {
	Score float64    `json:"score"` // LSI
	Tasks []Task `json:"tasks" form:"tasks" binding:"required"`
	Meta  Meta   `json:"meta" dynamo:",set"`
}

// RecordSchema ...
type RecordSchema struct {
	Key
	Record
}
