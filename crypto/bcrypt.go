package crypto

import (
	basic "com.lh.basic/locales"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

// 加密密码
func Encrypt(p string, c *gin.Context) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		val := basic.GetKey(c, []string{"errors", "encrypt"})
		return "", errors.New(val)
	}
	return string(password), err
}

// 对加密校验
func Decrypt(oP string, nP string, c *gin.Context) error {
	err := bcrypt.CompareHashAndPassword([]byte(oP), []byte(nP))
	if err != nil {
		arrs := []string{"errors", "verifyPW"}
		msg := basic.GetKey(c, arrs)
		return errors.New(msg)
	}
	return err
}

// UUID生成解码
func UUID(c *gin.Context) (string, error) {
	v4 := uuid.NewV4()
	id, err := uuid.FromString(v4.String())
	if err != nil {
		arrs := []string{"errors", "uuid"}
		msg := basic.GetKey(c, arrs)
		return "", errors.New(msg)
	}
	return id.String(), err
}

// md5加密
func Md5(value string) string {
	if value == "" {
		now := time.Now()
		value = fmt.Sprintf("%d_%d", now.UnixNano(), now.Nanosecond())
	}
	md := md5.New()
	io.WriteString(md, value)
	id := md.Sum(nil)
	return hex.EncodeToString(id)
}

// base64转码
func EnBase(value string) string {
	input := []byte(value)
	return base64.StdEncoding.EncodeToString(input)
}

// base64解码
func DeBase(value string, c *gin.Context) (string, error) {
	base, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		arrs := []string{"errors", "decode"}
		msg := basic.GetKey(c, arrs)
		return "", errors.New(fmt.Sprintf("base64 %s", msg))
	}
	return string(base), err
}

// hmac加密
func Hmac(value string, key string) string {
	hma := hmac.New(md5.New, []byte(key))
	hma.Write([]byte(value))
	return hex.EncodeToString(hma.Sum([]byte("")))
}
