# Golang Unit Test 

```
SLIDE PRESENTASI UNIT TEST BY PROGRAMMER ZAMAN NOW
https://docs.google.com/presentation/d/1XxMEaA-JsPHr9BUw2oIOPlEL_psI3EaUFUpuvdlDB_Q/edit#slide=id.gb233370586_0_320
```

```
1. Unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah method
2. Unit test biasanya dibuat kecil dan cepat, oleh karena itu biasanya kadang kode unit test lebih banyak dari kode program aslinya, karena semua skenario pengujian akan dicoba di unit test
3. Unit test bisa digunakan sebagai cara untuk meningkatkan kualitas kode program kita
```

## Testing Package
```
1. Di bahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test
2. Berbeda dengan Go-Lang, di Go-Lang sudah untuk unit test sudah disediakan sebuah package khusus bernama testing
3. Selain itu untuk menjalankan unit test, di Go-Lang juga sudah disediakan perintah nya
4. Hal ini membuat implementasi unit testing di golang sangat mudah dibanding dengan bahasa pemrograman yang lain
5. https://golang.org/pkg/testing/

=> Testing Package
1. testing.T
2. testing.M
3. testing.B
```

## Aturan File Test
```
1.Go-Lang memiliki aturan cara membuat file khusus untuk unit test
2.Untuk membuat file unit test, kita harus menggunakan akhiran \_test
3.Jadi kita misal kita membuat file hello_world.go, artinya untuk membuat unit testnya, kita harus membuat file hello_world_test.go
```

## Aturan Function Unit Test
```
1. Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test
2. Nama function untuk unit test harus diawali dengan nama Test
3. Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama TestHelloWorld
4. Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawalin dengan kata Test
5. Selanjutnya harus memiliki parameter (t \*testing.T) dan tidak mengembalikan return value
```

## Command Run File Test
```
1. go test -v => menjalankan semua file test yang ada 
2. go test -v -run TestHelloWorld => menjalankan file test yang ada dengan nama TestHelloWorld
NOTE : go test -v -run TestHelloWorld bisa dijalankan hanya didalam folder location yang sama dengan file testnya ketika dijalankan dari luar folder
maka akan error "no Go files in E:\development\golang-dev\src\golang-unit-test
".
3. go test -v ./... => menjalankan semua file test dari luar folder
4. go test -v ./... -run TestHelloWorld => menjalankan satu file test dari luar folder dengan nama TestHelloWorld
```

## Menggagalkan Test
```
1. Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
2. Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
3. Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test
```
### t.Fail() dan t.FailNow()
```
1. Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa bedanya?
2. Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
3. FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
```

### t.Error(args...) dan t.Fatal(args...)
```
1. Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
2. Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
3. Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
4. Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
```

## Assertion
```
1. Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
2. Apalagi jika result data yang harus di cek itu banyak
3. Oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan
4. Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini
```

### Testify
```
1. Salah satu library assertion yang paling populer di Go-Lang adalah Testify
2. Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
3. https://github.com/stretchr/testify
4. Kita bisa menambahkannya di Go module kita :
go get github.com/stretchr/testify
```

### assert vs require
```
1. Testify menyediakan dua package untuk assertion, yaitu assert dan require
2. Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
3. Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
```

## Skip Test
```
1. Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
2. Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
3. Untuk membatalkan unit test kita bisa menggunakan function Skip()
```

## Before dan After Test
```
1. Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
2. Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
3. Untungnya di Go-Lang terdapat fitur yang bernama testing.M
4. Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test
```
## testing.M
```
1. Untuk mengatur ekeskusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
2. Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
3. Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
4. Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test

```

## Sub Test
```
1. Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
2. Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
3. Untuk membuat sub test, kita bisa menggunakan function Run()
```

### Menjalankan Hanya Sub Test
```
1. Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
go test -run TestNamaFunction
2. Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
go test -run TestNamaFunction/NamaSubTest
3. Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
go test -run /NamaSubTest

```

## Table Test
```
1. Sebelumnya kita sudah belajar tentang sub test
2. Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
3. Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan konsep table test
4. Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi hasil dari unit test
5. Lalu slice tersebut kita iterasi menggunakan sub test
```

## Mock
```
1. Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal
2. Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
3. Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita mau
4. Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

```

## Testify Mock
```
1. Untuk membuat mock object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion
2. Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
3. Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik
4. Mari kita buat contoh kasus

```

## Contoh MOCK => Aplikasi Query Ke Database
```
1. Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
2. Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
3. Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

```

## Benchmark
```
1. Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark
2. Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
3. Benchmark di Go-Lang dilakukan dengan cara secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
4. Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package
```

## testing.B
```
1. testing.B adalah struct yang digunakan untuk melakukan benchmark.
2. testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dan lain-lain
3. Yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
4. Salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark
```

## Cara Kerja Benchmark
```
1. Cara kerja benchmark di Go-Lang sangat sederhana
2. Gimana kita hanya perlu membuat perulangan sejumlah N attribute
3. Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark nya dalam waktu

```
## Membuat Benchmark
### Benchmark Function
```
1. Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama function nya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
2. Selain itu, harus memiliki parameter (b \*testing.B)
3. Dan tidak boleh mengembalikan return value
4. Untuk nama file benchmark, sama seperti unit test, diakhiri dengan \_test, misal hello_world_test.go
```
### Menjalankan Benchmark
```
1. Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :
go test -v -bench=.
2. Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :
go test -v -run=NotMathUnitTest -bench=.
3. Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah :
go test -v -run=NotMathUnitTest -bench=BenchmarkTest
4. Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :
go test -v -bench=. ./...

```

## Sub Benchmark
```
Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()

```
### Menjalankan Hanya Sub Benchmark
```
1. Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
2. Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :
go test -v -bench=BenchmarkNama/NamaSub

```

## Table Benchmark
```
1. Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga
2. Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function

```







