package setting

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/guregu/dynamo"
)

type ServiceInterface interface {
	setTaskAlert(dto *setTaskAlertDto)
}

type service struct {
	DB *database.DB
}

func (s *service) setTaskAlert(user *userWithSetting) {
	m, err := dynamo.MarshalItem(user)
	if err != nil {
		panic(err)
	}

	err = s.DB.CoreTable.Update("PK", m["PK"]).Range("SK", m["SK"]).Set("TaskAlertSetting", m["TaskAlertSetting"]).Run()
	if err != nil {
		panic(err)
	}
}

func newService(DB *database.DB) *service {
	return &service{DB}
}
