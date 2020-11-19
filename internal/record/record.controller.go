package record

import (
	"net/http"

	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type controllerInterface interface {
	getRecord(c *gin.Context)
	createRecord(c *gin.Context)

	getDashboardData(c *gin.Context)
}

type controller struct {
	service *Service
}

func (rc *controller) createRecord(c *gin.Context) {
	var dto createRecordDto
	err := c.Bind(&dto)
	if err != nil {
		log.Error(err)
		shared.FormError(c, err.Error())
	}

	profile := middlewares.GetUserProfile(c)

	ret, err := rc.service.createRecord(database.GetUUIDFromPK(profile.PK), &dto)

	if err != nil {
		log.Error(err)
		shared.BadRequestError(c, err.Error())
	}

	c.JSON(http.StatusOK, ret)
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
	c.POST("", rc.createRecord)
}
