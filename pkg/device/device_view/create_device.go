package device_view

import (
	"data_monitor/common"
	"data_monitor/pkg/app"
	"data_monitor/pkg/device/device_model"
	"data_monitor/pkg/device/device_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	request := &CreateDeviceRequest{}

	if err := c.BindJSON(request); err!=nil{
		appG.Response(http.StatusBadRequest, app.ERROR, nil)
	}

	if err := request.CheckParams(); err!= nil{
		appG.Response(http.StatusBadRequest, app.ERROR, err)
		return
	}

	m := &device_model.Device{}
	request.CopyToModel(m)
	innerErrCode, errMsg := device_service.CreateDevice(m)
	switch innerErrCode {
	case common.INNER_ERR_CODE_UNKNOWN_ERR:
		appG.Response(http.StatusNotFound, app.ERROR, errMsg)
	case common.INNER_ERR_CODE_SUCCESS:
		appG.Response(http.StatusOK, app.SUCCESS, "create device success")
	default:
		appG.Response(http.StatusBadRequest, app.ERROR, errMsg)
	}
}