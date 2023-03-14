package main

import "fmt"

func main() {

	var num int
	var a, b, c int

	fmt.Scanln(&num)

	if num < 0 {
		num = -num
	}

	a = num / 100

	b = (num / 10) % 10

	c = num % 10

	fmt.Println(a + b + c)

}
