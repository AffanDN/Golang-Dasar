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

## Tipe Data Array
- Array adalah tipe data yang berisikan kumpulan data dengan tipe data yang sama
- Saat Membuat Array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
- Daya tampung Array tidak bisa bertambah setelah Array dibuat
- Dapat membuat Array secara langsung saat deklarasi variabel

    ### Function Array
    - `len(array)` : Untuk mendapatkan panjang array
    - `array[index]`: Mendapat data di posisi index
    - `array[index] = value`: Mengubah data di posisi index
    
## Tipe Data Slice
- Tipe data Slice adalah potongan dari Array
- Slice mirip dengan array, yang membedakan adalah ukuran Slice bisa berubah
- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di array

## Detail tipe data Slice
- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice
- Capacity adalah kapasitas dari slice (dari pointer sampai akhir), dimana length tidak boleh dari capacity

    ### Membuat Slice dari Array
    - `array[low:high]`: Membuat slice dari array dimulai dari index low sampai index sebelum high
    - `array[low:]`: Membuat slice dari array dimulai dari index low sampai index terakhir di array
    - `array[:high]`: Membuat slice dari array dimulai index 0 sampai index sebelum high
    - `array[:]`: Membuat slice dari array dimulai index 0 sampai index terakhir di array

    ### Function di Slice
    - `len(slice)`: untuk mendapatkan panjang
    - `cap(slice)`: untuk mendapatkan kapasitas dari pointer pertama
    - `append(slice, data)`: Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
    - `make([]typedata, length, capacity)`: Membuat Slice baru
    - `copy(destination, source)`: menyalin slice dari source ke destination

## Tipe Data Map
- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
- Map adalah tipe data lain yang berisikan sekumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang kita gunakan
- Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukan kedalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda. Jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

## Function Map
- `len(map)` : untuk mendapatkan jumlah data di map
- `map[key]` : mengambil data di map dengan key
- `map[key] = value` : mengubah data di map dengan key
- `make(map[TypeKey]TypeValue)` : membuat map baru
- `delete(map, key)` : menghapus data di map dengan key

## If Expression
- If adalah salah satu kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir semua bahasa pemrograman mendukung if expression

    ### Else Expression
    - Blok if akan dieksekusi ketika kondisi if bernilai true
    - Kadang kita ingin melakukan eksekusi program tertentu jika kondisi bernilai false
    - Hal ini bisa dilakukan menggunakan else expression

    ### Else if expression
    - Kadang dalam if, kita butuh membuat beberapa kondisi
    - Kasus seperti ini, kita dapat menggunakan else if expression

    ### If dengan Short Statement
    - If mendukung short statement sebelum kondisi
    - Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

## Switch Expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan switch expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

    ### Switch dengan Short Statement
    - Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan dicek kondisinya

    ### Switch Tanpa Kondisi
    - Kondisi di switch expression tidak wajib
    - Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut disetiap casenya

## For Loops
- Perulangan

    ### For Statement
    - Dalam for, kita dapat menambahkan statement. Dimana terdapat 2 statement yang bisa ditambahkan di For:
        - Init Statement, yaitu statement sebelum for dieksekusi
        - Post Statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
    
    ### For Range
    - For bisa digunakan untuk melakukan iterasi terhadap semua data collection
    - Data collection contohnya Array, Slice dan Map

## Break and Continue
- Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan
- Break digunakan untuk menghentikan seluruh perulangan
- Continue adalah digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya

## Function 
- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu Function main
- Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
- Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti dengan nama function nya dan blok kode isi function nya
- Setelah Membuat Function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function nya diikuti dengan tanda kurung buka dan kurung tutup

    ### Function Parameter
    - Saat membuat function kita membutuhkan data dari luar / external, yang dimana kita bisa sebut parameter
    - Kita bisa menambahkan parameter di function, bisa lebih dari satu
    - Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
    - Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukan data ke parameternya

    ### Function Return Value
    - Function Bisa Mengembalikan Data
    - Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
    - Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib didalam function nya kita harus mengembalikan data
    - Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

    ### Returning Multiple Value
    - Function tidak hanya dapat mengembalikan satu value, tapi juga dapat mengembalika satu value
    - Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return valuenya di function

        ### Menghiraukan Return Value
        - Multiple return value wajib ditangkap semua valuenya
        - Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (underscore)

