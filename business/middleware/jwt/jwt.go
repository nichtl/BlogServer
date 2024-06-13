package jwt

import (
	"blogServe/business/global"
	"blogServe/business/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = global.SUCCESS_CODE
		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("token")
		}
		if token == "" {
			code = global.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = global.INVALID_TOKEN
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = global.EXPIRE_TOKEN
			}
			val, err := global.RedisClient.Get(context.Background(), global.TOKEN_PREFIX+claims.Account).Result()

			if err != nil {
				code = global.ERROR_CODE
			}
			if val != token {
				code = global.INVALID_TOKEN
			}
		}
		if code != global.SUCCESS_CODE {
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
