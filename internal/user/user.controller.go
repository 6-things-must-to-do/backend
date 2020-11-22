package user

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

type controllerInterface interface {
	remove(c *gin.Context) // 유저 삭제
	myPage(c *gin.Context) // 유저 정보 가져오기
	getOpenness(c *gin.Context) // 공개 범위 설정 상태
	taskAlert(c *gin.Context) // 알람 관련 설정
	//getSetting(c *gin.Context)
	//getMyProfile(c *gin.Context)
	//
	//// PUT
	//updateUsername(c *gin.Context)
	//updateSetting(c *gin.Context)
	////
}

type controller struct {
	service *service
}

func (uc *controller) remove(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)
	err := uc.service.removeUser(profile.PK)

	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (uc *controller) getOpenness(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret, err := uc.service.getUserOpenness(profile.PK)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (uc *controller) myPage(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret := transformUserProfileFromProfileSchema(profile)
	c.JSON(http.StatusOK, ret)
	return
}

func (uc *controller) taskAlert(c *gin.Context) {
	var dto SetTaskAlertDTO
	err := c.Bind(&dto)
	if err != nil {
		log.Error(err)
		return
	}

	user := middlewares.GetUserProfile(c)

	setting := &ProfileWithSetting{
		PK: user.PK,
		SK: user.SK,
		TaskAlertSetting: schema.TaskAlertSetting{
			Hour:   dto.Hour,
			Minute: dto.Minute,
			Offset: dto.Offset,
		},
	}

	uc.service.setTaskAlert(setting)

	c.JSON(http.StatusOK, setting)
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(r *gin.RouterGroup, service *service) {
	c := newController(service)

	r.GET("", c.myPage)
	r.DELETE("", c.remove)
	r.PUT("/settings/alarm", c.taskAlert)
	r.GET("/settings/openness", c.getOpenness)
}
