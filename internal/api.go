package internal

import (
	"github.com/6-things-must-to-do/server/internal/auth"
	"github.com/6-things-must-to-do/server/internal/setting"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/6-things-must-to-do/server/internal/task"
	"github.com/6-things-must-to-do/server/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func GetAPI() *gin.Engine {
	r := gin.Default()

	env := os.Getenv("ENV")
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	db := database.GetDB()

	api := r.Group("/api")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello!! "})
	})

	authGroup := api.Group("/auth")
	auth.InitModule(authGroup, db)

	authenticated := api.Group("")
	authenticated.Use(middlewares.AuthRequired())

	userGroup := authenticated.Group("/users")
	user.InitModule(userGroup, db)

	taskGroup := authenticated.Group("/tasks")
	task.InitModule(taskGroup, db)

	settingGroup := authenticated.Group("/settings")
	setting.InitModule(settingGroup, db)
	//{
	//	taskGroup := authenticated.Group("/tasks")
	//	router.InitTaskRouter(taskGroup)
	//
	//	socialGroup := authenticated.Group("/social")
	//	router.InitSocialRouter(socialGroup)
	//
	//	userGroup := authenticated.Group("/users")
	//	router.InitUserRouter(userGroup)
	//}
	return r
}
