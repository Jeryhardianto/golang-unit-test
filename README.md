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

