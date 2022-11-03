# Argument & Flag

**Arguments** adalah data opsional yang disisipkan ketika eksekusi program. Sedangkan **flag** merupakan ekstensi dari argument. Dengan flag, penulisan argument menjadi lebih rapi dan terstruktur.

## Penggunaan Arguments

Data arguments bisa didapat lewat variabel `os.Args` (package `os` perlu di-import terlebih dahulu). Data tersebut tersimpan dalam bentuk array dengan pemisah adalah tanda spasi.

Pada saat eksekusi program disisipkan juga argument-nya. Sebagai contoh disisipkan 3 buah data sebagai argumen, yaitu: `banana`, `potato`, dan `ice cream`.
Untuk eksekusinya sendiri bisa menggunakan `go run` ataupun dengan cara build-execute. Data argumen yang ada karakter spasi nya ( ), maka harus diapit tanda petik ("), agar tidak dideteksi sebagai 2 argumen.

- Menggunakan `go run` : `go run bab45.go banana potato "ice cream"`
- Menggunakan `go build` : `go build bab45.go` & `./bab45 banana potato "ice cream"`

Variabel `os.Args` mengembalikan tak hanya arguments saja, tapi juga path file executable (jika eksekusi-nya menggunakan `go run` maka path akan merujuk ke folder temporary). Gunakan `os.Args[1:]` untuk mengambil slice arguments-nya saja.

## Penggunaan Flag

Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk *parameterize* eksekusi program, dengan penulisan dalam bentuk key-value.

Cara penulisan arguments menggunakan flag:
`go run bab45.go -name="john wick" -age=28`

Tiap argument harus ditentukan key, tipe data, dan nilai default-nya. Agar lebih mudah dipahami, mari kita bahas kode berikut.

```Go
var dataName = flag.String("name", "anonymous", "type your name")
fmt.Println(*dataName)
```

Kode tersebut maksudnya adalah, disiapkan flag bertipe `string`, dengan key adalah `name`, dengan nilai default `"anonymous"`, dan keterangan `"type your name"`. Nilai flag nya sendiri akan disimpan ke dalam variabel dataName.
Nilai balik fungsi `flag.String()` adalah string pointer, jadi perlu di-dereference terlebih dahulu agar bisa mendapatkan nilai aslinya (`*dataName`).
Flag yang nilainya tidak di set, secara otomatis akan mengembalikan nilai default. Misalnya contoh diataas, jika flagnya tidak di-set maka `dataName` akan berisi `"anonymous"`.

Tabel lengkap fungsi flag yang tersedia di GoLang dapat dilihat pada ebook Dasar Pemrograman GoLang by NovalAgung

## Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data

Sebenarnya ada 2 cara deklarasi flag yang bisa digunakan, dan cara di atas merupakan cara pertama.

Cara kedua mirip dengan cara pertama, perbedannya adalah kalau di cara pertama nilai pointer flag dikembalikan lalu ditampung variabel. Sedangkan pada cara kedua, nilainya diambil lewat parameter pointer.

Tinggal tambahkan akhiran `Var` pada pemanggilan nama fungsi flag yang digunakan (contoh `flag.IntVar()`, `flag.BoolVar()`, dll), lalu disisipkan referensi variabel penampung flag sebagai parameter pertama.

Kegunaan dari parameter terakhir method-method flag adalah untuk memunculkan hints atau petunjuk arguments apa saja yang bisa dipakai, ketika argument `--help` ditambahkan saat eksekusi program.
