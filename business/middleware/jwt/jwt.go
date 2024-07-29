package jwt

import (
	"blogServe/business/global"
	"blogServe/business/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = global.SuccessCode
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			code = global.InvalidParams
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = global.InvalidToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = global.ExpireToken
			}
			val, err := utils.Get(global.TokenPrefix + claims.ID)
			if err != nil {
				code = global.ErrorCode
			}
			if val != token {
				code = global.InvalidToken
			}
			c.AddParam("token", token)
		}
		if code != global.SuccessCode {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "Token鉴权不通过",
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
