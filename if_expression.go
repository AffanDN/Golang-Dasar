package main

import "fmt"

func main() {
	name := "Forsaken"

	if name == "jow" {
		fmt.Println("Hello", name)
	} else if name == "pon"{
		fmt.Println("Ouh bukan jow ya ?")
	} else {
		fmt.Println("Terus kamu namanya siapa ?")
	}

	// if with short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}
}