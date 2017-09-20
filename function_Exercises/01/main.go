package main

import "fmt"

func takesInt(x int)  (int, bool){
	y := x/2
	var isEven bool
	if x%2 == 0{
		isEven = true
	}else{
		isEven = false
	}
	return y,isEven
}

func main() {
	fmt.Println(takesInt(2))
}
