/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-22 13:39:12
 * @LastEditTime: 2024-05-22 13:42:45
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package strutil

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Uuid() string {
	id, _ := uuid.NewUUID()
	return strings.ReplaceAll(id.String(), "-", "")
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encoding(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
func Base64Decoding(s string) string {
	decodeStr, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(decodeStr)
}

// 随机生成字符串
func GetRandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成纯字符串
func GetRandomPureString(l int) string {
	str := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成数字字符串
func GetRandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
