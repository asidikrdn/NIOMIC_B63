package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	// membuat routing file server (berguna untuk mengakses file-file static/aset yang dibutuhkan, misalnya file css, gambar, js, dll)
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("35-Web_Server/2-Web_Template/assets"))))

	// membuat route '/'
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Nama":  "Ahmad Sidik Rudini",
			"Prodi": "Teknik Informatika",
			"NPM":   "201106041165",
		}

		// mengambil file template
		// t, err := template.ParseFiles("/home/asidikrdn/Documents/Programming/NIOMIC-Golang/35-Web_Server/2-Web_Template/template.html")	// cara 1
		t, err := template.ParseFiles("35-Web_Server/2-Web_Template/template.html") // cara 2
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// mengeksekusi template, dan memparse "data" ke template
		t.Execute(w, data)
	})

	// membuat route '/test'
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Nama":  "Ahmad Sidik Rudini",
			"Prodi": "Teknik Informatika",
			"NPM":   "201106041165",
		}

		// mengambil file template
		// t, err := template.ParseFiles("/home/asidikrdn/Documents/Programming/NIOMIC-Golang/35-Web_Server/2-Web_Template/template.html")	// cara 1
		t, err := template.ParseFiles("35-Web_Server/2-Web_Template/template2.html") // cara 2
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// mengeksekusi template, dan memparse "data" ke template
		t.Execute(w, data)
	})

	fmt.Println("menjalankan web server di http://localhost:8080/")

	// menjalankan web server di port 8080
	http.ListenAndServe(":8080", nil)
}
