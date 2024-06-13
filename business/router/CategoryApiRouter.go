package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type CategoryApiRouter struct{}

func (b *BaseRouter) InitCategoryApiRouter(c *gin.RouterGroup) (R gin.IRoutes) {
	group := c.Group("category")

	categoryApi := api.ApiGroupApp.BaseApi

	{
		group.GET("/detail/:id", categoryApi.QueryCategoryById)
		group.GET("/sub/:id", categoryApi.QueryChildrenById)
		group.GET("/main", categoryApi.QueryAllFirstCategory)
	}

	return group
}
