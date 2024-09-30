package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// membuat container list / double link list
	var data list.List = *list.New()
	// menambahkan data ke dalam list
	data.PushFront("Ivan")
    data.PushFront("Adi")
    data.PushFront("Saputra")

    // mencetak data list
    fmt.Println("Data list:")
    for e := data.Front(); e!= nil; e = e.Next() {
        fmt.Print(e.Value, " ")
    }
    fmt.Println()

    // menghapus data pada posisi tertentu
    data.Remove(data.Front())
    fmt.Println("Data list setelah menghapus data pertama:")
    for e := data.Front(); e!= nil; e = e.Next() {
        fmt.Print(e.Value, " ")
    }
    fmt.Println()

    // mencari data di dalam list
    found := data.Front()
    for found!= nil {
		if found.Value == "Adi" {
            break
        }
		found = found.Next()
	}

	// container/ring adalah implementasi struktur data circular list
	// circular ring adalah struktur data ring dimana akhir element akan dikembalikan ke awal (HEAD)
	var ringData *ring.Ring = ring.New(5)
    
	for i := 0; i < ringData.Len(); i++ {
		ringData.Value = "Data " + strconv.Itoa(i)  // Menggunakan Itoa untuk mengonversi integer ke string
        ringData = ringData.Next()
	}

	ringData.Do(func(value any) {
		fmt.Println(value)
	})
}