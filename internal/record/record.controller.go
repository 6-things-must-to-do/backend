package record

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerInterface interface {
	getRecordDetail(c *gin.Context)
	getDashboardData(c *gin.Context)
}

type controller struct {
	service *Service
}

func (rc *controller) getRecordDetail(c *gin.Context) {
	t := c.Param("timestamp")
	timestamp := transformUtil.ToInt(t)
	profile := middlewares.GetUserProfile(c)
	ret, err := rc.service.getRecordDetail(profile.PK, int64(timestamp))

	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (rc *controller) getDashboardData(c *gin.Context) {
	y := c.Query("year")
	m := c.Query("month")
	d := c.Query("day")

	if y == "" || m == "" || d == "" {
		shared.FormError(c, "year, month, day query required")
		return
	}

	year := transformUtil.ToInt(y)
	month := transformUtil.ToInt(m)
	day := transformUtil.ToInt(d)

	profile := middlewares.GetUserProfile(c)
	ret, err := rc.service.getRecordMetaList(profile.PK, year, month, day)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func newController(service *Service) controllerInterface {
	return &controller{service: service}
}

func initController(c *gin.RouterGroup, service *Service) {
	rc := newController(service)
	c.GET("", rc.getDashboardData)
	c.GET("/:timestamp/detail", rc.getRecordDetail)
}
