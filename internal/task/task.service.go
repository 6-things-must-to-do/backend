package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	uuid2 "github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
)

// ServiceInterface ...
type ServiceInterface interface {
	getCurrentTasks(userID string) ([]schema.Task, error)
}

// Service ...
type Service struct {
	DB *database.DB
}

func (s *Service) getCurrentTasks(uid string) ([]schema.Task, error) {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return nil, err
	}

	ret := make([]schema.Task, 0)

	err = s.DB.CoreTable.Get("PK", database.GetUserPK(uuid)).Range("SK", dynamo.BeginsWith, "TASK#").All(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) addTask(uid string, dto *AddCurrentTaskDto) {
	//
}

func (s *Service) updateTaskDetail(uid string, index int, task *schema.Task) error {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return err
	}

	uTarget := s.DB.CoreTable.Update("PK", database.GetUserPK(uuid)).Range("SK", database.GetTaskSk(index))

	//
	uTarget.Set("CompletedAt", task.CompletedAt).Set("EstimatedMinutes", task.EstimatedMinutes).Set("Memo", task.Memo).Set("Todos", task.Todos).Set("Where", task.Where)

	return nil
}

func (s *Service) updateTaskPriority(uid string, from int, to int) error {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return err
	}

	fromUpdate := s.DB.CoreTable.Update("PK", database.GetUserPK(uuid)).Range("SK", database.GetTaskSk(from)).Set("SK", database.GetTaskSk(to))
	toUpdate := s.DB.CoreTable.Update("PK", database.GetUserPK(uuid)).Range("SK", database.GetTaskSk(to)).Set("SK", database.GetTaskSk(from))

	err = s.DB.DynamoDB.WriteTx().Update(fromUpdate).Update(toUpdate).Run()
	if err != nil {
		return err
	}

	return nil
}

var cachedService *Service

// GetService ...
func GetService(DB *database.DB) *Service {
	if cachedService != nil {
		return cachedService
	}
	cachedService = &Service{DB: DB}
	return cachedService
}
