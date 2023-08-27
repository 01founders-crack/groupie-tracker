// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"html/template"
	"net/http"
	"path/filepath"
)

func main() {
	// Serve static files (CSS, images, etc.) from the frontend directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))

	http.HandleFunc("/", handleNotFound)
	http.HandleFunc("/500", handle500)

	_, err := handlers.GetArtists()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	port := "3000"
	println("Server listening on port http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

/*
func printRelatedDatesLocations(url string) {
	relations, err := handlers.GetRelations(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for location, relatedDatesLocations := range relations.DatesLocations {
		fmt.Printf("location: %s\n", location)
		fmt.Println("Related DatesLocations:")
		for _, relatedDateL := range relatedDatesLocations {
			fmt.Println(relatedDateL)
		}
	}
}
*/

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

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Serve the index.html page for the root path
		combinedData, err := handlers.GetArtistsWithRelations()
		if err != nil {
			fmt.Println("Error:", err) // Print the error for debugging
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}

		renderTemplate(w, "index", combinedData)
	} else {
		// Serve the 404.html page for other not-found routes
		data := struct{}{} // Data for rendering, if needed
		renderTemplate(w, "404", data)
	}
}

func handle500(w http.ResponseWriter, r *http.Request) {
	data := struct{}{} // Data for rendering, if needed
	renderTemplate(w, "500", data)
}
