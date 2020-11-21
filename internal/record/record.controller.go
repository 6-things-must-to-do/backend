package record

import (
	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	getRecord(c *gin.Context)
	getDashboardData(c *gin.Context)
}

type controller struct {
	service *Service
}

func (rc *controller) getRecord(c *gin.Context) {
	// 본인 현재 리스트는 기기에서 관리됨. 과거 기록 또는 타인 것 가져올 때 사용한다.
	// Query -> userID?, RecordUUID
}

func (rc *controller) getDashboardData(c *gin.Context) {
	//
}
func newController(service *Service) controllerInterface {
	return &controller{service: service}
}

func initController(c *gin.RouterGroup, service *Service) {
	rc := newController(service)
	c.GET("", rc.getRecord)
}
