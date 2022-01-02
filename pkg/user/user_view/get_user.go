package user_view

import (
	"data_monitor/pkg/app"
	"data_monitor/pkg/user/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))


	user, err := user_service.GetUserById(int64(id))
	if err != nil{
		appG.Response(http.StatusNotFound, app.ERROR, nil)
		return
	}

	if user.Name == ""{
		appG.Response(http.StatusNotFound, app.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, app.SUCCESS, map[string]interface{}{
		"user": user,
	})
}