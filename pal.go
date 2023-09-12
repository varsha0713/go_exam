package main

import (
	"fmt"
	"strconv"
)

func ispalin(n int) bool {
	str := strconv.Itoa(n)
	leng := len(str)
	for i := 0; i < leng/2; i++ {
		if str[i] != str[leng-i-1] {
			return false
		}
	}
	return true
}
func palinprod(start, end int) (int, int, int) {
	largetspali := 0
	var m1, m2 int
	for i := end; i >= start; i-- {
		for j := i; j > start; j-- {
			product := i * j
			if product < largetspali {
				break
			}
			if ispalin(product) && product > largetspali {
				largetspali = product
				m1 = i
				m2 = j

			}
		}

	}
	return largetspali, m1, m2
}
func main() {
	var start, end int
	fmt.Print("enter start:")
	fmt.Scan(&start)
	fmt.Print("enter end:")
	fmt.Scan(&end)
	result, m1, m2 := palinprod(start, end)
	fmt.Printf("the palind of %dand %d is %d", start, end, result)
	fmt.Printf("the multpiplicamnds ar%dand %d", m1, m2)

}
