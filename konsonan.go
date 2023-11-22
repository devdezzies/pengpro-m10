package main

import "fmt"

func main() {
	var huruf string

	fmt.Scan(&huruf)
	if (huruf == "A" || huruf == "a") || (huruf == "E" || huruf == "e") || (huruf == "U" || huruf == "u") || (huruf == "I" || huruf == "i") || (huruf == "O" || huruf == "o") {
		fmt.Println("bukan konsonan")
	} else {
		fmt.Println("konsonan")
	}
}