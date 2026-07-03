package main

import (
	"fmt"
	"encoding/json"
)


type Artist struct {
	Name       string `json:"name"` // json tag
	FirstAlbum string `json:"first_album"` // json tag
}

func main() {
	data := []byte(`{"name": "Queen", "first_album": "1973"}`)
	var a Artist
	// marshal convert Go value to JSON
	json.Unmarshal(data, &a) // unmarshal convert JSON to Go value.
	fmt.Println(a)
}