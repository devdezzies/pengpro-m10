package main 

import "fmt"

func main() {
	var inp int 

	fmt.Scan(&inp)
	if inp >= 0 {
		fmt.Println(inp)
	} else {
		fmt.Println(inp * -1)
	}
}