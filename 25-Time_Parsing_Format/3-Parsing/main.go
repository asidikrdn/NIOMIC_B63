package main

import (
	"fmt"
	"time"
)

func main() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05 Z0700"
	value = "1999-04-14 05:00:59 +0700"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "11/11/2022 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	/**
	Layout format time di golang berbeda dibanding bahasa lain. Umumnya layout format yang
	digunakan adalah seperti "DD/MM/YYYY" dll, di Golang tidak.

	Golang memiliki standar layout format yang cukup unik, contohnya seperti pada kode di atas
	"2006-01-02 15:04:05" . Golang menggunakan 2006 untuk parsing tahun, bukan YYYY;
	01 untuk parsing bulan; 02 untuk parsing hari; dan seterusnya.
	*/

	date, _ = time.Parse(time.RFC822, "14 Mar 99 05:00 WIB")
	fmt.Println(date.String())
	// 1999-03-14 05:00:00 +0000 WIB
}
