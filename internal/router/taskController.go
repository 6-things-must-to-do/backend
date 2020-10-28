package router

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/gin-gonic/gin"
)

type taskRouterInterface interface {
	getTaskDetail(c *gin.Context)
	getDashboardData(c *gin.Context)
}

type taskRouter struct {}


func (r *taskRouter) getTaskDetail(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		shared.RequiredFieldError(c, "date field is required in query")
	}
}

func (r *taskRouter) getDashboardData(c *gin.Context) {
	//
}

func getTaskRouter() taskRouterInterface {
	return new(taskRouter)
}

func InitTaskRouter (router *gin.RouterGroup) {
	r := getTaskRouter()

	router.GET("", r.getTaskDetail)
	router.GET("/dashboard/:type", r.getDashboardData)
}

