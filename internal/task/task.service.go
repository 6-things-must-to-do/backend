package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/guregu/dynamo"
	"time"
)

// ServiceInterface ...
type ServiceInterface interface {
	getCurrentTasks(userPK string) (*[]schema.Task, error)
	getTaskDetail(userPK string, index int) (*schema.Task, error)
	lockCurrentTasks(userPK string, dto *LockCurrentTasksDTO) (*[]schema.Task, error)
	clearCurrentTasks(userPK string) error
}

// Service ...
type Service struct {
	DB *database.DB
}

func (s *Service) lockCurrentTasks(userPK string, dto *LockCurrentTasksDTO) (*[]schema.Task, error) {

	tx := s.DB.DynamoDB.WriteTx()

	for _, task := range dto.Current.Tasks {
		base := &schema.TaskSchema{
			Key:  schema.Key{
				PK: userPK,
				SK: database.GetTaskSK(task.Index),
			},
			Task: schema.Task{
				Todos:            task.Todos,
				Title:            task.Title,
				Index:            task.Index,
				Memo:             task.Memo,
				Where:            task.Where,
				WillStart:        task.WillStart,
				EstimatedMinutes: task.EstimatedMinutes,
				CompletedAt:      task.CompletedAt,
				CreatedAt:        task.CreatedAt,
			},
		}
		tx = tx.Put(s.DB.CoreTable.Put(base))
	}

	err := tx.Run()
	if err != nil {
		return nil, err
	}

	return &dto.Current.Tasks, nil
}

func (s *Service) getCurrentTasks(userPK string) (*[]schema.Task, error) {
	ret := make([]schema.Task, 0)

	err := s.DB.CoreTable.Get("PK", userPK).Range("SK", dynamo.BeginsWith, "TASK#").All(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Service) getTaskDetail(userPK string, index int) (*schema.Task, error) {
	var task schema.Task

	err := s.DB.CoreTable.Get("PK", userPK).Range("SK", dynamo.Equal, database.GetTaskSK(index)).One(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *Service) clearCurrentTasks(userPK string) (*schema.Record, error) {
	var tasks []schema.Task

	err := s.DB.CoreTable.Get("PK", userPK).Range("SK", dynamo.BeginsWith, "TASK#").All(&tasks)
	if err != nil {
		return nil, err
	}

	var keys []dynamo.Keyed

	record := schema.Record{
		Score: 0,
		Tasks: tasks,
	}
	recordSchema := schema.RecordSchema{
		Key:    schema.Key{
			PK: userPK,
			SK: database.GetRecordSK(time.Now()),
		},
		Record: record,
	}

	for _, task := range tasks {
		key := dynamo.Keys{userPK, database.GetTaskSK(task.Index)}

		keys = append(keys, key)
	}

	_, err = s.DB.CoreTable.Batch("PK", "SK").Write().Delete(keys...).Put(&recordSchema).Run()
	if err != nil {
		return nil, err
	}

	return &record, nil
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
