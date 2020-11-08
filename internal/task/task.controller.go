package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerInterface interface {
	getTask(c *gin.Context)
}

type controller struct {
	service *service
}

func (tc *controller) getTask(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(c *gin.RouterGroup, service *service) {
	tc := newController(service)
	c.GET("/*date", tc.getTask)
}
