# Goroutine

Goroutine bukanlah thread. Sebuah native thread bisa berisikan sangat banyak goroutine. Mungkin lebih pas kalau goroutine disebut sebagai mini thread. Goroutine sangat ringan, hanya dibutuhkan sekitar 2kB memori saja untuk satu buah goroutine. Ekseksui goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.

## Penerapan Goroutine

Untuk menerapkan goroutine, proses yang akan dieksekusi sebagai goroutine harus dibungkus kedalam fungsi. Pada saat pemanggilan fungsi tersebut, ditambahkan keyword go didepannya. Dengan demikian proses yang ada dalam fungsi tersebut dideteksi sebagai goroutine baru.

Penjelasan lengkapnya ada pada ebook Dasar Pemrograman Golang by Noval Agung...
