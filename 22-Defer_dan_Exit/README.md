# Defer & Exit

`Defer` digunakan untuk mengakhirkan eksekusi sebuah statement. Sedangkan `Exit` digunakan untuk menghentikan program. 2 topik ini sengaja digabung agar hubungan antara keduanya lebih mudah dipahami.

## Penerapan keyword `defer`

Seperti yang sudah dijelaskan secara singkat di atas, bahwa defer digunakan untuk mengakhirkan eksekusi baris kode. Ketika eksekusi sudah sampai pada akhir blok fungsi, statement yang di defer baru akan dijalankan. Defer bisa ditempatkan di mana saja, awal maupun akhir blok.

Ketika ada banyak statement yang di-defer, maka statement tersebut akan dieksekusi di akhir secara berurutan.

## Penerapan Fungsi os.Exit()

Exit digunakan untuk menghentikan program secara paksa pada saat itu juga. Semua statement setelah exit tidak akan di eksekusi, termasuk juga defer.

Fungsi `os.Exit()` berada dalam package `os`. Fungsi ini memiliki sebuah parameter bertipe numerik yang wajib diisi. Angka yang dimasukkan akan muncul sebagai exit status ketika program berhenti.
