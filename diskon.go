package main

import "fmt"

func main() {
	var ujian bool 
	var pembelian float64

	fmt.Scan(&pembelian, &ujian)
	if ujian {
		pembelian = pembelian * 0.65
	}
	fmt.Println(pembelian)
}