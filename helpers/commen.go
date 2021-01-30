package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func Md5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	//将str写入到w中
	return fmt.Sprintf("%x", w.Sum(nil))
}
