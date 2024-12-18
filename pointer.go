package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by Value
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // Copy value from address1

	// address2.City = "Solo"
	// fmt.Println(address1) // Tidak berubah
	// fmt.Println(address2) // Berubah menjadi Solo

	// Pass by reference
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Solo"
	fmt.Println(address1) // Ikut berubah menjadi Solo
	fmt.Println(address2) // Berubah menjadi Solo
}