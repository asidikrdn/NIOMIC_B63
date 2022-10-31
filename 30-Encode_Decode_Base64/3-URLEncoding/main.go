package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	urlStr := "https://asidikrdn.netlify.app"

	encodedUrl := base64.URLEncoding.EncodeToString([]byte(urlStr))
	fmt.Println(encodedUrl)

	decodedUrlByte, _ := base64.URLEncoding.DecodeString(encodedUrl)
	decodedUrl := string(decodedUrlByte)
	fmt.Println(decodedUrl)
}
