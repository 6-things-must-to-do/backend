package internal

import (
	"github.com/6-things-must-to-do/server/internal/router"
	"github.com/gin-gonic/gin"
)

func GetAPI() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("", func (c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello!! "})
	})

	authGroup := api.Group("/auth")
	router.InitAuthRouter(authGroup)

	authenticated := api.Group("")
	{
		taskGroup := authenticated.Group("/tasks")
		router.InitTaskRouter(taskGroup)

		socialGroup := authenticated.Group("/social")
		router.InitSocialRouter(socialGroup)

		userGroup := authenticated.Group("/users")
		router.InitUserRouter(userGroup)
	}
	return r
}