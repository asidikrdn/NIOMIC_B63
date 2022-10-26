# Channel

Channel digunakan untuk menghubungkan gorutine satu dengan goroutine lainnya. Mekanismenya adalah dengan cara serah-terima data lewat channel tersebut. Goroutine pengirim dan penerima harus berada pada channel yang berbeda (konsep ini disebut buffered channel). Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

## Buffered Channel

Channel secara default adalah un-buffered, tidak di-buffer di memori. Ketika ada goroutine yang mengirimkan data lewat channel, harus ada goroutine lain yang bertugas menerima data dari channel yang sama, dengan proses serah-terima yang bersifat blocking. Maksudnya, baris kode di bagian pengiriman dan penerimaan data, tidak akan akan diproses sebelum proses serah-terima-nya selesai.

Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan jumlah buffer-nya. Angka tersebut akan menjadi penentu kapan kita bisa mengirimkan data. Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan asynchronous (tidak blocking).

Ketika jumlah data yang dikirim sudah melewati batas buffer, maka pengiriman data hanya
bisa dilakukan ketika salah satu data sudah diambil dari channel, sehingga ada slot channel
yang kosong. Dengan proses penerimaan-nya sendiri bersifat blocking.

Penerapan buffered channel pada dasarnya mirip seperti channel biasa. Perbedannya pada channel jenis ini perlu disiapkan jumlah buffer-nya.

## Channel - Select

Adanya channel memang sangat membantu pengontrolan goroutine, jumlah goroutine yang banyak bukan lagi masalah. Ada kalanya dimana kita butuh tak hanya satu channel saja untuk manage goroutine yang
juga banyak, dibutuhkan beberapa atau mungkin banyak channel.

Disinilah kegunaan dari `select`. Select memudahkan pengontrolan komunikasi data lewat channel. Cara penggunaannya sama seperti seleksi kondisi `switch`.

Program pencarian rata-rata dan nilai tertinggi berikut merupakan contoh sederhana penerapan select dalam channel. Akan ada 2 buah goroutine yang masing-masing di-handle oleh sebuah channel. Setiap kali goroutine selesai dieksekusi, akan dikirimkan datanya ke channel yang bersangkutan. Lalu dengan menggunakan select, akan dikontrol penerimaan datanya.

Pertama, kita siapkan terlebih dahulu 2 fungsi yang akan dieksekusi sebagai goroutine baru. Fungsi pertama digunakan untuk mencari rata-rata, dan fungsi kedua untuk penentuan nilai tertinggi dari sebuah slice.

## Channel - Range dan Close

Penerimaan data lewat channel yang dipakai oleh banyak goroutine, akan lebih mudah dengan memanfaatkan keyword `for - range`.

`for - range` jika diterapkan pada channel, akan melakukan perulangan tanpa henti. Perulangan tersebut tetap berjalan meski tidak ada transaksi pada channel, dan hanya akan berhenti jika status channel berubah menjadi closed atau sudah ditutup. Fungsi `close` digunakan utuk menutup channel.

Channel yang sudah ditutup tidak bisa digunakan lagi untuk menerima maupun mengirim data. Menjadikan penerimaan data menggunakan `for - range` juga ikut berhenti.

### Channel Direction

Ada yang unik dengan fitur parameter channel yang disediakan Golang. Level akses channel bisa ditentukan, apakah hanya sebagai penerima, pengirim, atau penerima sekaligus pengirim. Konsep ini disebut dengan `channel direction`. Cara pemberian level akses adalah dengan menambahkan tanda `<-` sebelum atau setelah keyword `chan`. Untuk lebih jelasnya bisa dilihat di list berikut.

- `cH chan string`    : Parameter `cH` bisa digunakan untuk mengirim dan menerima data
- `cH chan<- string`  : Parameter `cH` hanya bisa digunakan untuk mengirim data
- `cH <-chan string`  : Parameter `cH` hanya bisa digunakan untuk menerima data

Sebagai contoh fungsi `sendMessage(cH chan<- string)` yang parameter `cH` dideklarasikan dengan level akses untuk pengiriman data saja. Channel tersebut hanya bisa digunakan untuk mengirim, contohnya: `cH <- fmt.Sprintf("data %d", i)`.

Dan sebaliknya pada fungsi `printMessage(cH <-chan string)`, channel `cH` hanya bisa digunakan untuk menerima data saja.
