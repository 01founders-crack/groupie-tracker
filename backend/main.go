// backend/main.go
package main

import (
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

	port := "8080"
	println("Server listening on port " + port)
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
