# Channel

Channel digunakan untuk menghubungkan gorutine satu dengan goroutine lainnya. Mekanismenya adalah dengan cara serah-terima data lewat channel tersebut. Goroutine pengirim dan penerima harus berada pada channel yang berbeda (konsep ini disebut buffered channel). Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.
