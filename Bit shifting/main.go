package main

import "fmt"

func main() {
	// bit shifting binary
	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5)
}
