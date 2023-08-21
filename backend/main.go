// backend/main.go
package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/backend/handlers"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/404", handle404)
	http.HandleFunc("/500", handle500)

	// Serve static files (CSS, images, etc.) from the frontend directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))


	artists, err := handlers.GetArtists()
	
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    for _, artist := range artists {
        fmt.Printf("Artist ID: %d, Name: %s\n", artist.ID, artist.Name)
    }

	// for i, p := range response.Persons {
	// 		fmt.Println("")
	// }

	port := "3000"
	println("Server listening on port http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = filepath.Join("frontend", tmpl+".html")
	layout := filepath.Join("frontend", "layout.html")
	t, err := template.ParseFiles(layout, tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	data := struct{}{} // Data for rendering, if needed
	renderTemplate(w, "index", data)
}

func handle404(w http.ResponseWriter, r *http.Request) {
	data := struct{}{} // Data for rendering, if needed
	renderTemplate(w, "404", data)
}

func handle500(w http.ResponseWriter, r *http.Request) {
	data := struct{}{} // Data for rendering, if needed
	renderTemplate(w, "500", data)
}

type Response struct {
    Page       int `json:"page"`
    PerPage    int `json:"per_page"`
    Total      int `json:"total"`
    TotalPages int `json:"total_pages"`
    Data       []struct {
        ID        int    	`json:"id"`
        Name     string 	`json:"name"`
        Members string 		`json:"members"`
        CreationDate  string `json:"creationDate"`
        Avatar    string `json:"avatar"`
    } `json:"data"`
    Support struct {
        URL  string `json:"url"`
        Text string `json:"text"`
    } `json:"support"`
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

