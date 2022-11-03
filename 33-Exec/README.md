# Exec

**Exec** digunakan untuk eksekusi perintah command line lewat kode program. Command yang bisa dieksekusi adalah semua command yang bisa dieksekusi di terminal (atau CMD untuk pengguna Windows).

## Penggunaan Exec

Go menyediakan package `exec` berisikan banyak fungsi untuk keperluan eksekusi perintah CLI.

Cara untuk eksekusi command cukup mudah, yaitu dengan menuliskan command dalam bentuk string, diikuti arguments-nya (jika ada) sebagai parameter variadic pada fungsi `exec.Command()`.

Fungsi `exec.Command()` digunakan untuk menjalankan command. Fungsi tersebut bisa langsung di-chain dengan method `Output()`, jika ingin mendapatkan outputnya. Output yang dihasilkan berbentuk `[]byte`, gunakan cast ke string untuk mengambil bentuk string-nya.

## Rekomendasi Penggunaan Exec

Kadang kala, pada saat eksekusi command yang sudah jelas-jelas ada (seperti `ls`, `dir`, atau lainnya) kita menemui error yang mengatakan command not found. Hal itu terjadi karena executable dari command-command tersebut tidak ada. Seperti di windows tidak ada `dir.exe` dan lainnya. Di OS non-windows-pun juga demikian.

Untuk mengatasi masalah ini, tambahkan `bash -c` pada linux/nix command atau `cmd /C` untuk windows. Untuk memeriksa OS yang digunakan user, bisa menggunakan fungsi `runtime.GOOS`. Kemudian hasilnya dijadikan premis pada if-else atau switch-case

## Method Lainnya

Selain `.Output()` ada sangat banyak sekali API untuk keperluan komunikasi dengan OS/CLI yang bisa dipergunakan. Detailnya silakan langsung merujuk ke dokumentasi <https://golang.org/pkg/os/exec/>
