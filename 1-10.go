package main

import "fmt"

func main() {

	var a int
	var hours, minutes int

	fmt.Scanln(&a)

	hours = a % (60 * 24) / 60
	minutes = a % 60

	fmt.Println(hours, minutes)

}
