package main

import "fmt"

// func ..() <tipe data return> {}
func getHello(name string) string {
	return "Hello " + name
}

// Dapat disimpan hasil function nya kedalam variabel
func main() {
	result := getHello("Jow")
	fmt.Println(result)
}