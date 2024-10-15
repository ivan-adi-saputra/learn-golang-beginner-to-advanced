package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

// GOMAXPROCS
// Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
// Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
// Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
// Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
// Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()

func TestGetGomacprocs(t *testing.T)  {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1) // tidak mengubah thread, karena kalau di atas 0 akan mengubah thread
	fmt.Println("Total Thread ", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines ", totalGoroutines)
}

func TestChangeThread(t *testing.T)  {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	runtime.GOMAXPROCS(20) // thread akan menjadi 20
	totalThread := runtime.GOMAXPROCS(-1) 
	fmt.Println("Total Thread ", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines ", totalGoroutines)
}

// Peringatan
// Menambah jumlah thread tidak berarti membuat aplikasi kita menjadi lebih cepat
// Karena pada saat yang sama, 1 CPU hanya akan menjalankan  1 goroutine dengan 1 thread
// Oleh karena ini, jika ingin menambah throughput aplikasi, disarankan lakukan vertical scaling (dengan menambah jumlah CPU) atau horizontal scaling (menambah node baru)
