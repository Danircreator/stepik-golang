package main

import "fmt"

func main() {

	var h1, m1, s1, h2, m2, s2, res int

	fmt.Scanln(&h1)
	fmt.Scanln(&m1)
	fmt.Scanln(&s1)
	fmt.Scanln(&h2)
	fmt.Scanln(&m2)
	fmt.Scanln(&s2)

	h1 = h1 * 60 * 60
	m1 = m1 * 60
	s1 = s1 * 1

	h2 = h2 * 60 * 60
	m2 = m2 * 60
	s2 = s2 * 1

	res = (h2 + m2 + s2) - (h1 + m1 + s1)

	fmt.Println(res)

}
