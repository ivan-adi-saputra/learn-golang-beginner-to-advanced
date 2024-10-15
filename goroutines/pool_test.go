package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool
// Pool adalah implementasi design pattern bernama object pool pattern.
// Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
// Implementasi Pool di Go-Lang ini sudah aman dari problem race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// jika pool nil maka akan mengembalikan New
        New: func() interface{} {
            return "New"
        },
    }

	pool.Put("ivan")
	pool.Put("adi")
	pool.Put("saputra")

	for i := 0; i < 10; i++ {
		go func ()  {
			data := pool.Get() // menambil data pool (otomatis data yang diambil kosong)
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // mengembalikan data ke pool
		}()
	}

	time.Sleep(10 * time.Second)
}