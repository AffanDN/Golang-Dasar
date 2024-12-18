package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Asterisk Operator
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Solo"
	fmt.Println(address1) // Ikut berubah menjadi Solo
	fmt.Println(address2) // Berubah menjadi Solo

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
