package views

import (
	"data_monitor/pkg/data/data_view"
	"data_monitor/pkg/device/device_view"
	"data_monitor/pkg/user/user_view"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	InitUserRoutes(e)
	InitDeviceRoutes(e)
	InitDataRoutes(e)
}

func InitUserRoutes(e *gin.Engine) {
	user := e.Group("/data_monitor/user")
	{
		user.GET("/:id", user_view.GetUser)
		user.POST("/", user_view.CreateUser)
		user.DELETE("/:id", user_view.DeleteUser)
		user.PUT("/", user_view.UpdateUser)
	}
}

func InitDeviceRoutes(e *gin.Engine) {
	device := e.Group("/data_monitor/device")
	{
		device.GET("/:id", device_view.GetDevice)
		device.POST("/", device_view.CreateDevice)
		device.PUT("/", device_view.UpdateDevice)
		device.DELETE("/:id", device_view.DeleteDevice)
		device.POST("/assign_device_to_user", device_view.AssignDeviceToUser)
		device.POST("/delete_device_to_user", device_view.DeleteDeviceToUser)
	}
}

func InitDataRoutes(e *gin.Engine) {
	data := e.Group("/data_monitor/data")
	{
		data.GET("/", data_view.GetDataList)
		data.GET("/get_data_by_device_id/:id", data_view.GetDataByDeviceId)
	}
}