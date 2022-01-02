package data_view

import (
	"data_monitor/pkg/app"
	"data_monitor/pkg/data/data_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDataList(c *gin.Context) {
	appG := app.Gin{C: c}

	data, err := data_service.GetDataList()
	if err != nil{
		appG.Response(http.StatusNotFound, app.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, app.SUCCESS, map[string]interface{}{
		"data": data,
	})
}