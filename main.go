package main

import (
	"com.lh.basic/config"
	"com.lh.basic/locales"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	config.InitConfig("com.lh.basic")
	configs := config.GetConfig("basic")
	yugabyte.InitConfig()
	router.Use(tools.Cors())
	//router.Use(tools.MiddleWare(configs))
	locales.Init()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	router.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
