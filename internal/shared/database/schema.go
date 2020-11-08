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
	Score int `json:"score"` // LSI
}

type Task struct {
	Key
	Todo             []Todo    `json:"todo"`
	Memo             string    `json:"memo"`
	Where            string    `json:"where"`
	WillStart        time.Time `json:"willStart"`
	EstimatedMinutes int       `json:"estimatedMinutes"`
	CompletedAt      time.Time `json:"completedAt"`
	CreatedAt        time.Time `json:"createdAt"`
}

type Todo struct {
	IsCompleted bool      `json:"isCompleted"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}
