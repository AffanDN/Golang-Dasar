package main

import (
	"Golang-dasar/database"
	_ "Golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
