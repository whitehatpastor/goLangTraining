package main

import (
	"fmt"
)

//breaks down an int number by number and checks to see if all the numbers in the int are odd
func isOdd(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 {
		remainder := n % 10
		if remainder%2 == 0 { //checks to see if there is a leading 0
			return false
		}
		n /= 10
	}
	return true
}

//adds 2 int's passed to it and returns an int
func add(x int, y int) int {
	return x + y
}

//reverses the int passed to it and returns an int.
//also checks if there is a leading 0 when reversed. If there is then it just returns a 0.
func reverse(n int) int {
	leading := true
	r := 0
	for n > 0 {
		remainder := n % 10
		if leading {
			if remainder == 0 {
				return 0
			}
			leading = false
		}
		r *= 10
		r += remainder
		n /= 10
	}
	return r
}

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		flipped := reverse(i)
		if flipped == 0 {
			continue
		}
		sum := add(flipped, i)
		if isOdd(sum) {
			counter++
			//fmt.Println(counter, ":", i, "+", flipped, "=", sum)
		}
	}
	fmt.Println(counter)
}


/*
****************Project Euler Problem Description*****************************
*
*How many reversible numbers are there below one-billion?
*Problem 145
*Some positive integers n have the property that the sum [ n + reverse(n) ] consists entirely of odd (decimal) digits.
*For instance, 36 + 63 = 99 and 409 + 904 = *313. We will call such numbers reversible; so 36, 63, 409, and 904
*are reversible. Leading zeroes are not allowed in either n or reverse(n).
*
*There are 120 reversible numbers below one-thousand.
*
*How many reversible numbers are there below one-billion (109)
******************************************************************************/
