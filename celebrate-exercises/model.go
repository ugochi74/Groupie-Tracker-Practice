// package main
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )




// type Artist struct {
// 	ID int  `json:"id"`
// 	Name string  `json:"name"`
// 	Members []string  `json:"members"`
// 	CreationDate int  `json:"creationDate"`
// 	FirstAlbum string   `json:"firstAlbum"`
// 	ConcertDates string  `json:"concertDates"`
// 	Locations string  `json:"locations"`
// 	Relations string   `json:"relations"`
// 	Image string      `json:"image"`
// }
// func main() {
// 	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	content, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	var artists []Artist
// 	err = json.Unmarshal(content, &artists)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}

// 		fmt.Printf("Successfully parsed %d artists!\n", len(artists))

// 	fmt.Printf("%+v\n", artists)
	
// }


package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

// Artist matches the API structure perfectly
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

// HTML template layout to display cards visually in the browser
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Groupie Trackers</title>
    <style>
        body { font-family: Arial, sans-serif; background-color: #f4f4f9; margin: 0; padding: 20px; }
        h1 { text-align: center; color: #333; }
        .container { display: flex; flex-wrap: wrap; gap: 20px; justify-content: center; }
        .card { background: white; border-radius: 10px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); width: 300px; padding: 20px; text-align: center; }
        .card img { max-width: 100%; border-radius: 8px; height: auto; }
        .card h2 { color: #111; margin: 10px 0; }
        .info { text-align: left; font-size: 14px; color: #555; }
        ul { padding-left: 20px; margin: 5px 0; }
    </style>
</head>
<body>
    <h1>🎵 Groupie Trackers Artists 🎵</h1>
    <div class="container">
        {{range .}}
        <div class="card">
            <img src="{{.Image}}" alt="{{.Name}}">
            <h2>{{.Name}}</h2>
            <div class="info">
                <p><strong>First Album:</strong> {{.FirstAlbum}}</p>
                <p><strong>Created In:</strong> {{.CreationDate}}</p>
                <p><strong>Members:</strong></p>
                <ul>
                    {{range .Members}}
                    <li>{{.}}</li>
                    {{end}}
                </ul>
            </div>
        </div>
        {{end}}
    </div>
</body>
</html>
`

func fetchArtists() ([]Artist, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	err = json.Unmarshal(content, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	artists, err := fetchArtists()
	if err != nil {
		http.Error(w, "Failed to load data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse and execute HTML templates to output visual page structures
	tmpl, err := template.New("webpage").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Template configuration error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, artists)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("🚀 Server running smoothly! Open your browser at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
