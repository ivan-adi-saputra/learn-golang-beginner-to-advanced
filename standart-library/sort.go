package main

import (
	"fmt"
	"sort"
)

// sort digunakan untuk proses pengurutan data

type User struct {
	Name string
	Age  int
}

type sliceUser []User 

// untuk menghitung panjang
func (s sliceUser) Len() int {
	return len(s)
}

// untuk mengganti posisi
func (s sliceUser) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// untuk mengkomparasi
    // return true jika i lebih kecil dibandingkan j, false jika i lebih besar dibandingkan j, dan true jika i sama dengan j
    // jika age sama, maka akan diurutkan berdasarkan nama
    // jika age dan nama sama, maka akan diurutkan berdasarkan umur
    // jika age, nama, dan umur sama, maka akan diurutkan berdasarkan nama kemudian umur, dan seterusnya  (dari kiri ke kanan)
    // jika age, nama, dan umur sama, maka akan diurutkan berdasarkan nama kemudian umur, dan seterusnya  (dari kiri ke kanan)
    // jika age, n
func (s sliceUser) Less(i, j int) bool {
    return s[i].Age < s[j].Age
}

func main() {
	users := []User{
        {Name: "ivan", Age: 25},
        {Name: "adi", Age: 23},
        {Name: "saputra", Age: 25},
    }

    // sorting
    sort.Sort(sliceUser(users))

	fmt.Println(users)
}