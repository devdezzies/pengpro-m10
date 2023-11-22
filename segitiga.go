package main 

import "fmt"

func main() {
	var s1, s2, s3 int 

	fmt.Scan(&s1, &s2, &s3)
	if (s1 == s2 && s2 == s3) {
		fmt.Println("Segitiga Sama Sisi")
	} else if (s1 == s3) || (s1 == s2) || (s2 == s3) {
		fmt.Println("Segitiga Sama Kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}