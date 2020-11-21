package user

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type Profile struct {
	Email            string                  `json:"email"`
	UUID             string                  `json:"uuid"`
	ProfileImage     string                  `json:"profileImage,omitempty"`
	Nickname         string                  `json:"nickname"`
	TaskAlertSetting schema.TaskAlertSetting `json:"taskAlertSetting,omitempty"`
}

type Openness struct {
	Account int `json:"account"`
	Task int `json:"task"`
	Record int `json:"record"`
}

type SetTaskAlertDTO struct {
	schema.TaskAlertSetting
}

type ProfileWithSetting struct {
	PK               string                  `json:"-"`
	SK               string                  `json:"-"`
	TaskAlertSetting schema.TaskAlertSetting `json:"taskAlertSetting" dynamo:",set"`
}
