package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer
// Timer adalah representasi satu kejadian
// Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
// Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)

func TestTimer(t *testing.T)  {
	timer := time.NewTimer(time.Second * 5)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Printf("Timer expired at %v\n", time)
}

// time.After()
// Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
// Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)

func TestAfter(t *testing.T)  {
	timer := time.After(time.Second * 5)
    fmt.Println(time.Now())

    time := <-timer
    fmt.Printf("Timer expired at %v\n", time)
}

// time.AfterFunc()
// Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
// Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
// Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadiannya

func TestAfterFunc(t *testing.T)  {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(time.Second * 5, func() {
        fmt.Println("Timer expired")
		group.Done()
    })

	group.Wait()

    time.Sleep(time.Second * 6)
}