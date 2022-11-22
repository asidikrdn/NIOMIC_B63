package main

import (
	"fmt"
	"regexp"
)

func main() {
	var email string

	regex, _ := regexp.Compile(`^\w+\d*@\w*\.\w*$`)

	fmt.Println("Masukkan alamat email anda :")
	fmt.Scanln(&email)

	fmt.Println()

	if regex.MatchString(email) {
		fmt.Println("Email Anda Valid")

		regex, _ = regexp.Compile(`@\w*\.\w*`)
		domain := regex.FindString(email)
		fmt.Println("Domain alamat email anda :", domain)
	} else {
		fmt.Println("Email Anda Invalid")
	}

}
