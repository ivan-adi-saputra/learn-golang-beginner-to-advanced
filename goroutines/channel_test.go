package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Membuat Channel
// Channel di Go-Lang direpresentasikan dengan tipe data chan
// Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
// Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut

// Mengirim dan Menerima Data dari Channel
// Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
// Untuk mengirim data, kita bisa gunakan kode : channel <- data
// Sedangkan untuk menerima data, bisa gunakan kode : data <- channel
// Jika selesai, jangan lupa untuk menutup channel menggunakan function close()


func TestChannel(t *testing.T) {
	ch := make(chan string)

	// mengirimkan data 
	// ch <- "Ivan"

	// menerima data
	// data := <- ch // data akan berisi data dari ch

	go func ()  {
		ch <- "Ivan"
	}()
		
	data := <- ch	
	fmt.Println(data)
	// defer memastikan close() tereksekusi walaupun aplikasi error atau berhasil
	defer close(ch)
}

// channel sebagai parameter
// Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
// Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference). 
// Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut

func GiveMeChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ivan adi saputra"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeChannel(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	defer close(channel)
}

// Channel In dan Out
// Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
// Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
// Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerima data)

func SendData(channel chan<- string) {
    channel <- "Ivan adi saputra"
}

func ReceiveData(channel <-chan string) {
    data := <- channel
    fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)

    go SendData(channel)
    go ReceiveData(channel)

    time.Sleep(5 * time.Second)
    defer close(channel)
}