package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	next := n + 1
	prev := n - 1

	fmt.Printf("The next number for the number %d is %d.\n", n, next)
	fmt.Printf("The previous number for the number %d is %d.", n, prev)
}
