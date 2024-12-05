package main

import "fmt"

func main() {

	// map[tipe_data_key]tipe_data_value
	person := map[string]string{
		"name":  "Hiko",
		"hobby": "Bermain Game",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["hobby"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)

}