package record

import (
	"time"

	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
)

// ServiceInterface ...
type ServiceInterface interface {
	getRecord(uid string, date time.Time) ([]schema.Task, error)
}

// Service ...
type Service struct {
	DB *database.DB
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
