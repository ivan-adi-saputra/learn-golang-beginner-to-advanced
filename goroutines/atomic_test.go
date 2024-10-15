package goroutines

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// Atomic
// Go-Lang memiliki package yang bernama sync/atomic
// Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
// Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya bisa digunakan menggunakan Atomic package
// Ada banyak sekali function di atomic package, kita bisa eksplore sendiri di halaman dokumentasinya
// https://golang.org/pkg/sync/atomic/

func TestAtomic(t *testing.T) {
	var counter int32

    // Increment counter atomically
    go func() {
        for i := 0; i < 10000; i++ {
            atomic.AddInt32(&counter, 1)
        }
    }()

    // Decrement counter atomically
    go func() {
        for i := 0; i < 10000; i++ {
            atomic.AddInt32(&counter, -1)
        }
    }()

    time.Sleep(time.Second)

	fmt.Println(counter)

    // Counter should be 0
    if counter != 0 {
        t.Error("Counter is not 0")
    }
}