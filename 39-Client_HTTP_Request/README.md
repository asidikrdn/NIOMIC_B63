# Simple Client HTTP Request

Pada chapter sebelumnya telah dibahas bagaimana membuat Web Service API yang mem-provide data JSON, pada chapter ini kita akan belajar mengenai cara untuk mengkonsumsi data tersebut.

## Penggunaan HTTP Request

Package `net/http`, selain berisikan tools untuk keperluan pembuatan web, juga berisikan fungsi-fungsi untuk melakukan http request. Salah satunya adalah `http.NewRequest()` yang akan kita bahas di sini.

Sebelumnya, import package yang dibutuhkan. Dan siapkan struct `mahasiswa` yang nantinya akan dipakai sebagai tipe data reponse dari web API. Struk tersebut skema nya sama dengan yang ada pada chapter `38. Web Service API Server`.

Setelah itu buat fungsi `fetchUsers()`. Fungsi ini bertugas melakukan request ke `http://localhost:4135/users`, menerima response dari request tersebut, lalu menampilkannya.

Statement `&http.Client{}` menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.

Fungsi `http.NewRequest()` digunakan untuk membuat request baru. Fungsi tersebut memiliki 3 parameter yang wajib diisi.

- Parameter pertama, berisikan tipe request **POST** atau **GET** atau lainnya
- Parameter kedua, adalah URL tujuan request
- Parameter ketiga, form data request (jika ada)

Fungsi tersebut menghasilkan instance bertipe `http.Request`. Objek tersebut nantinya disisipkan pada saat eksekusi request.

Cara eksekusi request sendiri adalah dengan memanggil method `Do()` pada instance `http.Client` yang sudah dibuat, dengan parameter adalah instance request-nya. Contohnya seperti pada `client.Do(request)`.

Method tersebut mengembalikan instance bertipe `http.Response`, yang di dalamnya berisikan informasi yang dikembalikan dari web API.

Data response bisa diambil lewat property `Body` dalam bentuk string. Gunakan JSON Decoder untuk mengkonversinya menjadi bentuk JSON. Contohnya, `json.NewDecoder(response.Body).Decode(&data)`. Setelah itu barulah kita bisa menampilkannya.

Perlu diketahui, data response perlu di-**close** setelah tidak dipakai. Caranya seperti pada kode `defer response.Body.Close()`.

Selanjutnya, eksekusi fungsi `fetchUsers()` dalam fungsi `main()`.

Setelah itu, jalankan program `Web Service API` dan `Client HTTP Request`.

## HTTP Request dengan Form Data

Untuk menyisipkan data pada sebuah request, ada beberapa hal yang perlu ditambahkan. Yang pertama, import beberapa package lagi, `bytes` dan `net/url`.

Buat fungsi baru, isinya request ke `http://localhost:4135/user` dengan data yang disisipkan adalah `ID`.

Statement `url.Values{}` akan menghasilkan objek yang nantinya digunakan sebagai form data request. Pada objek tersebut perlu di set data apa saja yang ingin dikirimkan menggunakan fungsi `Set()` seperti pada `param.Set("id", ID)`.

Statement `bytes.NewBufferString(param.Encode())` maksudnya, objek form data di-encode lalu diubah menjadi bentuk `bytes.Buffer`, yang nantinya disisipkan pada parameter ketiga pemanggilan fungsi `http.NewRequest()`.

Karena data yang akan dikirim di-encode, maka pada header perlu di set tipe konten request-nya. Kode `request.Header.Set("Content-Type", "application/x-www-form-urlencoded")` artinya tipe konten request di set sebagai `application/x-www-form-urlencoded`. Akan tetapi, kita juga harus meng-set tipe konten request di REST API (Web Service API Server) yang ada di materi sebelumnya menjadi `application/x-www-form-urlencoded` juga.

Jika kita tidak mau merubah tipe konten di REST API (Web Service API Server), maka set saja tipe konten requestnya menjadi `application/form-data`.
