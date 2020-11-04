package internal

import (
	"github.com/6-things-must-to-do/server/internal/auth"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
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

	db := database.InitDB(true)

	api := r.Group("/api")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Hello!! "})
	})

	authGroup := api.Group("/auth")
	auth.InitModule(authGroup, db)

	//authenticated := api.Group("")
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
