package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// Map
// Go-Lang memiliki sebuah struct beranama sync.Map
// Map ini mirip Go-Lang map, namun yang membedakan, Map ini aman untuk menggunaan concurrent menggunakan goroutine
// Ada beberapa function yang bisa kita gunakan di Map :
// Store(key, value) untuk menyimpan data ke Map
// Load(key) untuk mengambil data dari Map menggunakan key
// Delete(key) untuk menghapus data di Map menggunakan key
// Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup)  {
	group.Done()

	group.Add(1)

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	
	for i := 0; i < 100; i++ {
        go AddToMap(data, i, group)
    }

	data.Range(func (key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}