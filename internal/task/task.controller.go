package task

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

type controllerInterface interface {
	getTaskList(c *gin.Context)
	getTaskDetail(c *gin.Context)
	getDashboardData(c *gin.Context)
	saveRecord(c *gin.Context)
}

type controller struct {
	service *service
}

func (tc *controller) saveRecord(c *gin.Context) {
	var dto SaveRecordDto
	err := c.Bind(&dto)
	if err != nil {
		log.Error(err)
		return
	}


}

func (tc *controller) getDashboardData(c *gin.Context) {
	//
}

func (tc *controller) getTaskDetail(c *gin.Context) {
	// 본인 현재 디테일 내용은 기기에서 관리된다. 과거 기록을 보거나 타인 기록 가져올 때 사용한다.


	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func (tc *controller) getTaskList(c *gin.Context) {
	// 본인 현재 리스트는 기기에서 관리됨. 과거 기록 또는 타인 것 가져올 때 사용한다.
	// Query -> userID?, RecordUUID
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(c *gin.RouterGroup, service *service) {
	tc := newController(service)
	c.PUT("/records", tc.saveRecord)
	c.GET("/records", tc.getTaskList)
	c.GET("/records/:index", tc.getTaskDetail)
}
