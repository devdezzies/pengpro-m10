package main

import "fmt"

func main() {
	var x int 
	var y bool

	fmt.Scan(&x, &y)
	if y {
		x = x * 0
	}
	fmt.Println(x)
}