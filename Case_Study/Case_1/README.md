# Case 1 - Membuat simpel REST API dengan fitur CRUD

REST API ini hanya menggunakan fitur bawaan Go tanpa menggunakan `third-party library`.

Data yang digunakan pada REST API ini masih berbentuk local storage yang disimpan langsung dalam kode program dengan berbentuk `slice` dan `struct`.

Berikut list fitur yang dapat digunakan :

- `/users` `GET` : Mengambil seluruh data mahasiswa
- `/users` `POST` : Menambahkan data mahasiswa baru, isikan data baru (hanya data barunya saja) ke dalam `request body`
- `/user?id=xxx` `DELETE` : Menghapus data mahasiswa dengan user id `xxx`
- `/user?id=xxx` `PATCH` : Mengedit data mahasiswa, , isikan data baru untuk mahasiswa yang ingin diedit ke dalam `request body`
