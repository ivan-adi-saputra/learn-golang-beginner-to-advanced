package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello ", value.getName())
}

// implementasi menggunakan struct
type Person struct {
	name string
}

func (person Person) getName() string {
	return person.name
}

type Animal struct {
	name string
}

func (animal Animal) getName() string {
	return animal.name
}

// interface kosong atau any
func Ups() interface{} {
	// return 1
	return "Hello"
}

func Ups1() any {
	// return 1
	return "Any"
}

func main() {
	person := Person{name: "Ivan"}
	sayHello(person)

	animal := Animal{name: "Kucing"}
	sayHello(animal)

	// interface kosong atau any
	ups := Ups()
	fmt.Println(ups)

	ups1 := Ups1()
	fmt.Println(ups1)
}