package main

import "fmt"

func main() {
	// Mengggunakan if else
	var nilai = 8
	if nilai == 10 {
		fmt.Println("Lulus dengan sempurna")
	} else if nilai >= 7 {
		fmt.Println("Lulus")
	} else if nilai == 6 {
		fmt.Println("Sedikit lagi")
	} else {
		fmt.Println("Belajar lagi")
	}

	// Menggunakan switch case
	nilai = 10
	switch nilai {
	// Kode program dalam case, bisa menggunakan kurung kurawal seperti dibawah
	case 10:
		{
			fmt.Println("Lulus dengan sempurna")
		}
	case 9:
		{
			fmt.Println("Lulus")
		}
	case 8:
		{
			fmt.Println("Lulus")
		}
	case 7:
		{
			fmt.Println("Lulus")
		}
	// Kode program dalam case, bisa juga tidak menggunakan kurung kurawal seperti dibawah
	case 6:
		fmt.Println("Sedikit lagi")
	default:
		fmt.Println("Belajar lagi")
	}
}
