package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlStr := "https://asidikrdn.netlify.app/contact?name=sidik&age=23"

	urlObj, err := url.Parse(urlStr)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("URL : ", urlStr)
	fmt.Println()

	fmt.Printf("Protocol : %s\n", urlObj.Scheme) // https
	fmt.Printf("Host     : %s\n", urlObj.Host)   // asidikrdn.netlify.app
	fmt.Printf("Path     : %s\n", urlObj.Path)   // /contact

	fmt.Printf("Raw Query    : %s\n", urlObj.RawQuery) // name=sidik&age=23
	fmt.Printf("Query    : %s\n", urlObj.Query())      // map[age:[23] name:[sidik]]

	nama := urlObj.Query()["name"][0]                // sidik
	umur := urlObj.Query()["age"][0]                 // 23
	fmt.Printf("Nama : %s, Umur : %s\n", nama, umur) // Nama : sidik, Umur : 23
}
