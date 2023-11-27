package main 

import "fmt"

func main() {
	var orang int 
	fmt.Scan(&orang)

	if orang >= 25 {
		fmt.Println(3, "Dewasa", 5, "Kecil", orang - 25, "Tidak bisa Ikut")
	} else if orang >= 15 && orang < 25 {
		if (orang - 15) % 2 > 0 {
			fmt.Println(3, "Dewasa", (orang - 15) / 2 + 1, "Kecil")
		} else {
			fmt.Println(3, "Dewasa", (orang - 15) / 2, "Kecil")
		}
	} else if orang < 15 && orang != 10 && orang != 5 {
		if orang % 15 > 0 {
			fmt.Println(orang / 5 + 1, "Dewasa")
		} else {
			fmt.Println(orang / 5, "Dewasa")
		}
	} else if orang == 5 || orang == 10 {
		fmt.Println(orang / 5, "Dewasa")
	}
}

