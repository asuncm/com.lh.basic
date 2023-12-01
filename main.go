package main

import (
	"com.lh.service/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Config() (tools.MiddleConf, error) {
	platform := tools.Platform("")
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	paths := []string{dir, "config", platform.Env + ".config.yaml"}
	pathname := strings.Join(paths, "/")
	configs, err := tools.Yaml(pathname)
	if err != nil {
		return tools.MiddleConf{}, err
	}
	devServe := configs.Services["basic"]
	root := configs.Root[platform.Platform]
	database := configs.Database["badger"]
	return tools.MiddleConf{
		Platform:  platform.Platform,
		Serve:     "basic",
		Root:      root,
		Host:      devServe.Host,
		Port:      devServe.Port,
		DataCache: database[platform.Platform],
		DataPort:  devServe.DataPort,
	}, err
}
func main() {
	router := gin.Default()
	configs, _ := Config()
	router.Use(tools.MiddleWare(configs))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	address := []string{configs.Host, configs.Port}
	router.Run(strings.Join(address, ":"))
}
