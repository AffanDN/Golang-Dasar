package main

import "fmt"

func getFullName() (string, string) {
	return "Grim", "Jow"
}

// Function untuk ignore salah satu return
func getFullNameIgnoreOneReturn() (string, string) {
	return "Grim", "Jow"
}

func main() {
	// Jika hanya ingin mengambil 1 return tinggal kasih _ (undescore)
	firstname, _ := getFullName()
	fmt.Println(firstname)
}