package main

import "fmt"

func main() {

	var a int
	fmt.Scanln(&a)

	a = (a / 10) % 10

	fmt.Println(a)
}
