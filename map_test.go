package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"p0001", "name":"Apple Mac Book Pro", "price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id" : "P0001",
		"name" : "Apple Mac Book Pro",
		"price" : 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
