package route

import (
	"blogServe/business/global"
	"blogServe/business/middleware/cros"
	"blogServe/business/middleware/jwt"
	"blogServe/business/middleware/log"
	"blogServe/business/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Route /**
func Route() *gin.Engine {

	Router := gin.New()
	Router.Use(log.ZapLogger(), gin.Recovery(), cros.Cors())

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
		PublicGroup.GET("/getJwt", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg":     "success",
				"success": true,
				"code":    710000,
				"data":    "1SQMALINSKQ61JNMSkM=",
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

	}

	PrivateGroup.Use(jwt.JWT())
	{
		baseApiRouter.BaseRouter.InitUserAPIPublicRouter(PrivateGroup)
		baseApiRouter.BaseRouter.InitCategoryAPIRouter(PrivateGroup)
		baseApiRouter.BaseRouter.InitArticleApiRouter(PrivateGroup)
		baseApiRouter.BaseRouter.InitTagAPIRouter(PrivateGroup)

	}

	return Router
}
