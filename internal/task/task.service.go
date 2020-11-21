package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
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

func getTasksAndMeta(table *dynamo.Table, userPK string) (*[]schema.Task, *schema.Meta, error) {
	var tasks []schema.Task

	err := table.Get("PK", userPK).
		Range("SK", dynamo.BeginsWith, "TASK#").
		Project("Priority", "Title", "CreatedAt", "CompletedAt").
		Filter("attribute_exists(Priority)").
		All(&tasks)
	if err != nil {
		return nil, nil, err
	}

	var meta schema.Meta

	err = table.Get("PK", userPK).Range("SK", dynamo.Equal, "TASK#meta").One(&meta)
	if err != nil {
		return nil, nil, err
	}

	meta.Percent = transformUtil.GetRecordPercent(meta.InComplete, meta.Complete)
	return &tasks, &meta, nil
}

func (s *Service) lockCurrentTasks(userPK string, dto *LockCurrentTasksDTO) (*[]schema.Task, *schema.MetaSchema, error) {
	batch := s.DB.CoreTable.Batch("PK", "SK").Write()

	var items []interface{}

	for _, task := range dto.Current.Tasks {
		base := schema.TaskSchema{
			Key:  schema.Key{
				PK: userPK,
				SK: database.GetTaskSK(task.Priority),
			},
			Task: schema.Task{
				Todos:            task.Todos,
				Title:            task.Title,
				Priority:         task.Priority,
				Memo:             task.Memo,
				Where:            task.Where,
				WillStart:        task.WillStart,
				EstimatedMinutes: task.EstimatedMinutes,
				CompletedAt:      task.CompletedAt,
				CreatedAt:        task.CreatedAt,
			},
		}
		items = append(items, base)
	}



	meta := schema.MetaSchema{
		Key: schema.Key{
			PK: userPK,
			SK: "TASK#meta",
		},
		Meta: schema.Meta{
			InComplete: len(dto.Current.Tasks),
			Complete:   0,
			LockTime:   dto.LockTime,
			Percent: 0.0,
		},
	}
	_, err := batch.Put(items...).Put(meta).Run()
	if err != nil {
		return nil, nil, err
	}

	return &dto.Current.Tasks, &meta, nil
}

func (s *Service) getCurrentTasks(userPK string) (*[]schema.Task, *schema.Meta, error) {
	return getTasksAndMeta(&s.DB.CoreTable, userPK)
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
	tasks, meta, err := getTasksAndMeta(&s.DB.CoreTable, userPK)
	if err != nil {
		return nil, err
	}

	var keys []dynamo.Keyed

	record := schema.Record{
		Score: 0,
		Tasks: *tasks,
		Meta: *meta,
	}

	recordSchema := schema.RecordSchema{
		Key:    schema.Key{
			PK: userPK,
			SK: database.GetRecordSK(time.Now()),
		},
		Record: record,
	}

	for _, task := range *tasks {
		key := dynamo.Keys{userPK, database.GetTaskSK(task.Priority)}

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
