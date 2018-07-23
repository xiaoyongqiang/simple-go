package utils

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)

func RandStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateRandImgName(path string, file *multipart.FileHeader) (string, error) {
	if !checkFileIsExist(path) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return "", err
		}
	}

	tail := strings.Split(file.Filename, ".")
	imgname := RandStr(16)
	imgstr := fmt.Sprintf("%s.%s", imgname, tail[len(tail)-1])

	for checkFileIsExist(path + imgstr) {
		imgname = RandStr(16)
		imgstr = fmt.Sprintf("%s.%s", imgname, tail)
	}

	return path + imgstr, nil
}
