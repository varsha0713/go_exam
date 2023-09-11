package main

import "fmt"

func main() {
	var prev, cur, sum, limit int
	prev, cur, sum = 0, 1, 0
	fmt.Print("enter the limit of the series:")
	fmt.Scan(&limit)
	for prev+cur < limit {
		next := prev + cur
		fmt.Println(next)
		if next%2 == 0 {
			sum += next

		}
		prev, cur = cur, next
	}
	fmt.Print("the sum of even terms", sum)
}
