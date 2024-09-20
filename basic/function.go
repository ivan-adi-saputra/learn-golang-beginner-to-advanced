package main

import "fmt"

func sayHello(name string) string {
	return "Hello" + name
}

// func return multiple / mengembalikan lebih dari 1 data
func getFullName() (string, string) {
	return "Ivan", "Saputra"
}

// named return values
// jika semua type data sama maka bisa deklarasi di akhir
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "ivan"
	middleName = "adi"
	lastName = "saputra"
	return firstName, middleName, lastName
}

// variadic -> data bisa menerima lebih dari 1 input semacam array
func sumAll(numbers ...int) int {
	total:= 0

	for _, number := range numbers {
        total += number
    }
	return total
}

func getGoodBye(name string) string {
	return "Goodbye" + name
}

// function as parameter / function yang dijadikan parameter function
type Filter func(string) string
func sayHellos(name string, filter Filter) {
	filteredNamed := filter(name)
	fmt.Println(filteredNamed)
}
func filtered(name string) string {
	if name == "ivan" {
		return "Halo"
	} else {
		return "..."
	}
}

// anonymous func
type Blacklist func(string) bool
func anonymous(name string, blacklist Blacklist) {
	if(blacklist(name)) {
		fmt.Println("Blacklisted")
	} else {
		fmt.Println("Welcome", name)
	}
}

// defer -> digunakan ketika ingin memanggil di akhir func tsb
func loging() {
	fmt.Println("Selesai...")

	// recover -> digunakan untuk menangkap error pada panic
	// recover ditaruh di loging karena di defer -> func yang akan dijalankan di akhir func tsb walaupun terjasi error
	message := recover()
	fmt.Println(message)
}
func runApps() {
	defer loging()
    fmt.Println("Mulai...")
    // ketika fungsi runApps di panggil, fungsi loging akan dipanggil di akhir fungsi runApps
}

// panic -> digunakan untuk menghentikan program
func runPanic(error bool) {
	defer loging()
	if error {
        panic("Terjadi error")
    } else {
        fmt.Println("Berhasil...")
    }
}

func main() {
	result := sayHello("ivan")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	// jika ingin menghiraukan salah satu valu tinggal kasih _
	// _, lastName := getFullName()

	a, _, c := getCompleteName()
	fmt.Println("First Name:", a)
	fmt.Println("Last Name:", c)

	// variadic function
	resultSum := sumAll(1, 2, 3, 4, 5)
    fmt.Println("Sum:", resultSum)
	// ketika data yang dikirimkan slice
	numbers := []int{1,2,3,4,5,6,7,7}
	sliceNumber := sumAll(numbers...)
	fmt.Println("Sum using slice:", sliceNumber)

	// menjadikan function sebagai value
	// ketika ingin menjadikan values harus tidak menggunakan () karena ketika menggunaakannya maka akan di asumsikan memanggil func 
	hello := getGoodBye
	fmt.Println(hello("ivan"))

	// function as parameter / function yang dijadikan parameter function
	sayHellos("Hi", filtered)

	// anonymous func
	anonymous("ivan", func(name string) bool {
		return name != "ivan"
	})

	// defer -> digunakan ketika ingin memanggil di akhir func tsb
	runApps()
    // ketika fungsi runApps di panggil, fungsi loging akan dipanggil di akhir fungsi runApps
    // ketika menggunakan defer, semua fungsi yang di defer akan dipanggil di akhir fungsi main

	// panis
	runPanic(true) // ketika error, program akan berhenti
}