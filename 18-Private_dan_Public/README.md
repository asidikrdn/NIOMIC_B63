# Public dan Private dalam GoLang

Pengembangan aplikasi dalam real development pasti membutuhkan banyak sekali file program. Dan tidak mungkin semuanya di set sebagai package main . Dengan pertimbangan tersebut biasanya file-file tersebut dipisah sebagai package baru.

Folder proyek selain berisikan file-file .go juga bisa berisikan folder. Subfolder tersebut nantinya akan menjadi package baru. Di Golang, 1 subfolder adalah 1 package (kecuali package main yang berada langsung didalam folder proyek). Menjadikan file yang ada di dalam suatu folder memiliki nama package berbeda dengan di folder lainnya. Bahkan antara folder dan subfolder juga bisa memiliki nama package yang berbeda.

Fungsi, struct, dan variabel yang dibuat di package lain, jika diakses dari package main caranya tidak seperti biasanya. Perlu adanya penentuan hak akses yang tepat (apakah public atau private) agar kode tidak kacau balau. Package public, berarti bisa diakses dari package berbeda, sedangkan private berarti hanya bisa di akses dari package yang sama.

Penentuan level akses atau modifier sendiri di golang sangat mudah, ditandai dengan character case nama fungsi/struct/variabel yang ingin di akses. Ketika namanya diawali dengan huruf kapital menandakan bahwa modifier-nya public. Dan sebaliknya, jika diawali huruf kecil, berarti private.

Penjelasan lengkapnya ada pada ebook Dasar Pemrograman Golang by Noval Agung...

NOTE :
Apabila project dibuat diluar folder GOPATH dan GOROOT maka jalankan perintah berikut pada terminal :
`go mod init [nama_project]`
Tanpa menjalankan kode diatas, kita tidak akan dapat mengimport local package.
