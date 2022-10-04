# Mini CRUD Mahasiswa

Dalam programming, CRUD merupakan singkatan dari Create Read Update dan Delete. Yakni aplikasi yang berisi pengolahan data. Biasanya CRUD butuh database sebagai media penyimpanan. Akan tetapi untuk sementara ini app Mini CRUD CRUD lebih fokus ke kode React, CRUD ini hanya menggunakan dummy Rest API dengan json-server.

Cara menjalankan app di server lokal :

- Clone repository
- Buka repository tugas_19_react_js di file explorer
- Jalankan terminal pada folder tugas_19_react_js, kemudian ketik perintah `npm install` untuk mendownload dependency yang dibutuhkan dalam app tugas_19_react_js
- Setelah semua dependency terinstall, ketik perintah `npm start` pada terminal untuk menjalankan app tugas_19_react_js di server lokal
- Selamat, app telah berjalan di server lokal pada url `http://localhost:3000`

Langkah selanjutnya adalah menjalankan `json-server` untuk mengaktifkan dummy REST API, caranya sebagai berikut :

- Tetap berada di repository tugas_19_react_js pada file explorer
- Buka terminal baru (terminal yang menjalankan perintah `npm start` jangan ditutup !)
- Pada terminal, ketik perintah `npx json-server --watch karyawan.json --port 3004`
- Selamat, dummy REST API telah berjalan pada url `http://localhost:3004`
