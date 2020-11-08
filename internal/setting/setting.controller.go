package setting

import (
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

type controllerInterface interface {
	taskAlert(c *gin.Context)
}

type controller struct {
	service *service
}

func (sc *controller) taskAlert(c *gin.Context) {
	var dto setTaskAlertDto
	err := c.Bind(&dto)
	if err != nil {
		log.Error(err)
		return
	}

	user := middlewares.GetUserProfile(c)

	setting := &userWithSetting{
		PK: user.PK,
		SK: user.SK,
		TaskAlertSetting: database.TaskAlertSetting{
			Hour:   dto.Hour,
			Minute: dto.Minute,
			Offset: dto.Offset,
		},
	}

	sc.service.setTaskAlert(setting)

	c.JSON(http.StatusOK, setting)
}

func newController(service *service) controllerInterface {
	return &controller{service}
}

func initController(router *gin.RouterGroup, service *service) {
	c := newController(service)
	router.PUT("/task", c.taskAlert)
}
