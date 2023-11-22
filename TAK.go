package main

import "fmt"

func main() {
	var grade float64 

	fmt.Scan(&grade)
	if grade < 2.00 {
		fmt.Println("Poor")
	} else if grade >= 2.00 && grade < 2.75 {
		fmt.Println("Fair")
	} else if grade >= 2.75 && grade < 3.00 {
		fmt.Println("Satisfactory")
	} else if grade >= 3.00 && grade < 3.50 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Excellent")
	}
}