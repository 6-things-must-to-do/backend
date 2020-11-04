package user

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/database"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerInterface interface {
	// GET
	getUser(c *gin.Context)
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

func (uc *controller) getMyProfile(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)

	ret := &userProfile{
		Email:        database.GetEmailFromSK(profile.SK),
		UUID:         database.GetUUIDFromPK(profile.PK),
		ProfileImage: profile.ProfileImage,
		Nickname:     profile.Nickname,
	}

	c.JSON(http.StatusOK, ret)
}

func (uc *controller) getUser(c *gin.Context) {
	uuid := c.Param("uuid")
	profile := middlewares.GetUserProfile(c)

	var ret *userProfile

	if database.GetUUIDFromPK(profile.PK) == uuid || uuid == "my-page" {
		ret = transformUserProfileFromProfileSchema(profile)
		c.JSON(http.StatusOK, ret)
		return
	}

	ret, err := uc.service.getUserProfile(database.CreateUserPK(uuid))
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(r *gin.RouterGroup, service *service) {
	c := newController(service)

	r.GET("/:uuid", c.getUser)
}