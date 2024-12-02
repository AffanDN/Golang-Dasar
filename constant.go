package main

import "fmt"

func main() {

	const firstname string = "Kruise"
	const lastname string = "Lactora"

	const (
		secondname string = "Daytona"
		thirdname  string = "Antony"
	)

	// Error
	// firstname = "Jino"
	// lastname = "Jira"

	fmt.Println(firstname)
	fmt.Println(secondname)
	fmt.Println(thirdname)
	fmt.Println(lastname)
}