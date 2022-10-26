# Channel

Channel digunakan untuk menghubungkan gorutine satu dengan goroutine lainnya. Mekanismenya adalah dengan cara serah-terima data lewat channel tersebut. Goroutine pengirim dan penerima harus berada pada channel yang berbeda (konsep ini disebut buffered channel). Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

## Buffered Channel

Channel secara default adalah un-buffered, tidak di-buffer di memori. Ketika ada goroutine yang mengirimkan data lewat channel, harus ada goroutine lain yang bertugas menerima data dari channel yang sama, dengan proses serah-terima yang bersifat blocking. Maksudnya, baris kode di bagian pengiriman dan penerimaan data, tidak akan akan diproses sebelum proses serah-terima-nya selesai.

Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan jumlah buffer-nya. Angka tersebut akan menjadi penentu kapan kita bisa mengirimkan data. Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan asynchronous (tidak blocking).

Ketika jumlah data yang dikirim sudah melewati batas buffer, maka pengiriman data hanya
bisa dilakukan ketika salah satu data sudah diambil dari channel, sehingga ada slot channel
yang kosong. Dengan proses penerimaan-nya sendiri bersifat blocking.

Penerapan buffered channel pada dasarnya mirip seperti channel biasa. Perbedannya pada channel jenis ini perlu disiapkan jumlah buffer-nya.
