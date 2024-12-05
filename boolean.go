package main

import "fmt"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	var nilaiAKhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool =  nilaiAKhir > 80
	var lulusAbsensi bool =  absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	var lulus2 bool = lulusNilaiAkhir || lulusAbsensi

	fmt.Println(lulus)
	fmt.Println(lulus2)

}