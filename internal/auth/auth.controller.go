package auth

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controllerInterface interface {
	login(c *gin.Context)
}

type controller struct {
	service *service
}

func (ac *controller) login(c *gin.Context) {
	email := c.PostForm("email")
	provider := c.PostForm("provider")
	id := c.PostForm("id")
	nickname := c.PostForm("nickname")
	profileImage := c.PostForm("profileImage")

	form := &loginForm{email, provider, id, nickname}
	err := loginFormValidator(form)
	if err != nil {
		shared.FormError(c, err.Error())
	}

	param := &getOrCreateUserParam{
		appId:        id,
		email:        email,
		provider:     provider,
		nickname:     nickname,
		profileImage: profileImage,
	}

	user, err := ac.service.getOrCreateUser(param)
	if err != nil {
		panic(err)
	}

	token := ac.service.getJwtToken(user.AppID)

	c.JSON(http.StatusOK , gin.H{"token": token})
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(router *gin.RouterGroup, serAuth *service) {
	c := newController(serAuth)
	router.POST("/login", c.login)
}
