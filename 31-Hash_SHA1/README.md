# Hash SHA1

Hash adalah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak. Jumlah karakter hasil hash selalu sama. Hash termasuk `one-way encryption`, membuat hasil dari hash tidak bisa dikembalikan ke text asli.

`SHA1` atau `Secure Hash Algorithm 1` merupakan salah satu algoritma hashing yang sering digunakan untuk enkripsi data. Hasil dari sha1 adalah data dengan lebar `20 byte` atau `160 bit`, biasa ditampilkan dalam bentuk bilangan heksadesimal 40 digit.

## Penerapan Hash SHA1

Go menyediakan package crypto/sha1, berisikan library untuk keperluan hashing.

## Metode Salting Pada Hash SHA1

Salt dalam konteks kriptografi adalah data acak yang digabungkan pada data asli sebelum proses hash dilakukan.

Hash merupakan enkripsi satu arah dengan lebar data yang sudah pasti, sangat mungkin sekali kalau hasil hash untuk beberapa data adalah sama. Di sinilah kegunaan `salt`, teknik ini berguna untuk mencegah serangan menggunakan metode pencocokan data-data yang hasil hash-nya adalah sama (dictionary attack).

Salt yang digunakan adalah hasil dari ekspresi `time.Now().UnixNano()`. Hasilnya akan selalu unik setiap detiknya, karena scope terendah waktu pada fungsi tersebut adalah nano second atau nano detik.

Metode ini sering dipakai untuk enkripsi password user. Salt dan data hasil hash harus disimpan pada database, karena digunakan dalam pencocokan password setiap user melakukan login.
