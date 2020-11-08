package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func InitModule(c *gin.RouterGroup, DB *database.DB) {
	initController(c, &service{DB: DB})
}
