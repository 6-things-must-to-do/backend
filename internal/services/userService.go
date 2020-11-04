package services

import (
	"github.com/guregu/dynamo"
)

type userServicesInterface interface {
	getLatestTasks(email string)
	getUserInfo(appId string, email string)
}

type UserServices struct {
	coreTable dynamo.Table
}
