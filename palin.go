package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i <= length/2; i++ {
		if str[i] != str[length-i-1] {
			return false

		}
	}
	return true
}

func largestPalindromproduct(start, end int) (int, int, int) {
	largestpalindrom := 0
	var m1, m2 int
	for i := end; i >= start; i-- {
		for j := i; j >= start; j-- {
			product := i * j
			if product < largestpalindrom {
				break
			}
			if isPalindrome(product) && product > largestpalindrom {
				largestpalindrom = product
				m1 = i
				m2 = j
			}
		}
	}
	return largestpalindrom, m1, m2
}
func main() {
	var start, end int
	fmt.Println("enter start num:")
	fmt.Scan(&start)
	fmt.Print("enter the end num:")
	fmt.Scan(&end)
	result, m1, m2 := largestPalindromproduct(start, end)
	fmt.Printf("The largest palindrom product of %d and %d is :%d\n", start, end, result)
	fmt.Printf("the meultiplicands are %d and %d\n", m1, m2)
}
