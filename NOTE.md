# Go-Lang Dasar
Welcome to Go-Lang Dasar :)

## Sejarah Go-Lang

- Dibuat di Google menggunakan bahasa pemrograman C
- Dirilis ke Public sebagai open source pada tahun 2009
- Go-Lang populer sejak digunakan untuk membuat Docker pada tahun 2011
- Saat ini mulai banyak teknologi baru yang dibuat menggunakan bahasa Golang
- Saat ini banyak digunakan untuk pembuatan API di microservice

## Alasan Belajar Go-Lang
- Bahasa Go sangat sederhana, tidak butuh waktu lama untuk mempelajarinya
- Golang mendukung baik Concurrency Programming, dimana kita hidup di zaman Multicore Processor
- Golang mendukung Garbage Collector, sehingga tidak membutuhkan management memory secara manual seperti di Bahasa C
- Bahasa yang sedang naik daun

## Instalasi Golang
- [Link Go-Lang](https://golang.org/)
- Download Compiler Go-Lang
- Install Compiler Go-Lang
- Cek Menggunakan perintah: `go version`

## Text Editor / IDE
- Visual Studio Code
- JetBrains GoLand

## Membuat Project
- Project di Go-Lang disebut sebagai Module
- Untuk membuat module, kita dapat menggunakan perintah sebagai berikut di tempat folder kita akan membuat folder tersebut: `go mod init nama-module`

## Main Function
- Pada Go-Lang mirip seperti bahasa pemrograman C/C++, dimana perlu ada yang namanya **Main** function
- Main Function adalah sebuah fungsi yang akan dijalankan ketika program berjalan
- Untuk membuat function, kita bisa menggunakan kata kunci **Func**
- **Main Function** harus terdapat didalam **Main package**
- Titik koma di Golang tidaklah wajib, artinya kita bisa menambahkan titik koma atau tidak, diakhir kode program

## Println
- Untuk menulis tulisan, kita perlu melakukan import module `fmt` terlebih dahulu
- Seperti pada halnya dengan Java

## Kompilasi File Go-Lang
1. `go build`: untuk mencompile project
2. `ls -l` : untuk menampilkan file direktori beserta informasi
3. `./nama file.exe` : untuk menjalankan filenya (**bash**)
    - `nama file` : untuk menjalankan filenya (**Command Prompt**)

## Menjalankan Tanpa Kompilasi
1. `go run nama-file` : untuk menjalankan tanpa kompilasi

## Multiple Main Function
- Di Golang, function dalam module / project adalah unik, artinya kita tidak boleh membuat nama function yang sama
- Oleh karena itu, jika kita membuat file baru, misal sample.go, lalu membuat nama function yang sama yaitu main.
- Maka kita tidak bisa melakukan build module, karena main function tersebut duplikat dengan yang ada di main function helloworld.go
    - **Untuk sementara kita menggunakan go run dulu untuk menjalankan satu per satu file terlebih dahulu**

## Tipe Data Number
- Ada 2 jenis tipe data Number yaitu: 
    - Integer
        - Tipe Data Integer (1) *dari **minus**: int8, int16 ...
        - Tipe Data Integer (2) *dari **0**: uint8, uint16 ... (UnAssign)
    - Floating Point: float32, float64, complex64, complex128
    > *Semakin besar kita menggunakan jenis tipe data nya maka semakin besar pula memory yang digunakan*

- Alias
    - byte: uint8
    - rune: int32
    - int: Minimal int32
    - uint: Minimal uint32       

## Tipe Data Boolean
- Tipe data yang memiliki dua nilai, yaitu benar (true) dan salah (false)
- Di Go-Lang, tipe data boolean direpresentasikan menggunakan kata kunci bool




