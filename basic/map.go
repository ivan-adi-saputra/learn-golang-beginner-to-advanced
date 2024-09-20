package main

import "fmt"

func main() {
	person := map[string]string{
		"name":  "Ivan",
        "age":   "25",
        "city":  "Bandung",
        "email": "ivan@gmail.com",
	}
	fmt.Println(person["name"])
	fmt.Println(person)
}

// function map
// make(map[typeDataKey]TypeData) -> membuat mao baru
// delete(map, key) -> menghapus map berdasarkan key