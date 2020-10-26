package router

import "github.com/gin-gonic/gin"

type authRouterInterface interface {
	login(c *gin.Context)
	signup(c *gin.Context)
}

type authRouter struct {}

func (r *authRouter) login(c *gin.Context) {
	//
}

func (r *authRouter) signup(c *gin.Context) {
	//
}

func getAuthRouter() authRouterInterface {
	return new(authRouter)
}

func InitAuthRouter (router *gin.RouterGroup) {
	r := getAuthRouter()

	router.POST("/login", r.login)
	router.POST("/signup", r.signup)
}