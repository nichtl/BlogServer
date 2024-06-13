package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

type TgTagApiRouter struct{}

func (b *BaseRouter) InitTagApiRouter(c *gin.RouterGroup) (R gin.IRoutes) {
	group := c.Group("tag")

	tagApi := api.ApiGroupApp.BaseApi
	{
		group.POST("/list", tagApi.QueryTagList)
	}

	return group
}
