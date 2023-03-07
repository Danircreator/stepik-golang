package main

import "fmt"

func main() {

	var a int

	fmt.Scanln(&a)

	if a < 0 {
		a = a / (-10)
		fmt.Println(a)
	} else {
		a = a / 10
		fmt.Println(a)
	}

}
