package views

import (
	"data_monitor/pkg/device/device_view"
	"data_monitor/pkg/user/user_view"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	InitUserRoutes(e)
	InitDeviceRoutes(e)
}

func InitUserRoutes(e *gin.Engine) {
	user := e.Group("/data_monitor/user")
	{
		user.GET("/user/:id", user_view.GetUser)
		user.POST("/user", user_view.CreateUser)
		user.POST("/delete_user", user_view.DeleteUser)
		user.POST("/update_user", user_view.UpdateUser)
	}
}

func InitDeviceRoutes(e *gin.Engine) {
	device := e.Group("/data_monitor/device")
	{
		device.GET("/device/:id", device_view.GetDevice)
		device.POST("/device", device_view.CreateDevice)
		device.POST("/assign_device_to_user", device_view.AssignDeviceToUser)
		device.POST("/delete_device_to_user", device_view.DeleteDeviceToUser)
	}
}

