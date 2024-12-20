package helper

import "fmt"

var version = "1.0.0"      // Tidak bisa dipakai diluar package
var Application = "Golang" // Bisa digunakan diluar package

// Tidak bisa digunakan diluar package
func sayGoodBye(name string) string {
	return "Good Bye" + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Eko")
	fmt.Println(version)
}