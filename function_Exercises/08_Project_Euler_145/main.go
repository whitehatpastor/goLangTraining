package main

import (
	"fmt"
	"strconv"
)

func isReversible(n int) bool {

	//check 0 in the end
	if n%10 == 0 {
		return false
	}

	number := n

	reversed := reverse_int(number)

	sum := reversed + n

	for sum > 0 {

		if (sum%10)%2 == 0 {
			fmt.Println((reversed + n))
			return false
		}

		sum = sum / 10
	}
	fmt.Println((reversed + n), "reversible")

	return true
}

func reverse_int(value int) int {

	intString := strconv.Itoa(value)

	newString := ""

	for x := len(intString); x > 0; x-- {
		newString += string(intString[x-1])
	}

	newInt, err := strconv.Atoi(newString)

	if err != nil {
		fmt.Println("Error converting string to int")
	}

	return newInt
}

func main() {
	count := 0
	for i := 1; i < 1000; i++ {
		if isReversible(i) {
			//fmt.Println(i)
			count++
		}
	}

	fmt.Println("Reversible Count: ", count)
}
