package main

import (
	"fmt"
	"errors"
)

// error interface
// golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

// custom error
type validationError struct {
	Message string
}
func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}
func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Errors"}
	}
	if id != "ivan" {
		return &notFoundError{"Not Found Errors"}
	}

	return nil
}

func main() {
	result, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil ", result)
	} else {
		fmt.Println("Errors", err.Error())
	}
	
	// custom error
	error_custom := saveData("ivan", nil)
	if error_custom != nil {
		switch result := error_custom.(type) {
		case *validationError:
			fmt.Println("Validation Error", result.Error())
		case *notFoundError:
			fmt.Println("NotFound", result.Error())
		default: 
			fmt.Println("Success")
		}
	} else {
		fmt.Println("Success")
	}
}