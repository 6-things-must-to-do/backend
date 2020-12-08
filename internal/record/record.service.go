package record

import (
	"fmt"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	"github.com/guregu/dynamo"
	"time"
)


// Service ...
type Service struct {
	DB *database.DB
}

func (s *Service) getRecordMetaList(userPK string, year int, month int, day int) (*[]Meta, error) {
	sk := database.RecordSKFactoryFromYMD(year, month)


	var records []schema.RecordSchema
	err := s.DB.CoreTable.
		Get("PK", userPK).
		Range("SK", dynamo.BeginsWith, sk).
		All(&records)

	if err != nil {
		return nil, err
	}

	var ret []Meta
	recIndex := 0
	recMaxIndex := len(records) - 1
	for dayOfMonth := 1; dayOfMonth <= day; dayOfMonth++ {
		fmt.Println(recIndex)
		rec := records[recIndex]
		t := transformUtil.GetTimeFromJSUnixTimestamp(rec.LockTime)

		meta := Meta {
			Year: year,
			Month: month,
			Day: dayOfMonth,
			Score: 0,
			Percent: 0,
			InComplete: 0,
			Complete: 0,
			LockTime: 0,
			DayOfYear: time.Date(year, time.Month(month), dayOfMonth, 0, 0, 0, 0, time.UTC).YearDay(),
		}

		if dayOfMonth == t.Day() {
			meta.Score = rec.Score
			meta.Percent = rec.Percent
			meta.LockTime = rec.LockTime
			meta.DayOfYear = t.YearDay()
			meta.InComplete = rec.InComplete
			meta.Complete = rec.Complete
			if recIndex < recMaxIndex {
				recIndex++
			}
		}

		ret = append(ret, meta)
	}

	return &ret, err
}

func (s *Service) getRecordDetail(userPK string, timestamp int64) (*schema.RecordSchema, error) {
	var ret schema.RecordSchema

	err := s.DB.CoreTable.
		Get("PK", userPK).
		Range("SK", dynamo.Equal, database.RecordSKFactoryByJSTimestamp(timestamp, "day")).
		One(&ret)

	if err != nil {
		return nil, err
	}

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
