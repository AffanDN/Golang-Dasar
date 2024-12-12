package main

import "fmt"

func main() {
	name := "Lika Liku"

	switch name {
	case "Lika Liku":
		fmt.Println("Hello Lika Liku")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Ini siapa kalo begitu ?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama nya kepanjangan")
	case false:
		fmt.Println("Nama nya pas")
	}

	fruit := "Buah Naga"
	panjang := len(fruit)
	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu panjang")
	case panjang > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sesuai")
	}
}
