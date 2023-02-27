package main

import "fmt"

func main() {
	var a int

	fmt.Scanln(&a)

	fmt.Println("The next number for the number", a, "is", a+1)
	fmt.Println("The previous number for the number", a, "is", a-1)

}
