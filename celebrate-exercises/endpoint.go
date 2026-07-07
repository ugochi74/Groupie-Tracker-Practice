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


type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// These endpoints return an object whose "index" field
// contains the actual slice.

type LocationResponse struct {
	Index []Location `json:"index"`
}

type DateResponse struct {
	Index []Date `json:"index"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}

// --------------------
// Package-level data
// --------------------

var (
	artists   []Artist
	locations []Location
	dates     []Date
	relations []Relation
)

// learning nerver ends
func fetchArtists()  error {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/artists")
	    if err != nil {
		return err
	
		}
		
	defer res.Body.Close()


	//var response ArtistResponse

	decoder := json.NewDecoder(res.Body)

	//var artists []Artist
	err = decoder.Decode(&artists)
	if err != nil {
		return err

   }
   //artists = response.Index
   return  nil
}


func fetchLocations()  error {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/locations")
	    if err != nil {
		return err
	
		}
		
	defer res.Body.Close()
	var response LocationResponse
	decoder := json.NewDecoder(res.Body)

	//var locations []Location
	err = decoder.Decode(&response)
	if err != nil {
		return err

   }
   locations = response.Index
   return  nil
}


func fetchRelations()  error {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/relation")
	    if err != nil {
		return err
	
		}
		
	defer res.Body.Close()
	var response RelationResponse

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&response)
	if err != nil {
		return err

   }
   relations = response.Index
   return nil
}


func fetchDates() error {
    res, err := http.Get ("https://groupietrackers.herokuapp.com/api/dates")
	    if err != nil {
		return err
	
		}
		
	defer res.Body.Close()
	var response DateResponse
	decoder := json.NewDecoder(res.Body)

	//var dates []Date
	err = decoder.Decode(&response)
	if err != nil {
		return err
	}
     dates = response.Index
    return  nil
  
}



func fetchAll() error {
	if err := fetchArtists(); err != nil{
	
		return err
	}

 if  err := fetchLocations(); err != nil{
	
		return err
 }

	 if err := fetchRelations(); err != nil{
		return err
	}

	 if err := fetchDates(); err != nil { 
	
		return err
	
}
 return nil
}

	func main() {
		 err := fetchAll()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		 for _, artist := range artists {
 	fmt.Println(artist.Name)
	fmt.Println(artist.Members)
	//fmt.Println(artist.Relations)

 }

	fmt.Println("artists", len(artists))
	fmt.Println("locations:", len(locations))
	fmt.Println("relations:", len(relations))
	fmt.Println("dates:", len(dates))
	fmt.Println()



for _, location := range locations {
    fmt.Println("Artist ID:", location.ID)


for _, place := range location.Locations{
	fmt.Println(place)
}
   fmt.Println()
}
// for _, relation := range relations {
//     fmt.Println("Artist ID:", relation.ID)

//     for Location, Dates := range relation. DatesLocations {
//         fmt.Println("Location:", Location)
// 		fmt.Println("Dates:", Dates)
//     }

    fmt.Println()
}
	
