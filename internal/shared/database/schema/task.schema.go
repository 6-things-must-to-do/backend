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
	Title			 string `json:"title" form:"title" binding:"required"`
	Index			 int 	`json:"index" form:"index" binding:"gte=0"`
	Memo             string `json:"memo,omitempty" form:"memo"`
	Where            string `json:"where,omitempty" form:"where"`
	WillStart        int64    `json:"willStart,omitempty" form:"willStart"`
	EstimatedMinutes int    `json:"estimatedMinutes,omitempty" form:"estimatedMinutes"`
	CompletedAt      int64    `json:"completedAt,omitempty" form:"completedAt"`
	CreatedAt        int64    `json:"createdAt" form:"createdAt" binding:"required"`
}

// TaskSchema ...
type TaskSchema struct {
	Key
	Task
}
