package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Grim"
	names[1] = "Grum"
	names[2] = "Grom"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		85,
		80,
	}
	fmt.Println(values)

	// var names [...]string = jika seperti ini tidak bisa, karena [...] hanya bisa dideklarasikan isinya
	var nilai = [...]int{ // [...] = untuk tidak mendefinisikan jumlah array nya
		10,
		20,
		30,
		42,
	}
	// Di Golang tidak ada operasi menghapus data array, hanya bisa dikosongkan

	fmt.Println(nilai)
	fmt.Println(len(nilai))
	nilai[0] = 100
	fmt.Println(nilai)
}
