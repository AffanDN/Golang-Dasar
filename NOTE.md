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

## Tipe Data String
- String ada tipe data kumpulan karakter
- Jumlah karakter didalam String bisa nol sampai tak terhingga
- Tipe data string di Golang direpresentasikan dengan kata kunci String
- Nilai data String di Golang selalu diawali dengan karakter " (Petik 2) dan diakhiri dengan karakter " (Petik 2)


    ### Function String
    - `len("string")`: Menghitung jumlah karakter disamping
    - `"string"[number]`: Mengambil karakter pada posisi yang ditentukan

## Variable
- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Golang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kunci var, lalu diikuti dengan nama variable dan tipe datanya

## Tipe Data Variable
- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
- Namun jika kita langsung menginisilisasikan data pada variablenya, maka kita tidak wajib menyebutkan tipe data variable nya

## Kata Kunci Var
- Di Golang, kata kunci var saat membuat variable tidak lah wajib
- Asalkan saat membuat variable kita langsung menginisilisasi datanya
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut
    - <var_name> := "jow"

## Deklarasi Multiple Variable
- Di Golang kita bisa membuat variable secara sekaligus banyak
- Code yang dibuat akan lebih bagus dan  mudah dibaca

## Constant
- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Sama seperti var, namun ini menggunakan *const*
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

## Deklarasi Multiple Constant
- Sama seperti variable, di Golang juga kita bisa membuat constant secara sekaligus banyak

## Konversi Tipe Data
- Di Golang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe data lain
- Misal kita ingin mengkonversi tipe data int32 ke int64 dll

## Type Declarations
- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

## Operasi Matematik
- Penjumlahan (+)
- Pengurangan (-)
- Penrkalian (*)
- Pembagian (/)
- Sisa Pembagian (%)

    ### Agumented Assignments
    - a+= 10 : a = a + 10
    - a-= 10 : a = a - 10
    - a*= 10 : a = a * 10
    - a/= 10 : a = a / 10
    - a%= 10 : a = a % 10

    ### Unary Operator
    - ++ : a = a + 1
    - -- : a = a - 1
    - (-) : Negative
    - (+) : Positive
    - ! : Boolean Kebalikan

## Operasi Perbandingan
- Operasi Perbandingan adalah operasi untuk membandingkan dua buah nilai data
- Operasi Perbandingan adalah operasi yang menghasilkan nilai boolean (true/false)
- Jika hasilnya benar maka true
- Jika hasilnya salah maka false

    ### Operator Perbandingan
    - (>) : Lebih dari
    - (<) : Kurang dari
    - (>=): Lebih dari sama dengan
    - (<=): Kurang dari sama dengan
    - (==): sama dengan
    - (!=): Tidak sama dengan

## Operasi Boolean
- && : dan
- || : atau
- !  : Kebalikan
