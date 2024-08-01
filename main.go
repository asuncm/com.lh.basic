package main

import (
	auth "com.lh.auth/locales"
	"com.lh.basic/config"
	basic "com.lh.basic/locales"
	serve "com.lh.service/locales"
	"com.lh.service/tools"
	"com.lh.service/yugabyte"
	user "com.lh.user/locales"
	web "com.lh.web.service/locales"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	config.InitConfig("com.lh.basic")
	configs := config.GetServe("basic")
	root := config.GetKey("Root").(string)
	yugabyte.InitConfig()
	router.Use(tools.Cors())
	//router.Use(tools.MiddleWare(configs))
	auth.Init(root)
	basic.Init(root)
	user.Init(root)
	web.Init(root)
	serve.Init(root)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	router.Run(fmt.Sprintf("%s:%s", configs.Host, configs.Port))
}
