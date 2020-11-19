package database

import "time"

type Key struct {
	PK string `json:"-"`
	SK string `json:"-"`
}

type TaskAlertSetting struct {
	Hour   *int `json:"hour" form:"hour" binding:"required"`
	Minute *int `json:"minute" form:"minute" binding:"required"`
	Offset *int `json:"offset" form:"offset" binding:"required"`
}

type ProfileWithSetting struct {
	Profile
	TaskAlertSetting TaskAlertSetting `dynamo:",set"`
}

type Profile struct {
	Key                 // USER#uuid PROFILE#email
	Provider     string // google | apple
	AppID        string // hashedAppId
	ProfileImage string // Image URL
	Nickname     string
}

type Record struct {
	Key
	Score int    `json:"score"` // LSI
	Tasks []Task `json:"tasks" form:"tasks" binding:"required"`
}

type Task struct {
	Key
	Todos             []Todo    `json:"todo" form:"todo" binding:"required"`
	Memo             string    `json:"memo,omitempty" form:"memo"`
	Where            string    `json:"where,omitempty" form:"where"`
	WillStart        time.Time `json:"willStart,omitempty"`
	EstimatedMinutes int       `json:"estimatedMinutes,omitempty"`
	CompletedAt      time.Time `json:"completedAt,omitempty"`
	CreatedAt        time.Time `json:"createdAt" form:"createdAt" binding:"required"`
}

type Todo struct {
	IsCompleted bool      `json:"isCompleted"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}
