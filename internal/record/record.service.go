package record

import (
	"time"

	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	uuid2 "github.com/gofrs/uuid"
)

// ServiceInterface ...
type ServiceInterface interface {
	getRecord(uid string, date time.Time) ([]schema.Task, error)
	createRecord(uid string, dto *createRecordDto) (*schema.RecordSchema, error)
}

// Service ...
type Service struct {
	DB *database.DB
}

func (s *Service) createRecord(uid string, dto *createRecordDto) (*schema.RecordSchema, error) {
	uuid, err := uuid2.FromString(uid)
	if err != nil {
		return nil, err
	}

	pk := database.GetUserPK(uuid)
	sk := database.GetRecordSK(dto.Date)

	record := &schema.RecordSchema{
		Key:    schema.Key{PK: pk, SK: sk},
		Record: schema.Record{Tasks: dto.Tasks},
	}

	err = s.DB.CoreTable.Put(record).Run()
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s *Service) getRecord(uid string, date time.Time) (*schema.RecordSchema, error) {
	var ret schema.RecordSchema

	return &ret, nil
}

var cachedService *Service

// GetService ...
func GetService(DB *database.DB) *Service {
	if cachedService != nil {
		return cachedService
	}
	cachedService = &Service{DB}
	return cachedService
}
