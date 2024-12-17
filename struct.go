package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var jow Customer
	jow.Name = "Grim Jow"
	jow.Age = 30
	jow.Address = "London"

	fmt.Println(jow)

	joko := Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 44,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 29}
	fmt.Println(budi)
}