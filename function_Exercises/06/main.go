package main

import "fmt"

func foo(x ...int)  {
	for _, v := range x{
		fmt.Println(v)
	}
}



func main() {
	foo(1, 2)
	fmt.Println(` `)
	foo(1, 2, 3)
	fmt.Println(` `)
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(` `)
	foo(aSlice...)
	foo()

}
