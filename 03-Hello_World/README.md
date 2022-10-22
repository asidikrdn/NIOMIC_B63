# Hello World

Membuat program sederhana untuk menampilkan "Hello World" menggunakan GoLang

Untuk membuat sebuah kode golang, kita harus menggunakan package main dan mengimport fmt.
Penggunaan package main, sama seperti saat kita sedang membuat kode dengan bahasa C atau Java.

Setelah kode dibuat, kita bisa meng-compilenya dengan perintah :
`go build [nama_file.go]`

Setelah dicompile, kita bisa menjalankan file binarynya dengan perintah :
`./[nama_file_binary]`

Pada saat proses development, bisa saja kita langsung menjalankan kode tanpa meng-compilenya terlebih dahulu.
Akan tetepi cara ini hanya disarankan untuk digunakan pada proses development, untuk proses production tetap menggunakan file binary hasil compilenya.
Perintah untuk menjalankan kode tanpa meng-compilenya adalah seperti berikut :
`go run [nama_file.go]`
