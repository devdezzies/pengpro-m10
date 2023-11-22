package main 

import "fmt"

func main() {
	var c1, c2, c3, c4, c5 float64 

	fmt.Scan(&c1, &c2, &c3, &c4, &c5)
	if (c1 < c2 && c2 < c3 && c3 < c4 && c4 < c5) {
		fmt.Println("Stabil naik")
	} else if (c1 > c2 && c2 > c3 && c3 > c4 && c4 > c5) {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}