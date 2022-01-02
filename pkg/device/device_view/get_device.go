package device_view

import (
	"data_monitor/pkg/app"
	"data_monitor/pkg/device/device_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))


	device, err := device_service.GetDevice(int64(id))
	if err != nil{
		appG.Response(http.StatusNotFound, app.ERROR, nil)
		return
	}

	if device.Name == ""{
		appG.Response(http.StatusNotFound, app.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, app.SUCCESS, map[string]interface{}{
		"device": device,
	})
}
