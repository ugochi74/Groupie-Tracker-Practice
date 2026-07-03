package main
import (

	"fmt"
	"net/http"
	"io"
)
// container/struct that holds the response, which you then have to do something further with to get the actual text out?
func main() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(body))
}