package util

import (
	"time"
	"math/rand"
	// "io/ioutil"
	// "encoding/base64"
)

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLXCVZXVNMQERQWRE")

	result := make([]byte, n)

	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

// func FormFileToBase64String(file *multipart.FileHeader) string {
// 	_file, _ := file.Open()
// 	defer _file.Close()

// 	fileData, _ := ioutil.ReadAll(_file)

// 	bs64 := base64.StdEncoding.EncodeToString(fileData)

// 	return bs64
// }