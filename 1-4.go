package main

import "fmt"

func main() {

	var l int = 109
	var v int
	var t int
	var result int

	fmt.Scanln(&v)
	fmt.Scanln(&t)

	result = ((v*t)%l + l) % l

	fmt.Println(result)

}
