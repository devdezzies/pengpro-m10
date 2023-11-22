package main

import "fmt"

func main() {
	var besar int 
	var akhir float64

	fmt.Scan(&besar)
	akhir = float64(besar)
	if besar >= 275000 {
		akhir = akhir * 0.95
	} 
	fmt.Println(akhir)
}