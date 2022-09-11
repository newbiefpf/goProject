package middleware

import (
	"github.com/gin-gonic/gin"
	"goProject/utility"
	"goProject/utility/ginResult"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		code := false
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = true
		} else {
			// 解析token
			claims, err := utility.ParseToken(token)
			if err != nil {
				code = true
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = true
			}
		}
		if code {
			c.JSON(200, ginResult.ErrSignParam)
			c.Abort()
			return
		}
		c.Next()
	}
}
