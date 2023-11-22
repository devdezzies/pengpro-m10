package main 

import "fmt"

func main() {
	var kartu, diskon, cashback, bersedia bool 
	var bayar, total float64

	fmt.Scan(&bayar, &bersedia)
	kartu = bersedia 
	diskon = bayar >= 100000
	cashback = bayar >= 200000 && kartu 
	if (cashback && diskon) {
		total = bayar * 0.9 - 75000
	} else if (diskon) {
		total = bayar * 0.9
	} else if (cashback) {
		total = bayar - 75000
	}
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp.", total)

}