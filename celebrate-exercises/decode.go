package main


import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertDates string   `json:"concertDates"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
	Image        string   `json:"image"`
}

// learning nerver ends
func main() {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/artists")
	    if err != nil {
        fmt.Println("error occured:", err)
		return
	
		}
		
	defer res.Body.Close()
	decoder:= json.NewDecoder(res.Body)

	var artists []Artist
	err = decoder.Decode(&artists)
	if err != nil {
		fmt.Println("Error", err)
		return
 	}
	for _, char := range artists {
		fmt.Println("ID:", char.ID)
         fmt.Println("Name:", char.Name)
		 fmt.Println("CreationDate:", char.CreationDate)
		 fmt.Println("FirstAlbum:", char.FirstAlbum)
		 fmt.Println("Members:", char.Members)
		 fmt.Println("Relations:", char.Relations)
		 fmt.Println("Locations:", char.Locations)
		 fmt.Println("ConcertDates:", char.ConcertDates)
		 fmt.Println("Image:", char.Image)
		

		 fmt.Println()
	}
	fmt.Println("Total artists:", len(artists))
}