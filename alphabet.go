package main

import "fmt"

func main() {
	var inp byte 

	fmt.Scanf("%c", &inp)
	if (inp <= 'z' && inp >= 'a') || (inp >= 'A' && inp <= 'Z') {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}