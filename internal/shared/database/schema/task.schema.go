package schema

// Meta ...
type Meta struct {
	InComplete int     `json:"inComplete"`
	Complete   int     `json:"complete"`
	Percent    float64 `json:"percent"`
	LockTime   int64     `json:"lockTime"`
}

// MetaSchema ...
type MetaSchema struct {
	Key
	Meta
}

// Todo ...
type Todo struct {
	IsCompleted bool   `json:"isCompleted"`
	Content     string `json:"content"`
	CreatedAt   int64    `json:"createdAt"`
}

// Task ...
type Task struct {
	Todos            []Todo `json:"todos" form:"todos" binding:"required"`
	Title            string `json:"title" form:"title" binding:"required"`
	Priority         int    `json:"priority" form:"priority" binding:"gte=0"`
	Memo             string `json:"memo,omitempty" form:"memo" dynamo:",omitempty"`
	Where            string `json:"where,omitempty" form:"where" dynamo:",omitempty"`
	WillStart        int64  `json:"willStart,omitempty" form:"willStart" dynamo:",omitempty"`
	EstimatedMinutes int    `json:"estimatedMinutes,omitempty" form:"estimatedMinutes" dynamo:",omitempty"`
	CompletedAt      int64  `json:"completedAt,omitempty" form:"completedAt" dynamo:",omitempty"`
	CreatedAt        int64  `json:"createdAt" form:"createdAt" binding:"required"`
}

// TaskSchema ...
type TaskSchema struct {
	Key
	Task
}
