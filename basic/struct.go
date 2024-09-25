package main

import "fmt"

// struct -> digunalan untuk menggabungkan nol atau type data lain dalam satu kesatuan (template)
type Customer struct {
	Name, Address string
	Age            int
}

// struct method / func dalam struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, " My name is ", customer.Name)
}

func main() {
	ivan := Customer{
		Name:      "Ivan",
        Address:   "Jln. Kemiri 10",
        Age:        19,
	}
	adi := Customer{"Adi", "Indonesia", 18}
	fmt.Println("Customer Ivan: ", ivan)
	fmt.Println("\nCustomer Adi: ", adi)

	ivan.sayHello("saputra")
}