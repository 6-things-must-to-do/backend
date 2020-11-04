package database

import "time"

type Key struct {
	PK string
	SK string
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
	Score int // LSI
}

type Task struct {
	Key
	Todo             []Todo
	Memo             string
	Where            string
	WillStart        time.Time
	EstimatedMinutes int
	CompletedAt      time.Time
	CreatedAt        time.Time
}

type Todo struct {
	IsCompleted bool
	Content     string
	CreatedAt   time.Time
}
