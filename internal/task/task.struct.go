package task

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type CurrentTasks struct {
	Tasks []schema.Task
}

type LockCurrentTasksDTO struct {
	Current CurrentTasks `json:"current" form:"current" binding:"required"`
}