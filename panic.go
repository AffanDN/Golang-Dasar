package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic:", message)
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("Ups ERROR ....")
	}
	
}

func main() {
	runApp(true)
	fmt.Println("Ini Main")
}