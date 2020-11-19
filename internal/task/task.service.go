package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	uuid2 "github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
)

type ServiceInterface interface {
	getCurrentTasks(userID string) ([]database.Task, error)
	createRecord(dto *SaveRecordDto) (*database.Task, error)
}

type service struct {
	DB *database.DB
}

func (s *service) createRecord(uid string, dto *SaveRecordDto) (*database.Record, error) {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return nil, err
	}

	pk := database.GetUserPK(uuid)
	sk := database.GetRecordSK(dto.Date)

	record := &database.Record{
		Key:   database.Key{PK: pk, SK: sk},
		Tasks: dto.Tasks,
	}

	err = s.DB.CoreTable.Put(record).Run()
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s *service) getCurrentTasks(uid string) ([]database.Task, error) {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return nil, err
	}

	var ret []database.Task

	err = s.DB.CoreTable.Get("PK", database.GetUserPK(uuid)).Range("SK", dynamo.BeginsWith, "TASK#").All(ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *service) updateTaskDetail(uid string, index int, task *database.Task) error {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return err
	}

	uTarget := s.DB.CoreTable.Update("PK", database.GetUserPK(uuid)).Range("SK", database.GetTaskSk(index))

	//
	uTarget.Set("CompletedAt", task.CompletedAt).Set("EstimatedMinutes", task.EstimatedMinutes).Set("Memo", task.Memo).Set("Todos", task.Todos).Set("Where", task.Where)

	return nil
}

func (s *service) updateTaskPriority(uid string, from int, to int) error {
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

var cachedService *service

func GetTaskService() *service {
	if cachedService != nil {
		return cachedService
	}
	cachedService = new(service)
	return cachedService
}
