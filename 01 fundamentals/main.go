package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {
	// slice
	// composite literal; slice literal
	// x := []int{2, 3, 4}
	/// fmt.Println("x:", x)
	// for i := range x {
	// 	fmt.Println(i, "-", x[i])
	// }
	// map1 := map[string]int{"Abc": 1, "aBc": 2, "abC": 3}
	// for k, v := range map1 {
	// 	fmt.Println(k, "-", v)
	// }
	// for k := range map1 {
	// 	fmt.Println(k, "-", map1[k])
	// }

	// y := make([]int, 0, 2)
	// y = append(y, 5)
	// y = append(y, 10)
	// y = append(y, 234)
	// // fmt.Println("y:", y)
	// for i, v := range y {
	// 	fmt.Println(i, "-", v)
	// }
	// y := make(map[string]int)
	// y["Abc"] = 1
	// y["aBc"] = 2
	// for k, v := range y {
	// 	fmt.Println(k, "-", v)
	// }
	p1 := person{"Youktinath", "Bhowmick"}
	p2 := person{"John", "Finis"}
	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(lastname(p1))
	// fmt.Println(lastname(p2))
	p1.speak()
	p2.speak()

	// a1 := struct{
	// 	breed string
	// 	name string
	// 	}{
	// 	"German Sheperd",
	// 	"Shasta",
	// 	}
	// fmt.Println(a1)
}

// func lastname(p person) string {
// 	return p.lname
// }

func (p person) speak() {
	fmt.Println(p.fname, `says: "Hello"`)
}