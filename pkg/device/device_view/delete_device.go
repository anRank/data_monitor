package device_view

import (
	"data_monitor/common"
	"data_monitor/pkg/app"
	"data_monitor/pkg/device/device_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	innerErrCode, errMsg := device_service.DeleteDevice(int64(id))
	switch innerErrCode {
	case common.INNER_ERR_CODE_UNKNOWN_ERR:
		appG.Response(http.StatusNotFound, app.ERROR, errMsg)
	case common.INNER_ERR_CODE_SUCCESS:
		appG.Response(http.StatusOK, app.SUCCESS, "delete user success")
	default:
		appG.Response(http.StatusBadRequest, app.ERROR, errMsg)
	}
}