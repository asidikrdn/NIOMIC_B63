# Web Service API Server

Pada chapter ini kita akan mengkombinasikan pembahasan 2 chapter sebelumnya, yaitu web programming dan JSON, untuk membuat sebuah web service API dengan tipe data reponse berbentuk JSON.

Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML.

## Pembuatan Web API

Pertama siapkan terlebih dahulu struct dan beberapa data sample.

Selanjutnya buat fungsi `users()` untuk handle endpoint `/users`. Di dalam fungsi tersebut ada proses deteksi jenis request lewat property `r.Method()`, untuk mencari tahu apakah jenis request adalah **POST** atau **GET** atau lainnya.

Jika request adalah GET (mengambil data), maka data yang di-encode ke JSON dijadikan sebagai response.

Statement `w.Header().Set("Content-Type", "application/json")` digunakan untuk menentukan tipe response, yaitu sebagai JSON. Sedangkan `r.Write()` digunakan untuk mendaftarkan data sebagai response.

Selebihnya, jika request tidak valid, response di set sebagai error menggunakan fungsi `http.Error()`.

Siapkan juga handler untuk endpoint `/user`. Perbedaan endpoint ini dengan `/users` di atas adalah:

- Endpoint `/users` mengembalikan semua sample data yang ada (array).
- Endpoint `/user` mengembalikan satu buah data saja, diambel dari data sample berdasarkan `ID`-nya. Pada endpoint ini, client harus mengirimkan juga informasi `ID` data yang dicari.

Method `r.FormValue()` digunakan untuk mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah `ID`.

Dengan menggunakan `ID` tersebut dicarilah data yang relevan. Jika ada, maka dikembalikan sebagai response. Jika tidak ada, maka error **400, Bad Request** dikembalikan dengan pesan **User Not Found**.

Terakhir, implementasikan kedua handler di atas.

Jalankan program, sekarang web server sudah live dan bisa dikonsumsi datanya.

## Test Web Service API via Postman

Setelah web server sudah berjalan, web service yang telah dibuat perlu untuk di-tes. Di sini saya menggunakan aplikasi bernama Postman untuk mengetes API yang sudah dibuat.

- Test endpoint `/users`, apakah data yang dikembalikan sudah benar.
- Test endpoint `/user`, isi form data id dengan nilai `U001`. (bisa dengan menambahkan query variabel, atau langsung menuliskannya di url)
