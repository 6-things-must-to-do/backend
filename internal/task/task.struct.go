package task

import (
	"time"

	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
)

// TaskList ...
type TaskList struct {
	schema.Task
}

// UpdateTask ...
type UpdateTask struct {
	Index       int           `json:"index" form:"index" binding:"required"`
	Todos       []schema.Todo `json:"todos" form:"todos"`
	CompletedAt time.Time     `json:"completedAt" form:"completedAt"`
}

// AddCurrentTaskDto ...
type AddCurrentTaskDto struct {
	Index int `json:"index" form:"index" binding:"required"`
	schema.Task
}
