package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	waktu := 5
	apakahWaktuHabis := make(chan bool)

	// Menjalankan fungsi setTimeout dan isTimeout secara async
	go setTimeout(waktu, apakahWaktuHabis)
	go isTimeout(waktu, apakahWaktuHabis)

	// Menjalankan beberapa blok kode secara synchronous / blocking
	var input string
	fmt.Print("Berapakah hasil dari 25/5 ? ")
	fmt.Scanf("%s", &input)

	if input == "5" {
		fmt.Println("Jawaban anda benar !")
	} else {
		fmt.Println("Jawaban anda salah !")
	}

	/**
	Dengan struktur kode diatas,
	maka saat program dijalankan fungsi setTimeout dan isTimeout akan langsung berjalan secara async.
	Bersamaan dengan hal tersebut, kode selanjutnya juga akan dieksekusi
	dan akan terhenti sementara setelah mencetak teks "Berapakah hasil dari 25/5 ?"

	Dari kondisi tersebut, ada 2 kemungkinan.

	Kemungkinan pertama, jika user tak menjawab dalam kurun waktu 5 detik,
	maka fungsi setTisTimeout akan mengisi channel apakahWaktuHabis dengan nilai true,
	lalu fungsi isTimeout akan menerima data channel apakahWaktuHabis, kemudian menampilkan sebuah teks
	dan menjalankan fungsi Exit yang artinya program akan berakhir.

	Kemungkinan kedua, jika user menjawab pertanyaan kurang dari 5 detik,
	jawaban user akan disimpan pada variabel input.
	Kemudian program akan menghasilkan sebuah teks yang isinya tergantung dari jawaban user tersebut.
	*/
}

// Membuat fungsi untuk mengirim nilai true ke sebuah channel dalam durasi tertentu.
// durasi dan channelnya diinput sebagai argument
func setTimeout(timeout int, cH chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		cH <- true
	})
}

// Membuat fungsi untuk mengamati apakah channel yang dibuat sudah bernilai true,
// Apabila sudah bernilai true, maka tampilkan teks dan akhiri program
func isTimeout(timeout int, cH <-chan bool) {
	<-cH
	fmt.Println("\nTimeout! tidak ada jawaban setelah lebih dari", timeout, "detik")
	os.Exit(0) // menutup program dengan kode '0'
}
