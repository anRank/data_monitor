package user_view

import (
	"data_monitor/common"
	"data_monitor/pkg/app"
	"data_monitor/pkg/user/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	request := &DeleteUserRequest{}

	if err := c.BindJSON(request); err!=nil{
		appG.Response(http.StatusBadRequest, app.ERROR, nil)
		return
	}

	innerErrCode, errMsg := user_service.DeleteUser(*request.Id)
	switch innerErrCode {
	case common.INNER_ERR_CODE_UNKNOWN_ERR:
		appG.Response(http.StatusNotFound, app.ERROR, errMsg)
	case common.INNER_ERR_CODE_SUCCESS:
		appG.Response(http.StatusOK, app.SUCCESS, "delete user success")
	default:
		appG.Response(http.StatusBadRequest, app.ERROR, errMsg)
	}
}