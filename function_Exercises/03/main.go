package main

import (
	"fmt"
)

func greatestNumb(x ...int)  int{
	var max int
	for _, v := range x{
		if max < v{
			max = v
		}
	}
	return max

}


func main() {
	max := greatestNumb(5,8,498,5,5,98,1,192,198,1,51,681,98)
	fmt.Println(max)
}
