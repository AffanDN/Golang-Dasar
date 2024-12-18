package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// Error akan terjadi Panic
	// resultInt := result.(int) // Panic
	// fmt.Println(resultInt)

	// Switch Expression
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}

}
