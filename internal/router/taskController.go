package router

import (
	"github.com/gin-gonic/gin"
)

type taskRouterInterface interface {
	getLatestTasks(c *gin.Context)
	getTaskDetail(c *gin.Context)
	getDashboardData(c *gin.Context)
}

type taskRouter struct {}

func (r *taskRouter) getLatestTasks(c *gin.Context) {
	//
}

func (r *taskRouter) getTaskDetail(c *gin.Context) {
	//
}

func (r *taskRouter) getDashboardData(c *gin.Context) {
	//
}

func getTaskRouter() taskRouterInterface {
	return new(taskRouter)
}

func InitTaskRouter (router *gin.RouterGroup) {
	r := getTaskRouter()

	router.GET("", r.getLatestTasks)
	router.GET("/:date", r.getTaskDetail)
	router.GET("/dashboard/:type", r.getDashboardData)
}

