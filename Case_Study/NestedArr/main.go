package main

import "fmt"

// membuat fungsi untuk transpose array
func transposeArray(arr [][]string) [][]string {
	var transposedArr [][]string

	// Cara 1 Menggunakan Perulangan For
	// for i := 0; i < len(arr[0]); i++ {
	// 	var innerArr []string

	// 	for j := 0; j < len(arr); j++ {
	// 		innerArr = append(innerArr, arr[j][i])
	// 	}

	// 	transposedArr = append(transposedArr, innerArr)
	// }

	// Cara 2 Menggunakan Perulangan For Range
	for indexOuter := range arr[0] {
		var innerArr []string

		for indexInner := range arr {
			innerArr = append(innerArr, arr[indexInner][indexOuter])
		}

		transposedArr = append(transposedArr, innerArr)
	}

	return transposedArr
}

func main() {
	nestedArr := [][]string{{"1", "2", "3", "4"}, {"Mark Z", "Ellon M", "Bill G", "Steve J"}, {"Facebook", "Tesla", "Microsoft", "Apple"}}

	fmt.Println("Array asli :")
	fmt.Println(nestedArr)

	fmt.Println()

	fmt.Println("Array hasil transpose :")
	fmt.Println(transposeArray(nestedArr))
}
