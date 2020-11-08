package user

import "github.com/6-things-must-to-do/server/internal/shared/database"

type userProfile struct {
	Email            string                    `json:"email"`
	UUID             string                    `json:"uuid"`
	ProfileImage     string                    `json:"profileImage"`
	Nickname         string                    `json:"nickname"`
	TaskAlertSetting database.TaskAlertSetting `json:"taskAlertSetting"`
}
