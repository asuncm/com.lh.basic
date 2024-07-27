package config

import (
	"com.lh.service/tools"
	"fmt"
	"os"
	"strings"
)

type Option = map[string]tools.ServeConf

type Conf struct {
	Root    string `json:"root"`
	Name    string `json:"namespace"`
	DataDir string `json:"data_dir"`
	Lists   Option `json:"lists"`
}

var Config Conf

func InitConfig(key string) {
	Config = Conf{}
	platform := tools.Platform("")
	path := tools.GetPath("LHPATH", "")
	Config.Root = path
	filename := fmt.Sprintf("%s/%s/%s/%s%s", path, key, "config", platform.Env, ".config.yaml")
	config, err := tools.Yaml(filename)
	if err != nil {
		Config.Lists = Option{}
	} else {
		Config.Name = config.Name
		Config.Lists = config.Services
		dir := os.Getenv(config.Database)
		dir = strings.ReplaceAll(dir, "\\", "/")
		Config.DataDir = dir
	}
}

func GetConfig(key string) tools.ServeConf {
	opts := Config.Lists
	return opts[key]
}
