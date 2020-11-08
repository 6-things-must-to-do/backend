package auth

import (
	"github.com/6-things-must-to-do/server/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

type controllerInterface interface {
	login(c *gin.Context)
}

type controller struct {
	service *service
}

func (ac *controller) login(c *gin.Context) {
	var dto loginDto
	err := c.Bind(&dto)
	if err != nil {
		log.Error(err)
		return
	}

	err = loginFormValidator(&dto)
	if err != nil {
		shared.FormError(c, err.Error())
		return
	}

	user, err := ac.service.getOrCreateUser(&dto)
	if err != nil {
		panic(err)
	}

	token := ac.service.getJwtToken(user.PK)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func newController(service *service) controllerInterface {
	return &controller{service: service}
}

func initController(router *gin.RouterGroup, serAuth *service) {
	c := newController(serAuth)
	router.POST("/login", c.login)
}
