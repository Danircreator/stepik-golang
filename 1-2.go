package main

import "fmt"

func main() {

	var a int
	var b int
	var result int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	result = (b - b%a) / a

	fmt.Println(result)

}
