package main

import (
	"data_monitor/config"
	"data_monitor/views"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置初始化
	config.InitConfig()

	r := gin.Default()

	views.InitRoutes(r)

	if err := r.Run(":8878"); err != nil { // listen and serve on 0.0.0.0:8080
		panic(err)
	}
}