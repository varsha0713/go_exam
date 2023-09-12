package main

import "fmt"

func main() {
	var prev, cur, sum, lim int
	prev, cur, sum = 0, 1, 0
	fmt.Print("enter the limit")
	fmt.Scan(&lim)
	for prev+cur < lim {
		next := prev + cur
		fmt.Println(next)
		if next%2 == 0 {
			sum += next
		}
		prev, cur = cur, next
	}
	fmt.Println("the sum of even terms is ", sum)

}
