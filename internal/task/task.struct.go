package task

import "github.com/6-things-must-to-do/server/internal/shared/database"

type TaskList struct {
	database.Task
}

type SaveRecordDto struct {
	Record []database.Task `json:"record" form:"record" binding:"required"`
}
