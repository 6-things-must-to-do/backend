package services

type UserServicesInterface interface {
	getLatestTasks(email string)
}

type UserServices struct {}