# Fungsi String

Go menyediakan package `strings`, isinya banyak fungsi untuk keperluan pengolahan data string.

## Fungsi `strings.Contains()`

Digunakan untuk mendeteksi string tertentu dalam sebuah string, jika string tersebut ditemukan maka fungsi akan mengembalikan nilai `true`.
Format penulisan fungsi :
`strings.Contains(stringUtama, stringYangDicari)`

## Fungsi `strings.HasPrefix()`

Digunakan untuk mendeteksi apakah sebuah string diawali dengan string tertentu. Fungsi ini juga akan mengembalikan nilai bool.
Format penulisan fungsi :
`strings.HasPrefix(stringUtama, stringPembuka)`

## Fungsi `strings.HasSuffix()`

Digunakan untuk mendeteksi apakah sebuah string diakhiri dengan string tertentu. Fungsi ini juga akan mengembalikan nilai bool.
Format penulisan fungsi :
`strings.HasSuffix(stringUtama, stringPenutup)`

## Fungsi `strings.Count()`

Digunakan untuk menghitung jumlah karakter tertentu dalam sebuah string. Fungsi ini akan mengembalikan jumlah karakternya.
Format penulisan fungsi :
`strings.Count(stringUtama, karakterYangDiinginkan)`

## Fungsi `strings.Index()`

Digunakan untuk mencari posisi indeks string tertentu dalam sebuah string. Jika terdapatbeberapa karakter yang sama, maka yang akan diambil indeksnya adalah karakter yang pertama (yang indeksnya paling mendekati 0).
Format penulisan fungsi :
`strings.Index(stringUtama, stringYangDicari)`

## Fungsi `strings.Replace()`

Digunakan untuk mereplace/mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa ditentukan, apakah hanya 1 string pertama, 2 string, atau seluruhnya.
Format penulisan fungsi :
`strings.Replace(stringUtama, stringYangInginDiganti, stringPengganti, jumlahSubstringYangInginDiganti)`
*Jika ingin mengganti seluruh string yang ingin diganti, maka `jumlahSubstringYangInginDiganti` bisa diisi dengan `-1`

## Fungsi `strings.Repeat()`

Digunakan untuk mengulang string sebanyak jumlah tertentu.
Format penulisan fungsi :
`strings.Repeat(stringYangInginDiRepeat, jumlahRepetisi)`

## Fungsi `strings.Split()`

Digunakan untuk memisah string dengan tanda pemisah tertentu. Hasilnya berupa slice string.
Format penulisan fungsi :
`string.Split(stringUtama, "tandaPemisah")`
*Jika tandaPemisah diisi `""` atau string kosong, maka stringUtama akan dipisahkan per 1 karakter

## Fungsi `strings.Join()`

Digunakan untuk menggabungkan slice string menjadi sebuah string dengan pemisah tertentu.
Format penulisan fungsi :
`strings.Join(sliceString, "tandaPemisah")`

## Fungsi `strings.ToLower()`

Mengubah seluruh huruf dalam sebuah string menjadi huruf kecil.
Format penulisan fungsi :
`strings.ToLower(stringUtama)`

## Fungsi `strings.ToUpper()`

Mengubah seluruh huruf dalam sebuah string menjadi huruf besar.
Format penulisan fungsi :
`strings.ToUpper(stringUtama)`
