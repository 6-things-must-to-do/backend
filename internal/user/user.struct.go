package user

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type userProfile struct {
	Email            string                  `json:"email"`
	UUID             string                  `json:"uuid"`
	ProfileImage     string                  `json:"profileImage,omitempty"`
	Nickname         string                  `json:"nickname"`
	TaskAlertSetting schema.TaskAlertSetting `json:"taskAlertSetting,omitempty"`
}
