package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Okta"
	middleName = "Kuwan"
	lastName = "Bimo"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a,b,c)
}