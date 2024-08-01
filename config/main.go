package config

import (
	"com.lh.service/tools"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Option = map[string]tools.ServeConf

type Conf struct {
	Root    string `json:"root"`
	Name    string `json:"namespace"`
	DataDir string `json:"data_dir"`
	Lists   Option `json:"lists"`
}

var options Conf

func InitConfig(key string) {
	options = Conf{}
	platform := tools.Platform("")
	path := tools.GetPath("LHPATH", "")
	options.Root = path
	filename := fmt.Sprintf("%s/%s/%s/%s%s", path, key, "config", platform.Env, ".config.yaml")
	config, err := tools.Yaml(filename)
	if err != nil {
		options.Lists = Option{}
	} else {
		options.Name = config.Name
		options.Lists = config.Services
		dir := os.Getenv(config.Database)
		dir = strings.ReplaceAll(dir, "\\", "/")
		options.DataDir = dir
	}
}

func GetServe(key string) tools.ServeConf {
	opts := options.Lists
	return opts[key]
}

func GetKey(key string) interface{} {
	vals := reflect.ValueOf(options)
	return vals.FieldByName(key).Interface()
}
