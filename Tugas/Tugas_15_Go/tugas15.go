package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	// Cara 1, tanpa template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := 1; i <= 100; i++ {
			fmt.Fprintln(w, i)
		}
	})

	var Bilangan []int
	for i := 1; i <= 100; i++ {
		Bilangan = append(Bilangan, i)
	}

	// Cara 2, dengan template
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {

		var indexHTML, err = template.ParseFiles("index.html")
		if err != nil {
			panic(err.Error())
		}

		data := map[string][]int{
			"Bilangan": Bilangan,
		}
		// fmt.Println(Bilangan)
		indexHTML.Execute(w, data)

	})

	fmt.Println("Web server running on http://localhost:4135/")
	http.ListenAndServe(":4135", nil)
}
