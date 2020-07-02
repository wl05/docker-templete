package middleware

import (
	"server/handler/v1"
	"server/pkg/errno"
	"server/pkg/token"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
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
