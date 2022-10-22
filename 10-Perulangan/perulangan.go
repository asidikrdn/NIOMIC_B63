package main

import "fmt"

func main() {
	// Perulangan for layaknya fungsi for pada umumnya
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan for ke-", i)
	}

	fmt.Println()

	// Perulangan for layaknya fungsi while
	var x = 1
	for x <= 5 {
		fmt.Println("Perulangan while ke-", x)
		x++
	}

	fmt.Println()

	// Perulangan for layaknya fungsi do-while
	var y = 1
	for {
		fmt.Println("Perulangan do-while ke-", y)
		y++
		if y == 5 {
			break
		}
	}
	fmt.Println()
	var z = 1
	for {
		fmt.Printf("Perulangan do-while ke-%d \n", z)
		z++
		if z > 5 {
			break
		}
	}
}
