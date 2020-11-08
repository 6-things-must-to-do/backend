package task

import "github.com/6-things-must-to-do/server/internal/shared/database"

type serviceInterface interface {
	getLatestTask(userID string) ([]database.Task, error)
}

type service struct {
	DB *database.DB
}

func (s *service) getLatestTask(userID string) ([]database.Task, error) {
	ret := []database.Task{database.Task{}}
	return ret, nil
}
