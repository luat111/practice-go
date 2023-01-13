package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
	c.Next()
}
