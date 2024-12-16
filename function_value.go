package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// function dimasukan kedalam variabel, tanpa ()
	goodbye := getGoodBye
	fmt.Println(goodbye("Jow"))
}