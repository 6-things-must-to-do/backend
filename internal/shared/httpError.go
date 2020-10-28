package shared

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequiredFieldError (c *gin.Context, errString string) {
	c.JSON(http.StatusNotAcceptable, gin.H{"message": errString})
}