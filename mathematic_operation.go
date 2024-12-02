package main

import "fmt"

func main() {

	var a = 10
	var b = 11
	var c = a + b
	var d = a - b
	var e = a * b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Augmented Assignment
	var v = 7
	v += 5 // v = v (7) + 5

	fmt.Println(v)

	// Unary Operator
	var j = 3
	j++ // j = j (3) + 1
	j++ // j = j (4) + 1
	fmt.Println(j) // Hasil: 5

	j-- // j = j (5) - 1
	j-- // j = j (4) - 1
	fmt.Println(j) // Hasil: 3

}
