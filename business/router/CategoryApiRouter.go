package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type CategoryAPIRouter struct{}

func (base *BaseRouter) InitCategoryAPIRouter(c *gin.RouterGroup) (r gin.IRoutes) {
	group := c.Group("category")

	categoryAPI := api.ApiGroupApp.BaseAPI

	{
		group.GET("/detail/:id", categoryAPI.QueryCategoryByID)
		group.GET("/children/:id", categoryAPI.QueryChildrenByID)
		group.GET("/FirstLevelCategory", categoryAPI.QueryAllFirstCategory)
		group.GET("/del/:id", categoryAPI.DeleteCategory)
		group.POST("/add", categoryAPI.CreateCategory)
	}

	return group
}
