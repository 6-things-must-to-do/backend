package auth

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func InitModule(router *gin.RouterGroup, DB *database.DB) {
	serAuth := newService(DB)
	initController(router, serAuth)
}
