package main

import "fmt"

func main() {
	// slice
	// composite literal; slice literal
	x := []int{2, 3, 4}
	/// fmt.Println("x:", x)
	for i := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 2)
	y = append(y, 5)
	y = append(y, 10)
	y = append(y, 234)
	// fmt.Println("y:", y)
	for i, v := range y {
		fmt.Println(i, "-", v)
	}
}