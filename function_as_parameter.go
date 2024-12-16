package main

import "fmt"

// Function Type Declarations
type Filter func(string)string // alias

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Naga" {
		return "...."
	} else {
		return name
	}
}

func main()  {
	sayHelloWithFilter("Naga", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Jow", filter)
}