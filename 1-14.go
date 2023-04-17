package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	days := (m + n - 1) / n
	fmt.Println(days)
}
