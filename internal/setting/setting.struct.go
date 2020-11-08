package setting

import "github.com/6-things-must-to-do/server/internal/shared/database"

type setTaskAlertDto struct {
	database.TaskAlertSetting
}

type userWithSetting struct {
	PK               string `json:"-"`
	SK               string `json:"-"`
	TaskAlertSetting database.TaskAlertSetting `json:"taskAlertSetting" dynamo:",set"`
}
