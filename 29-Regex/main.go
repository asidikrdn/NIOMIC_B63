package main

import (
	"fmt"
	"regexp"
)

func main() {
	// MatchString
	text1 := "banana burger soup"
	regex1, _ := regexp.Compile(`[a-z]+`)
	isMatch := regex1.MatchString(text1)
	fmt.Println(isMatch) // true

	fmt.Println()

	// FindString
	text2 := "banana burger soup"
	regex2, _ := regexp.Compile(`[a-z]+`)
	findStr := regex2.FindString(text2)
	fmt.Println(findStr) // banana

	fmt.Println()

	// FindStringIndex
	text3 := "banana burger soup"
	regex3, _ := regexp.Compile(`[a-z]+`)
	findIndexStr := regex3.FindStringIndex(text3)
	fmt.Println(findIndexStr) // [0 6] string ke 0 sampai 6 (6nya tidak termasuk)

	fmt.Println()

	// FindAllString
	text4 := "banana burger soup"
	regex4, _ := regexp.Compile(`[a-z]+`)
	findAllStr := regex4.FindAllString(text4, -1)
	fmt.Println(findAllStr) // [banana burger soup]
	findAllStr = regex4.FindAllString(text4, 2)
	fmt.Println(findAllStr) // [banana burger]

	fmt.Println()

	// ReplaceAllString
	text5 := "banana burger soup"
	regex5, _ := regexp.Compile(`[a-z]+`)
	replaceAllStr := regex5.ReplaceAllString(text5, "potato")
	fmt.Println(replaceAllStr) // potato potato potato

	fmt.Println()

	// ReplaceAllStringFunc
	text6 := "banana burger soup"
	regex6, _ := regexp.Compile(`[a-z]+`)
	replaceAllStrFunc := regex6.ReplaceAllStringFunc(text6, func(textElement string) string {
		if textElement == "banana" {
			return "potato"
		}
		return textElement
	})
	fmt.Println(replaceAllStrFunc) // potato burger soup

	fmt.Println()

	// Split
	text7 := "banana burger soup"
	regex7, _ := regexp.Compile(`[a-b]+`)
	splitStr := regex7.Split(text7, -1)
	fmt.Printf("%#v \n", splitStr) // []string{"", "n", "n", " ", "urger soup"}
	splitStr = regex7.Split(text7, 2)
	fmt.Printf("%#v \n", splitStr) // []string{"", "nana burger soup"} (di split menjadi 2 element slice)
	splitStr = regex7.Split(text7, 3)
	fmt.Printf("%#v \n", splitStr) // []string{"", "n", "na burger soup"} (di split menjadi 3 element slice)
	splitStr = regex7.Split(text7, 4)
	fmt.Printf("%#v \n", splitStr) // []string{"", "n", "n", " burger soup"} (di split menjadi 4 element slice)
}
