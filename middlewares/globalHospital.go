package middlewares

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SecretMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		hospital := c.GetHeader("Hospital")

		if hospital == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Hospital was not provided"})
			c.Abort()
			return
		}

		converted := strconv.Itoa(rand.Intn(100)) + hospital
		c.Set("Hospital", converted)
		c.Next()
	}
}
