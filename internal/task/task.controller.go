package task

import (
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"

	"net/http"

	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	getCurrentTasks(c *gin.Context)
	getTaskDetail(c *gin.Context)
	lockCurrentTasks(c *gin.Context)
	clearCurrentTasks(c *gin.Context)
}

type controller struct {
	service *Service
}

func (tc *controller) getCurrentTasks(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret, err := tc.service.getCurrentTasks(profile.PK)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (tc *controller) getTaskDetail(c *gin.Context) {
	rIndex := c.Param("index")
	targetIndex := transformUtil.ToInt(rIndex)

	profile := middlewares.GetUserProfile(c)

	ret, err := tc.service.getTaskDetail(profile.PK, targetIndex)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (tc *controller) lockCurrentTasks(c *gin.Context) {
	// 현재 상태를 최종으로 저장해둔다.
	var dto LockCurrentTasksDTO
	err := c.ShouldBind(&dto)
	if err != nil {
		shared.FormError(c, err.Error())
		return
	}

	profile := middlewares.GetUserProfile(c)

	ret, err := tc.service.lockCurrentTasks(profile.PK, &dto)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (tc *controller) clearCurrentTasks(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret, err := tc.service.clearCurrentTasks(profile.PK)
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
	tc := newController(service)
	c.GET("", tc.getCurrentTasks)
	c.POST("", tc.lockCurrentTasks)
	c.DELETE("", tc.clearCurrentTasks)
	c.GET("/:index", tc.getTaskDetail)
}
