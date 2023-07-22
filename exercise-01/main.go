package main

import "fmt"

func main() {
	fmt.Println("this is heloo 🥶🥶🥶")

	const nama, umur = "pingu", 40
	fmt.Printf("%s is %d years old. \n", nama, umur)
	fmt.Printf("%v is %v years old. \t and the type is %T and %T", nama, umur, nama, umur)
}