## Named Return Values
- Biasanya saat kita memberitahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variabel secara langsung di tipe data return functionnya

## Variadic Function
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
- Apa bedanya dengan parameter biasa dengan tipe data Array ?
    - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
    - Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu cukup gunakan tanda koma

## Slice Parameter
- Kadang ada kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice
- Kita dapat menjadikan slice sebagai vararg parameter

## Function Value
- Function adalah first class citizen
- Function juga merupakan tipe data, dan bisa disimpan dalam variabel

## Function as Parameter
- Function tidak hanya bisa kita simpan didalam variabel sebagai value
- Namun juga bisa kita gunakan sebagai parameter untuk function lain

    ## Function Type Declarations
    - Kadang jika function terlalu panjang, agak ribet untuk menuliskannya didalam parameter
    - Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

## Anonymous Function
- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variabel atau parameter tanpa harus membuat function terlebih dahulu
- Hal tersebut dinamakan anonymous function, atau function tanpa nama

## Recursive Function
- Recursive Function adalah function yang memanggil function dirinya sendiri
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

## Closure
- Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yg sama
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi 

## Defer, Panic dan Recover
- Defer
    - Defer Function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
    - Defer function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi
- Panic
    - Panic Function adalah function yang bisa kita gunakan untuk menghentikan program
    - Panic Function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
    - Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
- Recover
    - Recover adalah function yang bisa kita gunakan untuk menangkap data panic
    - Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan 

## Struct
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field

    ### Membuat Data Struct
    - Struct adalah template data atau prototype data
    - Struct tidak bisa langsung digunakan
    - Namun kita bisa membuat data / object dari struct yang telah kita buat

    ### Struct Literals
    - Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
    
    ### Struct Method
    - Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
    - Namun jika kita ingin menambahkan method kedalam structs, sehingga seakan-akan sebuah struct memiliki function
    - Method adalah function

## Interface
- Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi-definisi method
- Biasanya interface digunakan sebagai kontrak

    ### Implementasi Interface
    - Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
    - Sehingga kita tidak perlu mengimplementasikan interface secara manual
    - Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

    ### Interface Kosong
    - Golang bukanlah bahasa pemrograman yang berorientasi objek
    - Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
    - Contoh di java ada java.lang.Object
    - Untuk menangani kasus seperti ini, di Golang kita bisa menggunakan interface kosong
    - Interface Kosong adalah interface yang tidak memiliki deklarasi method satu pun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
    - Interface kosong, juga memiliki type alias bernama any

        ### Penggunaan Interface Kosong
        - Ada banyak contoh penggunaan interface kosong di Golang, seperti:
            - `fmt.Println(a...interface{})`
            - `panic(v interface{})`
            - `recover() interface{}`
            - Dan lain-lain
## Nil
- Biasanya didalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
- Berbeda dengan Go-Lang, di Go-Lang saat kita buat variabel dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
- Namun di Golang ada data Nil, yaitu data kosong
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function map, slice, pointer dan channel

## Type Assertions
- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sering kali digunakan ketika kita bertemu dengan data interface kosong

    ### Type Assertions Menggunakan Switch
    - Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
    - Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
    - Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

## Pointer
- Pointer adalahh kemampuan membuat reference ke lokasi data di memori yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

    ### Pass by Value
    - Secara default di Golang semua variable itu di passing by value, bukan by reference
    - Artinya, jika kita ingin mengirim sebuah variable kedalam function, method atau variabel lain, sebenarnya yang dikirim adalah duplikasi valuenya

    ### Operator (&)
    - Untuk membuat sebuah variable dengan nilai pointer ke variabel yang lain, kita bisa menggunakan operator & diikuti dengan nama variablenya

    ### Asterisk Operator (*)
    - Saat kita mengubah variabel pointer, maka yang berubah hanya variable tersebut
    - Semua variable yang mengacu ke data yang sama tidak akan berubah
    - Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

    ### Operator New
    - Sebelumnya untuk membuat pointer dengan menggunakan operator &
    - Golang juga memiliki function new yang bisa digunakan untuk membuat pointer
    - Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

    ### Pointer di Function
    - Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
    - Oleh karena itu, jika kita mengubah data didalam function, data aslinya tidak akan pernah berubah
    - Hal ini membuat variabel menjadi aman, karena tidak akan bisa diubah
    - Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
    - Untuk melakukan ini, kita menggunakan pointer di function
    - Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya