package task

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type UpdateTaskStatusDTO struct {
	Type string `json:"type" form:"type,omitempty" binding:"required"`
	Data struct {
		Todo []schema.Todo `json:"todo,omitempty" form:"todo,omitempty"`
	} `json:"data,omitempty" form:"data,omitempty"`
}

type CurrentTasks struct {
	Tasks []schema.Task
}

// LockTime int `json:"lockTime" form:"lockTime" binding:"required"`
type LockCurrentTasksDTO struct {
	LockTime int64 `json:"lockTime" form:"lockTime" binding:"required"`
	Current CurrentTasks `json:"current" form:"current" binding:"required"`
}

type CompleteLockTask struct {
	CompletedAt int64 `json:"completedAt" form:"completedAt" binding:"required"`
}