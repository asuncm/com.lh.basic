package config

import (
	"com.lh.service/tools"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var configuration tools.MapConf

type Locale = []string

func onRange(data tools.MapConf) {
	for key, value := range data {
		switch value.(type) {
		case string:
			configuration[key] = value.(string)
		default:
			configuration[key] = value.(tools.MapConf)
		}
	}
	return
}

func getConfig(data tools.MapConf, locales Locale) {
	for key, value := range data {
		configuration[key] = value
		root := configuration["root"]
		filepath := fmt.Sprintf("%s/%s/%s", root, key, "locales")
		for _, filename := range locales {
			str, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", filepath, filename))
			keys := configuration[key].(tools.MapConf)
			keys["locales"] = tools.MapConf{}
			vals := keys["locales"].(tools.MapConf)
			if err != nil {
				vals[filename] = tools.MapConf{}
			} else {
				locale := tools.MapConf{}
				err = yaml.Unmarshal(str, &locale)
				if err != nil {
					vals[filename] = tools.MapConf{}
				} else {
					vals[filename] = locale
				}
			}
		}
	}
}

func InitConfig(key string) tools.MapConf {
	locales := []string{"zh_Hans", "zh_Hant", "en_US"}
	configuration = tools.MapConf{}
	platform := tools.Platform("")
	path := tools.GetPath("LHPATH", "")
	configuration["root"] = path
	filename := fmt.Sprintf("%s/%s/%s/%s%s", path, key, "config", platform.Env, ".config.yaml")
	fmt.Println(filename, "oooooooppppppppppppppppp====s", path)
	config, err := tools.Yaml(filename)
	if err != nil {
		return tools.MapConf{}
	}
	configuration["name"] = config["name"]
	serve := config["services"].(tools.MapConf)
	database := config["database"].(tools.MapConf)
	onRange(database)
	dir := os.Getenv(database["root"].(string))
	dir = strings.ReplaceAll(dir, "\\", "/")
	configuration["database"] = dir
	getConfig(serve, locales)
	options := configuration[key].(tools.MapConf)
	fmt.Println(configuration, "config")
	return tools.MapConf{
		"host": options["host"],
		"port": options["port"],
	}
}

func GetServe(key string) tools.MapConf {
	opts := configuration[key]
	return opts.(tools.MapConf)
}

func GetLocale(key string, arrs Locale) interface{} {
	serve := configuration[key].(tools.MapConf)
	locales := serve["locales"]
	return locales
}
