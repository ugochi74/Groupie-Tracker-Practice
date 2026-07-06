package main


import (
	"fmt"
	"net/http"
	"io"
)
// learning nerver ends
func main() {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/artists")
	    if err != nil {
        fmt.Println("error occured:", err)
		return
	
		}
		
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}