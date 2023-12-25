package middlewares

import "github.com/gin-gonic/gin"

func RouterProtection() gin.HandlerFunc {
	return func(c *gin.Context) {
		protection := c.GetHeader("Protection")
		c.Set("Protection", "SDF8S7DF98S7DF"+protection)
		c.Next()
	}

}
