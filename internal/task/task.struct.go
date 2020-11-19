package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"time"
)

type TaskList struct {
	database.Task
}

type SaveRecordDto struct {
	Tasks []database.Task `json:"tasks" form:"tasks" binding:"required"`
	Date  time.Time       `json:"date" form:"date" binding:"required"`
}

type UpdateTask struct {
	Index       int             `json:"index" form:"index" binding:"required"`
	Todos       []database.Todo `json:"todos" form:"todos"`
	CompletedAt time.Time       `json:"completedAt" form:"completedAt"`
}
