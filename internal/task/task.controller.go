package task

import (
	"net/http"

	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/database"

	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
)

type controllerInterface interface {
	addCurrentTask(c *gin.Context)
	getCurrentTasks(c *gin.Context)
	updateCurrentTasks(c *gin.Context)
	getTaskDetail(c *gin.Context)
	updateTaskDetail(c *gin.Context)
}

type controller struct {
	service *Service
}

func (tc *controller) addCurrentTask(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	var dto AddCurrentTaskDto

	err := c.Bind(&dto)
	if err != nil {
		shared.FormError(c, err.Error())
		return
	}

	c.JSON(200, profile)
}

func (tc *controller) getCurrentTasks(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret, err := tc.service.getCurrentTasks(database.GetUUIDFromPK(profile.PK))
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (tc *controller) updateCurrentTasks(c *gin.Context) {
	//
}

func (tc *controller) getTaskDetail(c *gin.Context) {
	// 본인 현재 디테일 내용은 기기에서 관리된다. 과거 기록을 보거나 타인 기록 가져올 때 사용한다.
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func (tc *controller) updateTaskDetail(c *gin.Context) {
	// Update
}

func newController(service *Service) controllerInterface {
	return &controller{service: service}
}

func initController(c *gin.RouterGroup, service *Service) {
	tc := newController(service)
	c.GET("", tc.getCurrentTasks)
	c.POST("", tc.addCurrentTask)
}
