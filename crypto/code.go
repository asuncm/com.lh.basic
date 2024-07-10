package crypto

import "github.com/gin-gonic/gin"

type Option struct {
	ID  string `json:"id"`
	MD5 string `json:"md5"`
	Key string `json:"key"`
}

func ListId(value int64, c *gin.Context) (Option, error) {
	node, err := NewWorker(value)
	id := node.GetId()
	key, _ := UUID(c)
	md5 := Md5(key)
	config := Option{
		ID:  id,
		MD5: md5,
		Key: key,
	}
	return config, err.(error)
}
