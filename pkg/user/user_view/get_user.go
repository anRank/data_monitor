package user_view

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
}