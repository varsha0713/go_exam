package main

import "fmt"

func LF(n int) int {
	var largest int

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}

	return largest
}
func main() {
	var cap int
	fmt.Print("enter the num")
	fmt.Scan(&cap)
	fmt.Println("the lpf of ", cap, "is", LF(cap))
}
