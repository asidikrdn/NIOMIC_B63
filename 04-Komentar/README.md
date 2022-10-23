# Komentar pada GoLang

Komentar biasa dimanfaatkan untuk menyisipkan catatan pada kode program, menulis penjelasan atau deskripsi mengenai suatu blok kode, atau bisa juga digunakan untuk me-remark kode (men-non-aktifkan kode yg tidak digunakan). Komentar akan diabaikan ketika kompilasi maupun eksekusi program.

Ada 2 jenis komentar di Golang, inline & multiline.

## Komentar Inline

Penulisan komentar jenis ini di awali dengan tanda double slash ( `//` ) lalu diikuti pesan komentarnya. Komentar inline hanya berlaku utuk satu baris pesan saja. Jika pesan komentar lebih dari satu baris, maka tanda `//` harus ditulis lagi di baris selanjutnya.

Contohnya :
`// ini adalah komentar inline`

## Komentar Multiline

Komentar yang cukup panjang akan lebih rapi jika ditulis menggunakan teknik komentar multiline. Ciri dari komentar jenis ini adalah penulisannya diawali dengan tanda `/*` dan diakhiri `*/`.

Contoh :
`/* Ini adalah komentar multi-line */`

Contoh lebih lengkap dari keduanya dapat dilihat pada file `komentar.go`
