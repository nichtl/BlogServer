package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type ArticleApiRouter struct{}

func (s *BaseRouter) InitArticleApiRouter(Router *gin.RouterGroup) (r gin.IRoutes) {
	articleApiRouter := Router.Group("article")
	articleApi := api.ApiGroupApp.ArticleApi
	{
		articleApiRouter.POST("/list", articleApi.QueryAritcleList)
	}
	return articleApiRouter
}
