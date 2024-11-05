package json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Ivan",
		MiddleName: "Adi",
		LastName: "Saputra",
	}

	encoder.Encode(customer)
}