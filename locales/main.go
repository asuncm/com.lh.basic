package locales

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

type Option map[string]interface{}
type Locale = map[string]Option

var configs = Locale{}

func Init(key string) {
	dir := fmt.Sprintf("%s%s", key, "/com.lh.basic/locales")
	arrs := []string{"zh_Hans", "zh_Hant", "en_US"}
	for _, val := range arrs {
		str, err := os.ReadFile(fmt.Sprintf("%s/%s.yaml", dir, val))
		if err != nil {
			configs[val] = Option{}
		} else {
			locale := Option{}
			err = yaml.Unmarshal(str, &locale)
			if err != nil {
				configs[val] = Option{}
			} else {
				configs[val] = locale
			}
		}
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
	key := arrs[0]
	opts := arrs[1:len(arrs)]
	list := []interface{}{options[key]}
	for _, val := range opts {
		item := list[len(list)-1].(Option)
		vals := item[val]
		list = append(list, vals)
	}
	return list[len(list)-1]
}

func GetKey(c *gin.Context, arrs []string) string {
	key := GetLocele(c)
	locale := configs[key]
	var val string
	val = fmt.Sprint(getKeys(locale, arrs))
	return val
}
