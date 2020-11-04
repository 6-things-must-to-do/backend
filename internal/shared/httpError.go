package shared

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormError(c *gin.Context, errString string) {
	c.JSON(http.StatusNotAcceptable, gin.H{"message": "Form has error.", "field": errString})
}
