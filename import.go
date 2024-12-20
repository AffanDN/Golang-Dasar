package main

import (
	"Golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Jow")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // karena tidak capital nama variabel
	// fmt.Println(helper.sayGoodBye()) // karena tidak capital nama function nya
}