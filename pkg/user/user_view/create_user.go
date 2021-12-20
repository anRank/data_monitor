package user_view

import (
	"data_monitor/common"
	"data_monitor/pkg/app"
	"data_monitor/pkg/user/user_model"
	"data_monitor/pkg/user/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	request := &CreateUserRequest{}

	if err := c.BindJSON(request); err!=nil{
		appG.Response(http.StatusBadRequest, app.ERROR, nil)
		return
	}

	if err := request.CheckParams(); err != nil {
		appG.Response(http.StatusBadRequest, app.ERROR, err)
		return
	}

	m := &user_model.User{}
	request.CopyToModel(m)
	innerErrCode, errMsg := user_service.CreateUser(m)
	switch innerErrCode {
	case common.INNER_ERR_CODE_UNKNOWN_ERR:
		appG.Response(http.StatusNotFound, app.ERROR, errMsg)
	case common.INNER_ERR_CODE_SUCCESS:
		appG.Response(http.StatusOK, app.SUCCESS, "create user success")
	default:
		appG.Response(http.StatusBadRequest, app.ERROR, errMsg)
	}
}