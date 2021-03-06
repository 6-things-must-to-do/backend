package task

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/gin-gonic/gin"
)

// InitModule ...
func InitModule(c *gin.RouterGroup, DB *database.DB) {
	service := GetService(DB)
	initController(c, service)
}
