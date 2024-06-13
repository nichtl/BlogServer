package route

import (
	"blogServe/business/global"
	"blogServe/business/middleware/jwt"
	"blogServe/business/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route /**
func Route() *gin.Engine {

	Router := gin.Default()

	baseApiRouter := router.RouterGroupApp

	PublicGroup := Router.Group(global.Config.Serve.Name)
	PrivateGroup := Router.Group(global.Config.Serve.Name)

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"success": true,
				"data":    gin.H{"status": "ok"},
			})
		})
		PublicGroup.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":     "success",
				"success": true,
				"code":    710000,
				"data": gin.H{
					"info":    "success",
					"version": 18,
				},
			})
		})

		baseApiRouter.BaseRouter.InitUserApiRouter(PublicGroup)
		baseApiRouter.BaseRouter.InitCategoryApiRouter(PublicGroup)
		baseApiRouter.BaseRouter.InitArticleApiRouter(PublicGroup)
		baseApiRouter.BaseRouter.InitTagApiRouter(PublicGroup)
	}

	PrivateGroup.Use(jwt.JWT())
	{
		PrivateGroup.GET("/getJwt", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg":     "success",
				"success": true,
				"code":    710000,
				"data":    "1SQMALINSKQ61JNMSkM=",
			})
		})

	}

	return Router
}
