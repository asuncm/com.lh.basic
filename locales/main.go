package locales

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
	"strings"
)

type Option map[string]interface{}
type Locale struct {
	zh_Hant Option
	zh_Hans Option
	en_US   Option
}

var configs Locale

func Init() {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "\\", "/", -1)
	dir = strings.Replace(dir, "/com.lh.auth", "", -1)
	dir = strings.Join([]string{dir, "com.lh.auth", "locales"}, "/")
	Hans, err := os.ReadFile(dir + "/zh_Hans.yaml")
	if err != nil {
		configs.zh_Hans = nil

	} else {
		err = yaml.Unmarshal(Hans, &configs.zh_Hans)
		if err != nil {
			configs.zh_Hans = nil
		}
	}
	Hant, err := os.ReadFile(dir + "/zh_Hant.yaml")
	if err != nil {
		configs.zh_Hant = nil

	} else {
		err = yaml.Unmarshal(Hant, &configs.zh_Hant)
		if err != nil {
			configs.zh_Hant = nil
		}
	}
	ens, err := os.ReadFile(dir + "/en_US.yaml")
	if err != nil {
		configs.en_US = nil

	} else {
		err = yaml.Unmarshal(ens, &configs.en_US)
		if err != nil {
			configs.en_US = nil
		}
	}
}

func Get(key string) Option {
	if key == "zh_Hant" {
		return configs.zh_Hant
	} else if key == "en_US" {
		return configs.en_US
	} else {
		return configs.zh_Hans
	}
}

func GetLocele(c *gin.Context) string {
	locale := c.Request.Header.Get("Accept-Language")
	hans := regexp.MustCompile("^zh-(Hans|CN|CHS)")
	hant := regexp.MustCompile("^zh-(Hant|HK|TW|MO|SG|CHT)")
	if hans.MatchString(locale) {
		return "zh_Hans"
	} else if hant.MatchString(locale) {
		return "zh_Hant"
	} else {
		return "en_US"
	}
}

func getKeys(options Option, arrs []string) interface{} {
	list := []interface{}{options}
OuterLoop:
	for _, value := range arrs {
		item := list[len(list)-1]
		maps, ok := item.(Option)
		if ok {
			vals := maps[value]
			list = append(list, vals)
		} else {
			list = append(list, nil)
			break OuterLoop
		}
	}
	return list[len(list)-1]
}

func GetKey(c *gin.Context, arrs []string) string {
	key := GetLocele(c)
	locale := Get(key)
	var val string
	val = fmt.Sprint(getKeys(locale, arrs))
	return val
}
