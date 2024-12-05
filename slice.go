package main

import "fmt"

func main() {
	names := [...]string{
		"gum", "gim", "gam", "gem", "gom", "guam",
	}

	slice := names[2:4] // dimulai dari 2 dan berhenti sebelum 4 : 2-3
	fmt.Println("--0--",slice)

	slice1 := names[:3]
	fmt.Println("--1--",slice1)

	slice2 := names[4:]
	fmt.Println("--2--",slice2)

	slice3 := names[:]
	fmt.Println("--3--",slice3)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	// fmt.Println(slice[2])

	// Append Slice
	days := [...]string{
		"senin","selasa","rabu","kamis","jumat","sabtu","minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur NGAB")
	daySlice2[0] = "Haha"
	fmt.Println(daySlice2) // [Ups]
	fmt.Println(days) // [Ups]

	// Make Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Jow"
	newSlice[1] = "Jaw"

	fmt.Println("-- new slice --",newSlice)
	fmt.Println("-- len slice --",len(newSlice))
	fmt.Println("-- cap slice --",cap(newSlice))

	newSlice2 := append(newSlice, "Kora")
	fmt.Println(newSlice2)
	fmt.Println("-- len slice --",len(newSlice2))
	fmt.Println("-- cap slice --",cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println("Copy Slice: ",toSlice)

	// Array vs Slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5} // Tidak disebutkan ukurannya

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}