package main

import "fmt"

func main() {
	name := [...]string{"ivan", "adi", "saputra", "fullstack", "developer"}
	fmt.Println(name)

	slice1 := name[1:3]
	lenSlice1 := len(slice1) // 4
	capSlice1 := cap(slice1) // 2
	fmt.Println(capSlice1)
	fmt.Println(lenSlice1)
	fmt.Println(slice1)

	slice2 := name[:3]
	fmt.Println(slice2)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "sabtu lama"
	daysSlice1[1] = "minggu lama"
	fmt.Println(days) // data akan ke update jadi sabtu lama dan minggu lama
	fmt.Println(daysSlice1) // sabtu lama minggu lama

	// karena di array tidak memenuhi maka akan membuat array baru dan array days tidak ke update
	daysSlice2 := append(daysSlice1, "libur baru")
	fmt.Println(daysSlice2)

	// membuat slice baru
	// alasan membuat slice dari awal biar tidak membuat array baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "ivan"
	newSlice[1] = "adi"
	// newSlice[2] = "saputra" -> error karena panjangnya cuma 2, harusnya menggunakan append
	newSlice1 := append(newSlice, "saputra")
	fmt.Println(newSlice1) // [ivan adi saputra]
	fmt.Println(newSlice) // ivan adi

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// yang perlu diperhatikan dalam slice
	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// pilih array atau slice? 
	// kalau pakai golang rata" pakai slice daripada array
}

// membuat slice
// array[low:high] -> slice mulai dari index low sampai sebelum index high
// array[low:] -> slice mulai dari index low sampai terakhir
// array[:high] -> slice dari index 0 sampai sebelum high
// array[:] -> slice dari index 0 sampai akhir

// function slice
// len(slice) -> untuk mendapatkan panjang slice
// cap(slice) -> mendapatkan capasitas slice, mulai dari pointer
// append(slice, data) -> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas penuh, maka akan membuat array baru
// make([]TypeData, length, capacity) -> membuat slice baru
// copy(destination, source) -> menyalin slice dari source ke destination

