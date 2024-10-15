package goroutines

import (
	"fmt"
	"testing"
	"time"
)

// time.Ticker
// Ticker adalah representasi kejadian yang berulang
// Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
// Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
// Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()

func TestTicker(t *testing.T)  {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for data := range ticker.C {
		fmt.Println(data)
	}
}

// time.Tick
// Kadang kita tidak butuh data Ticker nya, kita hanya butuh channel nya saja
// Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja

func TestTick(t *testing.T)  {
	ticker := time.Tick(1 * time.Second)

    go func() {
        time.Sleep(5 * time.Second)
    }()

    for data := range ticker {
        fmt.Println(data)
    }
}