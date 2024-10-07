package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// menggagalkan unit test
// Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
// Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
// Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

// t.Fail() dan t.FailNow()
// Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa bedanya?
// Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
// FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

// t.Error(args...) dan t.Fatal(args...)
// Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
// Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
// Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
// Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

// Assertion
// Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
// Apalagi jika result data yang harus di cek itu banyak
// Oleh karena itu, sangat disarankan untuk menggunakan Assertion untuk melakukan pengecekan
// Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga kita butuh menambahkan library untuk melakukan assertion ini

// Testify
// Salah satu library assertion yang paling populer di Go-Lang adalah Testify
// Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
// https://github.com/stretchr/testify
// Kita bisa menambahkannya di Go module kita : go get github.com/stretchr/testify

// Table Test
// Sebelumnya kita sudah belajar tentang sub test
// Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
// Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan konsep table test
// Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi hasil dari unit test
// Lalu slice tersebut kita iterasi menggunakan sub test
func TestTableTest(t *testing.T) {
	// Table Test
    tests := []struct {
        input    int
        expected int
    }{
        {1, 2},
        {2, 3},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("Test Case %d", tt.input), func(t *testing.T) {
            result := tt.input + 1
            assert.Equal(t, tt.expected, result, "Test Case Failed")
        })
    }
}

// Sub Test
// Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
// Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
// Untuk membuat sub test, kita bisa menggunakan function Run()
// Menjalankan Hanya Sub Test
// Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah : go test -run TestNamaFunction
// Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah : go test -run TestNamaFunction/NamaSubTest
// Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah : go test -run /NamaSubTest
func TestSubTest(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
        result := 1 + 1
        assert.Equal(t, 2, result, "Test Case 1 Failed")
    })
    t.Run("Test Case 2", func(t *testing.T) {
        result := 1 + 2
        assert.Equal(t, 3, result, "Test Case 2 Failed")
    })
}

// Before dan After Test
// Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
// Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
// Untungnya di Go-Lang terdapat fitur yang bernama testing.M
// Fitur ini bernama Main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test
// testing.M
// Untuk mengatur ekeskusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
// Jika terdapat function TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
// Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
// Ingat, function TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap function unit test
func TestMain(m *testing.M) {
	fmt.Println("Before Testing")

	m.Run()

	fmt.Println("After Testing")
}

// Skip Test
// Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
// Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
// Untuk membatalkan unit test kita bisa menggunakan function Skip()
func TestSkip(t *testing.T) {
	// Kita ingin membatalkan unit test ini karena kita mau mencoba skip test lain
    t.Skip("skipping test")
    fmt.Println("This test will not be executed")
}

// assert vs require
// Testify menyediakan dua package untuk assertion, yaitu assert dan require
// Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
// Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
func TestHelloWorldRequire(t *testing.T) {
    result := HelloWorld("Ivan")
    require.Equal(t, "Hello, my name is Ivan", result)
}

func TestIvan(t *testing.T) {
    result := HelloWorld("Ivan")
    assert.Equal(t, "Hello, my name is Ivan", result, "Testing failed")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ivan1")
	if result != "Hello, my name is Ivan" {
		// error
		t.FailNow()
	}
}