package shared

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormError(c *gin.Context, errString string) {
	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Form has error.", "field": errString})
}

func UnAuthenticatedError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Not authenticated", "error": err})
}

func BadRequestError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err})
}
