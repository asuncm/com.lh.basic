package crypto

import (
	"bytes"
	"com.lh.service/src/tools"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

type AesConf struct {
	ID    string `json:"id"`
	Key   string `json:"Key"` // key支持16、24、32为加密数据
	Value string `json:"value"`
}

// PKCS7 填充模式
func cs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func cs7UnPadding(origData []byte, msg string) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New(msg)
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// 实现加密
func AesEnCrypt(data AesConf, pathname string) (string, error) {
	key := []byte(data.Key)
	pw := []byte(data.Value)

	cData, err := aes.NewCipher(key)
	if err != nil {
		config, _ := tools.Language(pathname)
		return "", errors.New(config["aes"]["cliperError"])
	}
	cSize := cData.BlockSize()
	//对数据进行填充，让数据长度满足需求
	pw = cs7Padding(pw, cSize)
	//采用AES加密方法中CBC加密模式
	cMode := cipher.NewCBCEncrypter(cData, key[:cSize])
	lists := make([]byte, len(pw))
	// 实现加密
	cMode.CryptBlocks(lists, pw)
	str := hex.EncodeToString(lists)
	return str, err
}

// 实现解密
func AesDeCrypt(data AesConf, pathname string) (string, error) {
	bytes, err := hex.DecodeString(data.Value)
	config, _ := tools.Language(pathname)
	if err != nil {
		return "", errors.New(config["aes"]["hexError"])
	}
	key := []byte(data.Key)
	cData, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New(config["aes"]["cliperError"])
	}
	//获取块大小
	cSize := cData.BlockSize()
	//创建加密客户端实例
	cMode := cipher.NewCBCDecrypter(cData, key[:cSize])
	lists := make([]byte, len(bytes))
	//这个函数也可以用来解密
	cMode.CryptBlocks(lists, bytes)
	lists, err = cs7UnPadding(lists, config["aes"]["csError"])
	if err != nil {
		return "", err
	}
	return string(lists), err
}
