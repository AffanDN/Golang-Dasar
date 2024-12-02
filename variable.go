package main

import "fmt"

func main() {
	var nama = "Affan Dharmawan"
	fmt.Println(nama)

	nama = "Affan Nusli"
	fmt.Println(nama)

	// Tanpa Kata Kunci Var
	name := "Banana Joe"
	fmt.Println(name)

	// Multiple Variable
	var (
		firstname = "Jojo"
		lastname = "Jaja"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}