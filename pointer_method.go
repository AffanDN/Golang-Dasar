package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Jow := Man{"Jow"}
	Jow.Married()
	fmt.Println(Jow.Name)
}