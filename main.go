package main

import (
	"com.lh.basic/config"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	opts := config.InitConfig("com.lh.basic")
	yugabyte.InitConfig()
	router.Use(tools.Cors())
	//router.Use(tools.MiddleWare(configs))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	router.Run(fmt.Sprintf("%s:%s", opts["host"], opts["port"]))
}
