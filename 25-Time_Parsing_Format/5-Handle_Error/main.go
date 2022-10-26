package main

import (
	"fmt"
	"time"
)

func main() {
	var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(date)
	}
}
