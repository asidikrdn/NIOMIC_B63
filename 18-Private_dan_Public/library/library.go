package library

import "fmt"

// Membuat sebuah fungsi private
// func iniprivate() {
// 	fmt.Println("Saya di private")
// }
func sebutnama(name string) string {
	return name
}

// Membuat sebuah fungsi public
func Inipublic() {
	fmt.Println("saya di public, nama saya", sebutnama("Sidik")) // menyisipkan fungsi private di fungsi public agar bisa dipanggil di file lain
}
