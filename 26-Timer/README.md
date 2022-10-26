# Timer

Ada beberapa fungsi dalam package `time` yang memiliki kegunaan sebagai timer. Dengan memanfaatkan fungsi-fungsi tersebut, kita bisa menunda eksekusi sebuah proses dengan durasi waktu tertentu.

## Fungsi `time.Sleep()`

Fungsi ini digunakan untuk menghentikan program sejenak. `time.Sleep()` bersifat blocking, sehingga statement dibawahnya tidak akan dieksekusi sampai waktu pemberhentian usai. Contoh sederhana penerapannya bisa dilihat pada kode berikut.

```Go
fmt.Println("start")
time.Sleep(time.Second * 4)
fmt.Println("after 4 seconds")
```

Tulisan `"start"` muncul, lalu 4 detik kemudian tulisan `"after 4 seconds"` muncul.

## Fungsi `time.NewTimer()`

Fungsi ini sedikit berbeda dengan `time.Sleep()`. Fungsi `time.NewTimer()` mengembalikan sebuah objek `*time.Timer` yang memiliki method `C`. Method ini mengembalikan sebuah channel dan akan dieksekusi dalam waktu yang sudah ditentukan. Contoh penerapannya bisa dilihat pada kode berikut.

```Go
var timer = time.NewTimer(4 * time.Second)
fmt.Println("start")
<-timer.C
fmt.Println("finish")
```

Tulisan `"finish"` akan muncul setelah delay 4 detik.

## Fungsi `time.AfterFunc()`

Fungsi `time.AfterFunc()` memiliki 2 parameter. Parameter pertama adalah durasi timer, dan parameter kedua adalah callback nya. Callback tersebut akan dieksekusi jika waktu sudah memenuhi durasi timer.
Contoh :

```Go
var ch = make(chan bool)
time.AfterFunc(4*time.Second, func() {
fmt.Println("expired")
ch <- true
})
fmt.Println("start")
<-ch
fmt.Println("finish")
```

Tulisan `"start"` akan muncul di awal. Diikuti 4 detik kemudian tulisan `"expired"`. Didalam callback terdapat proses transfer data lewat channel, mengakibatkan tulisan `"finish"` akan muncul tepat setelah tulisan `"expired"` muncul.

Beberapa hal yang perlu diketahui dalam menggunakan fungsi ini:

- Jika tidak ada serah terima data lewat channel, maka eksekusi `time.AfterFunc()` adalah asynchronous dan tidak blocking.
- Jika ada serah terima data lewat channel, maka fungsi akan tetap berjalan asynchronous dan tidak blocking hingga baris kode dimana penerimaan data channel dilakukan.

## Fungsi `time.After()`

Kegunaan fungsi ini mirip seperti `time.Sleep()`. Perbedaannya adalah, fungsi `timer.After()` akan mengembalikan data channel, sehingga perlu menggunakan tanda `<-` dalam penerapannya.

```Go
<-time.After(4 * time.Second)
fmt.Println("expired")
```

Tulisan `"expired"` akan muncul setelah 4 detik.
