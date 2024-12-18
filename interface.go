package main

import "fmt"

// Interface
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// Implementasi
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}
// Implementasi
func (animal Animal) GetName() string {
	return animal.Name
}

func main()  {
	person := Person{Name: "Jow"}
	animal := Animal{Name: "Cow"}
	SayHello(person)
	SayHello(animal)
}