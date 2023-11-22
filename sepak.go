package main 

import "fmt"

func main() {
	var g1, g2, g3, g4, max, min int 


	fmt.Scan(&g1, &g2, &g3, &g4)
	max = g1
	if g2 > g1  && g2 > g3 && g2 > g4 {
		max = g2
	} else if g3 > g2 && g3 > g1 && g3 > g4 {
		max = g3
	} else if g4 > g2 && g4 > g1 && g4 > g3 {
		max = g4
	}
	min = g1 
	if g2 < g1 && g2 < g3 && g2 < g4 {
		min = g2
	} else if g3 < g1 && g3 < g4 && g3 < g2 {
		min = g2
	} else if g4 < g1 && g4 < g3 && g4 < g2 {
		min = g4
	}
	fmt.Println(max, min)
}