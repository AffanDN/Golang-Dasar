package main

import "fmt"

func logging() {
	fmt.Println("Selesai Bekerja Run Application nya")
}

func runApplication()  {
	// defer: akan menjadi diakhir yg dijalankan pada runApplication ini
	defer logging()
	fmt.Println("Running Application ...")
}

func main() {
	runApplication()
}