package views

import (
	"data_monitor/pkg/user/user_view"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {

	InitUserRoutes(e)
}

func InitUserRoutes(e *gin.Engine) {
	//
	users := e.Group("/data_monitor/user")
	{
		users.GET("/user/:id", user_view.GetUser)
	}
}

