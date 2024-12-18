package main

import "fmt"

func Ups() any {
	// return true
	return "Ups haha"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}