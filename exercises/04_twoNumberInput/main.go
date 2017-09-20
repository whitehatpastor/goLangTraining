package main

import "fmt"

func main() {
	small := 0
	large := 0

	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a number larger than ", small, ": ")
	fmt.Scan(&large)

	x := large / small
	fmt.Println(x)
}
