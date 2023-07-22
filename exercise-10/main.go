package main

import "fmt"

var x int

func main() {
	fmt.Println(x)

	y := 42
	fmt.Println(y)

	a, b := 43, "heppy"
	fmt.Println(a, b)

	var c float64 = 42.02
	fmt.Printf("%v is of the type %T\n", c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)
}
