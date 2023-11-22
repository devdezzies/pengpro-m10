package main 

import "fmt"

func main() {
	var b1, b2 float64

	fmt.Scan(&b1, &b2)
	if (b1 < b2) {
		fmt.Println("Peningkatan sebesar", b2 - b1)
	} else if (b1 > b2) {
		fmt.Println("Penurunan sebesar", b1 - b2)
	} else {
		fmt.Println("Tetap")
	}
}