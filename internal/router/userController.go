package router

import "github.com/gin-gonic/gin"

type userRouterInterface interface {
	// GET
	getUser(c *gin.Context)
	getSetting(c *gin.Context)

	// PUT
	updateUsername(c *gin.Context)
	updateSetting(c *gin.Context)
	//
}

type userRouter struct {}

func (r *userRouter) getUser(c *gin.Context) {
	//
}

func (r *userRouter) getSetting(c *gin.Context) {
	//
}

func (r *userRouter) updateUsername(c *gin.Context) {
	//
}

func (r *userRouter) updateSetting(c *gin.Context) {
	//
}

func getUserRouter() userRouterInterface {
	return new(userRouter)
}

func InitUserRouter(router *gin.RouterGroup) {
	r := getUserRouter()

	router.GET("/settings", r.getSetting)
	router.PUT("/settings", r.updateSetting)

	router.GET("", r.getUser)
	router.PUT("", r.updateUsername)
}