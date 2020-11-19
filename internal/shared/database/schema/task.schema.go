package schema

// Todo ...
type Todo struct {
	IsCompleted bool   `json:"isCompleted"`
	Content     string `json:"content"`
	CreatedAt   int    `json:"createdAt"`
}

// Task ...
type Task struct {
	Todos            []Todo `json:"todos" form:"todos" binding:"required"`
	Memo             string `json:"memo,omitempty" form:"memo"`
	Where            string `json:"where,omitempty" form:"where"`
	WillStart        int    `json:"willStart,omitempty" form:"willStart"`
	EstimatedMinutes int    `json:"estimatedMinutes,omitempty" form:"estimatedMinutes"`
	CompletedAt      int    `json:"completedAt,omitempty"`
	CreatedAt        int    `json:"createdAt" form:"createdAt" binding:"required"`
}

// TaskSchema ...
type TaskSchema struct {
	Key
	Task
}
