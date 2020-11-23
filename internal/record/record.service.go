package record

import (
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	"github.com/guregu/dynamo"
	"time"

	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
)


// Service ...
type Service struct {
	DB *database.DB
}

func (s *Service) getRecordMetaList(userPK string, timestamp int64) (*[]Meta, error) {

	date := time.Unix(transformUtil.ToUnixTimestamp(timestamp), 0)

	maxDate, minDate := func () (time.Time, time.Time) {
		days := time.Hour * 24
		scopeDay := 3 * days
		max := date.Add(scopeDay)
		min := date.Add(-scopeDay)
		return max, min
	}()

	maxSK := database.RecordSKFactoryByJSTimestamp(transformUtil.GetJSUnixTimestampFromTime(maxDate), "day")
	minSK := database.RecordSKFactoryByJSTimestamp(transformUtil.GetJSUnixTimestampFromTime(minDate), "day")


	var records []schema.RecordSchema
	err := s.DB.CoreTable.
		Get("PK", userPK).
		Range("SK", dynamo.Between, minSK, maxSK).
		All(&records)

	if err != nil {
		return nil, err
	}

	var ret []Meta
	for _, rec := range records {
		t := transformUtil.GetTimeFromJSUnixTimestamp(rec.LockTime)
		meta := Meta{
			Year:     t.Year(),
			Month:    int(t.Month()),
			Day:      t.Day(),
			Score:    rec.Score,
			Percent:  rec.Percent,
			LockTime: rec.LockTime,
			DayOfYear: t.YearDay(),
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
