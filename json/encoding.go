package main

import (
	"encoding/json"
	"fmt"
)

type Artist struct {
	Name string
	Year int
}

func main() {
	data := []byte(`{"name": "Queen", "year": 1970}`)
// empty Artist struct
	var a Artist
	// fill the struct with values from the json
	json.Unmarshal(data, &a)
// print the struct
	fmt.Println(a)
}