package setting

import "github.com/6-things-must-to-do/server/internal/shared/database/schema"

type setTaskAlertDto struct {
	schema.TaskAlertSetting
}

type userWithSetting struct {
	PK               string                  `json:"-"`
	SK               string                  `json:"-"`
	TaskAlertSetting schema.TaskAlertSetting `json:"taskAlertSetting" dynamo:",set"`
}
