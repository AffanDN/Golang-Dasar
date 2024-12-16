package main

import "fmt"

func main() {
	counter := 1 // Init Statement

	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++ // Post Statement
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i, "|")
	}

	names := []string{"Banana", "Jow", "Grimm"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("Index:",index,"=",value)
	}

	// Jika index nya tidak butuh maka untuk bagian index diganti jadi _
	for _, value := range names {
		fmt.Println("Value:",value)
	}
}