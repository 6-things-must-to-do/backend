package record

import (
	"time"

	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
)

type createRecordDto struct {
	Tasks []schema.Task `json:"tasks" form:"tasks" binding:"required"`
	Date  time.Time     `json:"date" form:"date" binding:"required"`
}
