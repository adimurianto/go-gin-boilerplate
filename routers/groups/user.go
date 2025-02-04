package routers

import (
	controllers "github.com/adimurianto/go-gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, apiVersion string) {
	var ctrl controllers.UserController
	groupRoutes := route.Group(apiVersion)
	groupRoutes.GET("user/", ctrl.GetData)
	groupRoutes.POST("user/", ctrl.CreateData)
}
