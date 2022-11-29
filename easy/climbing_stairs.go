package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	n1 := 1
	n2 := 2
	for i := 0; i < n-2; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}

func main() {
	fmt.Println(climbStairs(4))
}
