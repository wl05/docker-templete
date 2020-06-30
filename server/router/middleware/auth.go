package middleware

import (
	"github.com/gin-gonic/gin"
)
// AuthMiddleware Parse the json web token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			v1.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
