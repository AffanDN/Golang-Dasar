package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Parameter struct sebelum nama function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	rully := Customer{Name: "Rully"}
	rully.sayHello("Jow")
}