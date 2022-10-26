package main

import (
	"fmt"
	"time"
)

func main() {
	// membuat sebuah object time dengan mem-parse dari string
	date, _ := time.Parse(time.RFC822, "14 Mar 99 05:00 WIB")

	// memformat object time menjadi string dengan menuliskan komponen-komponen format layouting
	dateSatu := date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateSatu :", dateSatu)

	// memformat object time menjadi string dengan predefined format layout
	dateDua := date.Format(time.RFC3339)
	fmt.Println("dateDua :", dateDua)
}
