package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Eko",
		MiddleName: "Kurniawan",
		LastName: "Khannedy",
	}

	encoder.Encode(customer)
}

