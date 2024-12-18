package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	alamat1 := new(Address)
	//sama seperti alamat1 := &Address{}

	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1) // alamat 1 berubah
	fmt.Println(alamat2)
}