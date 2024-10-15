package goroutines

// Membuat Goroutine
// Untuk membuat goroutine di Golang sangatlah sederhana
// Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
// Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
// Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, World!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // Membuat Goroutine baru / berjalan secara syncronus

    // Kode program lanjut berjalan ke kode program selanjutnya tanpa menunggu Goroutine yang kita buat selesai
	fmt.Println("This is main function")

	// Menunggu Goroutine yang kita buat selesai
	time.Sleep(1 * time.Second)
}

// Goroutine Sangat Ringan
// Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
// Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
// Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

func Display(i int) {
	fmt.Printf("Goroutine %d is running...\n", i)
}

func TestCreateMultipleGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
        go Display(i) // Membuat Goroutine baru
    }

    // Menunggu Goroutine yang kita buat selesai
    time.Sleep(10 * time.Second)
}

// Buffered Channel
// Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
// Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
// Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
// Untuknya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

// Buffer Capacity
// Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
// Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer.
// Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
// Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data

func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 5) // Membuat Channel dengan buffer capacity 5

    // Mengirimkan data ke Channel
    ch <- "Data 1"
    ch <- "Data 2"
    ch <- "Data 3"
    ch <- "Data 4"
    ch <- "Data 5"

    // Menunggu data di Channel selesai dikirim
    // Jika data belum selesai dikirim, maka kita akan menunggu sampai data terakhir keluar
    // Jika buffer capacity 5, maka ketika data 6 keluar, maka kita akan diminta menunggu sampai data 1 keluar
    <-ch

	fmt.Println(cap(ch)) // untuk melihat capasitas channel
	fmt.Println(len(ch)) // untuk melihat jumlah data yang ada di channel

	defer close(ch)
}

// Range Channel
// Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
// Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
// Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
// Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
// Ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual

func TestRangeChannel(t *testing.T) {
    ch := make(chan string)

    // Mengirimkan data ke Channel
    go func() {
        ch <- "Data 1"
        ch <- "Data 2"
        ch <- "Data 3"
        close(ch)
    }()

    // Menerima data dari Channel
    for data := range ch {
        fmt.Println(data)
    }
}

// Select Channel
// Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
// Lalu kita ingin mendapatkan data dari semua channel tersebut
// Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
// Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	defer close(ch1)
	defer close(ch2)

	// Menggunakan goroutine untuk mengirim data ke kanal secara paralel
	go func() {
		ch1 <- "Data dari ch1"
	}()
	go func() {
		ch2 <- "Data dari ch2"
	}()

	counter := 0

	for {
		select {
		case data := <-ch1:
			fmt.Println("Received from ch1:", data)
			counter++
		case data := <-ch2:
			fmt.Println("Received from ch2:", data)
			counter++
		default: 
			fmt.Println("Waiting...")
		}
		if counter == 2 {
			break
		}
	}
}


// Race Condition
// Masalah Dengan Goroutine
// Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
// Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
// Hal ini bisa menyebabkan masalah yang namanya Race Condition

// Race Condition Example
func TestRaceCondition(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
        go func() {
			for i := 0; i < 100; i++ {
				counter++
			}
		}()
    }

    // Menunggu Goroutine yang kita buat selesai
    time.Sleep(5 * time.Second)

    fmt.Println("Counter:", counter)
}

// Mutex (Mutual Exclusion)
// Untuk mengatasi masalah race condition tersebut, di Go-Lang terdapat sebuah struct bernama sync.Mutex
// Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
// Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi
// Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
func TestMutex(t *testing.T) {
	counter := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
        go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
    }

    // Menunggu Goroutine yang kita buat selesai
    time.Sleep(5 * time.Second)

    fmt.Println("Counter:", counter)
}

// RWMutex (Read Write Mutex)
// Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
// Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
// Di Go-Lang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write

// RWMutex Example
func TestRWMutex(t *testing.T) {
    var rwMutex sync.RWMutex
    counter := 0

    go func() {
        rwMutex.Lock()
        for i := 0; i < 1000; i++ {
            counter++
        }
        rwMutex.Unlock()
    }()

    go func() {
        rwMutex.RLock()
        for i := 0; i < 500; i++ {
            counter++
        }
        rwMutex.RUnlock()
    }()

    time.Sleep(5 * time.Second)

    fmt.Println("Counter:", counter)
}

// Deadlock
// Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi adalah Deadlock
// Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
// Sekarang kita coba simulasikan proses deadlock

func TestDeadlock(t *testing.T) {
	var lock1, lock2 sync.Mutex

    go func() {
        lock1.Lock()
        time.Sleep(1 * time.Second)
        lock2.Lock()
        lock1.Unlock()
        lock2.Unlock()
    }()

    go func() {
        lock2.Lock()
        time.Sleep(2 * time.Second)
        lock1.Lock()
        lock2.Unlock()
        lock1.Unlock()
    }()

    time.Sleep(5 * time.Second)
}
