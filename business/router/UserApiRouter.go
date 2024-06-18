package router

import (
	"blogServe/business/api"
	"github.com/gin-gonic/gin"
)

func (base *BaseRouter) InitUserAPIPublicRouter(c *gin.RouterGroup) (r gin.IRoutes) {
	userAPIGroup := c.Group("user")
	userAPI := api.ApiGroupApp.BaseAPI
	{
		userAPIGroup.POST("/login", userAPI.Login)
		userAPIGroup.POST("/update", userAPI.UpdateUser)
		userAPIGroup.POST("/register", userAPI.RegisterUser)
		userAPIGroup.POST("/loginOut", userAPI.LogoutByUser)
		userAPIGroup.POST("/logout", userAPI.LogoutByUser)
	}
	return userAPIGroup
}
