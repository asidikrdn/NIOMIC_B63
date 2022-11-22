# Tugas 9

Cara menjalankan program :

1. Clone source code project ini
2. Buka folder project `Tugas_9_GO`
3. Jalankan program `tugas9.go`

## Hasil analisa

Di dalam program yang saya buat, terdapat 1 buah defer dan error panic. Setelah saya coba menjalankan program tersebut, apabila tidak terjadi error maka kode yang di-defer akan dijalankan pada akhir program. Akan tetapi, jika terjadi error dan fungsi panic dijalankan, maka kode yang di-defer akan dijalankan tepat sebelum pesan error dari panic ditampilkan.
