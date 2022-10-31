package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains
	fmt.Println("FUNGSI Contains \n------------------")
	strIsExist := strings.Contains("Ahmad Sidik Rudini", "Sid")
	fmt.Println(strIsExist) // true
	strIsExist = strings.Contains("Ahmad Sidik Rudini", "sid")
	fmt.Println(strIsExist) // false
	strIsExist = strings.Contains("Ahmad Sidik Rudini", "di")
	fmt.Println(strIsExist) // true

	fmt.Println()

	// HasPrefix
	fmt.Println("FUNGSI HasPrefix \n------------------")
	strHasPrefix := strings.HasPrefix("Ahmad Sidik Rudini", "Ah")
	fmt.Println(strHasPrefix) // true
	strHasPrefix = strings.HasPrefix("Ahmad Sidik Rudini", "ah")
	fmt.Println(strHasPrefix) // false, HasPrefix bersifat case sensitive
	strHasPrefix = strings.HasPrefix("Ahmad Sidik Rudini", "sid")
	fmt.Println(strHasPrefix) // false

	fmt.Println()

	// HasSuffix
	fmt.Println("FUNGSI HasSuffix \n------------------")
	strHasSuffix := strings.HasSuffix("Ahmad Sidik Rudini", "ni")
	fmt.Println(strHasSuffix) // true
	strHasSuffix = strings.HasSuffix("Ahmad Sidik Rudini", "in")
	fmt.Println(strHasSuffix) // false

	fmt.Println()

	// Count
	fmt.Println("FUNGSI Count \n------------------")
	strCount := strings.Count("Ahmad Sidik Rudini", "i")
	fmt.Println(strCount) // 4
	strCount = strings.Count("Ahmad Sidik Rudini", "d")
	fmt.Println(strCount) // 3

	fmt.Println()

	// Index
	fmt.Println("FUNGSI Index \n------------------")
	strIndex := strings.Index("Ahmad Sidik Rudini", "id")
	fmt.Println(strIndex) // 7
	strIndex = strings.Index("Ahmad Sidik Rudini", "d")
	fmt.Println(strIndex) // 4

	fmt.Println()

	// Replace
	fmt.Println("FUNGSI Replace \n------------------")
	strBeforeReplace := "banana"
	strReplace := strings.Replace(strBeforeReplace, "a", "o", 1)
	fmt.Println(strReplace) // bonana
	strReplace = strings.Replace(strBeforeReplace, "a", "o", 2)
	fmt.Println(strReplace) // bonona
	strReplace = strings.Replace(strBeforeReplace, "a", "o", -1)
	fmt.Println(strReplace) // bonono
	// jika n diisi dengan 0 seperti dibawah, maka akan menyebabkan warning
	// strReplace = strings.Replace(strBeforeReplace, "a", "o", 0)
	// fmt.Println(strReplace) // banana

	fmt.Println()

	// Repeat
	fmt.Println("FUNGSI Repeat \n------------------")
	strRepeat := strings.Repeat("ha", 4)
	fmt.Println(strRepeat)

	fmt.Println()

	// Split
	fmt.Println("FUNGSI Split \n------------------")
	strSplit := strings.Split("Pisang, Apel, Mangga", ", ")
	fmt.Println(strSplit) // [Pisang Apel Mangga]
	strSplit = strings.Split("The Dark Knight", " ")
	fmt.Println(strSplit) // [The Dark Knight]
	strSplit = strings.Split("BATMAN", "")
	fmt.Println(strSplit) // [B A T M A N]

	fmt.Println()

	// Join
	fmt.Println("FUNGSI Join \n------------------")
	sliceData := []string{"Pisang", "Apel", "Mangga"}
	strJoin := strings.Join(sliceData, "-")
	fmt.Println(strJoin) // Pisang-Apel-Mangga
	strJoin = strings.Join([]string{"I", "Love", "You"}, " ")
	fmt.Println(strJoin) // I Love You

	fmt.Println()

	// ToLower
	fmt.Println("FUNGSI ToLower \n------------------")
	strLower := strings.ToLower("AlAy")
	fmt.Println(strLower)

	fmt.Println()

	// ToUpper
	fmt.Println("FUNGSI ToUpper \n------------------")
	strUpper := strings.ToUpper("cicak")
	fmt.Println(strUpper) // CICAK
}
