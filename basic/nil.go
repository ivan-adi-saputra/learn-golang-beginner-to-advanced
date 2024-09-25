package main

import "fmt"
// nil mengembalikan data kosong
// nil hanya bisa digunakan ketika data yang dikembalikan berupa function, interface, map, slice, pointer, channel
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name := newMap("Ivan")

	if name == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(name["name"])
	}
}