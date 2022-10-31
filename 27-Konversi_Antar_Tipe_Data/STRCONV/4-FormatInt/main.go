package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Kebalikan dari ParseInt, berguna untuk konversi numerik int64 ke string dengan basis numerik tertentu.

	bilInt := int64(24)
	fmt.Println(bilInt) // 24

	// Konversi bilangan int dengan nilai 24 menjadi string berbentuk number basis 8 (bilangan oktal)
	// varHasil := strconv.FormatInt(intYangInginDiKonversi, basisDataYangDiinginkan)
	bilString := strconv.FormatInt(bilInt, 8)
	fmt.Println(bilString) // 30

}
