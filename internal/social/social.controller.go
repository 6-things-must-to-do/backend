package social

import (
"github.com/gin-gonic/gin"
)

type socialRouterInterface interface {
	// GET
	getFriendList(c *gin.Context)
	getFriendDashboard(c *gin.Context)
	getLeaderboard(c *gin.Context)
	// POST
	addFriend(c *gin.Context)
	// DELETE
	removeFriend(c *gin.Context)
}

type socialRouter struct{}

func (r *socialRouter) getFriendList(c *gin.Context) {
	//
}

func (r *socialRouter) getFriendDashboard(c *gin.Context) {
	//
}

func (r *socialRouter) addFriend(c *gin.Context) {
	//
}

func (r *socialRouter) removeFriend(c *gin.Context) {
	//
}

func (r *socialRouter) getLeaderboard(c *gin.Context) {
	//
}

func newController() socialRouterInterface {
	return new(socialRouter)
}

func initController(router *gin.RouterGroup) {
	r := newController()

	router.GET("/friends", r.getFriendList)
	router.POST("/friends", r.addFriend)
	router.GET("/friends/:username", r.getFriendDashboard)
	router.DELETE("/friends/:username", r.addFriend)
	router.GET("/leaderboard", r.getLeaderboard)
}
