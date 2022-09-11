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
			c.JSON(200, ginResult.OK.WithData("token无效!"))
		} else {
			// 解析token
			claims, err := utility.ParseToken(token)
			if err != nil {
				code = true
				c.JSON(200, ginResult.OK.WithData("token无效!"))
				c.Abort()
				return

			} else if time.Now().Unix() > claims.ExpiresAt {
				code = true
				c.JSON(200, ginResult.OK.WithData("token过期!"))
				c.Abort()
				return
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
