package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type ArticleAPIRouter struct{}

func (base *BaseRouter) InitArticleApiRouter(router *gin.RouterGroup) (r gin.IRoutes) {
	articleApiRouter := router.Group("article")
	articleAPI := api.ApiGroupApp.ArticleAPI
	{
		articleApiRouter.POST("/list", articleAPI.QueryAritcleList)
		articleApiRouter.POST("/add", articleAPI.CreateArticle)
		articleApiRouter.Group("/del/:id", articleAPI.DeleteArticle)
		articleApiRouter.POST("/update", articleAPI.UpdateArticle)
	}
	return articleApiRouter
}
