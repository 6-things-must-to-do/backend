package social

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/database/schema"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	sliceUtil "github.com/6-things-must-to-do/server/internal/shared/utils/slice"
	transformUtil "github.com/6-things-must-to-do/server/internal/shared/utils/transform"
	validateUtil "github.com/6-things-must-to-do/server/internal/shared/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sc *controller) getFriendDashboard(c *gin.Context) {
	//
}

func (sc *controller) getAllLeaderboard(c *gin.Context) {
	mode := c.Query("mode")
	time := c.Query("time")

	if time == "" {
		shared.FormError(c, "query must have 'time'")
		return
	}

	jsTimestamp := transformUtil.ToInt(time)
	jsTimestamp64 := int64(jsTimestamp)


	availableMode := []interface{}{"month","week", "day", ""}
	if !sliceUtil.Includes(availableMode, mode) {
		shared.FormError(c, "invalid mod")
		return
	}

	var ret *[]schema.RankRecord
	var err error

	switch mode {
	default:
		ret, err = sc.service.getAllRankingByDay(jsTimestamp64)
	}
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": ret})
}

func (sc *controller) getFollowerList(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)
	follower, err := sc.service.getFollowerList(profile.SK)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"follower": follower, "count": len(*follower)})
}

func (sc* controller) getFollowingList(c *gin.Context) {
	profile := middlewares.GetUserProfile(c)
	following, err := sc.service.getFollowingList(profile.PK)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"following": following, "count": len(*following)})
}

func (sc *controller) getUser(c *gin.Context) {
	email := c.Param("email")
	if !validateUtil.IsEmail(email) {
		shared.FormError(c, "invalid email")
		return
	}

	ret, err := sc.service.getUser(email)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, ret)
}

func (sc *controller) follow(c *gin.Context) {
	email := c.Param("email")
	if !validateUtil.IsEmail(email) {
		shared.FormError(c, "invalid email")
		return
	}

	profile := middlewares.GetUserProfile(c)
	err := sc.service.follow(profile.PK, profile.SK, email)

	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (sc *controller) unfollow(c *gin.Context) {
	email := c.Param("email")
	if !validateUtil.IsEmail(email) {
		shared.FormError(c, "invalid email")
	}

	profile := middlewares.GetUserProfile(c)

	err := sc.service.unfollow(profile.PK, email)
	if err != nil {
		shared.BadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

type controllerInterface interface {

	getUser(c *gin.Context)
	follow(c *gin.Context)
	unfollow(c *gin.Context)
	getFollowerList(c *gin.Context)
	getFollowingList(c *gin.Context)

	getFriendDashboard(c *gin.Context)
	getAllLeaderboard(c *gin.Context)
}

type controller struct {
	service *service
}


func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(r *gin.RouterGroup, service *service) {
	c := newController(service)

	r.GET("/users/:email", c.getUser)
	r.POST("/users/:email", c.follow)
	r.DELETE("/users/:email", c.unfollow)

	r.GET("/following", c.getFollowingList)
	r.GET("/followers", c.getFollowerList)

	r.GET("/rank/all", c.getAllLeaderboard)
}
