package main

import "fmt"

func main() {

	var rub int
	var kop int
	var res int
	var n int

	fmt.Scanln(&rub)
	fmt.Scanln(&kop)
	fmt.Scanln(&n)

	res = (rub*100 + kop) * n

	rub = res / 100
	kop = res % 100

	fmt.Println(rub, kop)

}
