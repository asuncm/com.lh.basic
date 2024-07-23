package main

import (
	"com.lh.basic/locales"
	"com.lh.service/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Config() (tools.MiddleConf, error) {
	platform := tools.Platform("")
	root := tools.GetPath("LHPATH", "")
	pathname := fmt.Sprintf("%s%s%s%s", root, "/config/", platform.Env, ".config.yaml")
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["basic"]
	database := tools.GetPath(configs.Database, "pebble/basic")
	return tools.MiddleConf{
		Platform: platform.Platform,
		Serve:    fmt.Sprintf("%s%s", root, "/com.lh.basic"),
		Root:     root,
		Host:     devServe.Host,
		Port:     devServe.Port,
		DataDir:  database,
	}, err
}
func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(tools.MiddleWare(configs))
	locales.Init()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
