package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

func (base *BaseRouter) InitUserApiRouter(c *gin.RouterGroup) (R gin.IRoutes) {
	userApiGroup := c.Group("user")
	userApi := api.ApiGroupApp.BaseApi
	{
		userApiGroup.POST("/login", userApi.Login)
		userApiGroup.POST("/register", userApi.RegisterUser)
		userApiGroup.POST("/loginOut", userApi.LogoutByUser)
	}
	return userApiGroup
}
