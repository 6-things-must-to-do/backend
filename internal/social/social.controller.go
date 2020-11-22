package social

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/6-things-must-to-do/server/internal/shared/middlewares"
	validateUtil "github.com/6-things-must-to-do/server/internal/shared/utils/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sc *controller) getFriendDashboard(c *gin.Context) {
	//
}

func (sc *controller) getLeaderboard(c *gin.Context) {
	//
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
	//
}

type controllerInterface interface {

	getUser(c *gin.Context)
	follow(c *gin.Context)
	unfollow(c *gin.Context)

	getFollowerList(c *gin.Context)
	getFollowingList(c *gin.Context)
	getFriendDashboard(c *gin.Context)
	getLeaderboard(c *gin.Context)
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

	r.GET("/leaderboard", c.getLeaderboard)
}
