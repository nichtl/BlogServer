package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type TgTagAPIRouter struct{}

func (base *BaseRouter) InitTagAPIRouter(c *gin.RouterGroup) (r gin.IRoutes) {
	group := c.Group("tag")

	tagAPI := api.ApiGroupApp.BaseAPI
	{
		group.POST("/list", tagAPI.QueryTagList)
		group.POST("/add", tagAPI.CreateTag)
		group.GET("/del/:id", tagAPI.DeleteTag)
	}

	return group
}
