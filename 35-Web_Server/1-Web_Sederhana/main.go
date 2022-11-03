package main

import (
	"fmt"
	"net/http"
)

func main() {
	/**
	Format :
	http.HandleFunc(route, callback)

	Penulisan callback bisa dengan anonymous func yang dituliskan langsung, bisa juga dengan fungsi yang dituliskan terpisah,
	yang pasti, callbacknya harus berbentuk :
	func(w http.ResponseWriter, r *http.Request) {}
	Note :
	w dan r adalah variabel untuk menampung data yang bertipekan http.ResponseWriter dan *http.Request,
	*/

	// membuat route '/' dengan callback yang ditulis langsung
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halo")
	})

	// membuat route '/' dengan callback yang ditulis terpisah
	http.HandleFunc("/hai", hai)

	fmt.Println("menjalankan web server di http://localhost:8080/")

	// menjalankan web server di port 8080
	http.ListenAndServe(":8080", nil)
}

// fungsi callback untuk route '/hai'
func hai(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hai")
}
