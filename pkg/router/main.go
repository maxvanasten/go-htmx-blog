package router

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetRoutes() map[string]http.HandlerFunc {
	Routes := make(map[string]http.HandlerFunc)

	// Main page
	Routes["/"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/index.html", nil)
	}

	// Partials
	Routes["/partials/footer"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/partials/footer.html", nil)
	}

	// Pages
	files, err := os.ReadDir("html/pages")
	if err != nil {
		log.Fatal(err)
	}

	type PFile struct {
		Path string
		Name string
	}

	pfiles := make(map[string]PFile)
	for _, file := range files {
		path := strings.Replace(file.Name(), ".html", "", -1)

		pfiles[path] = PFile{
			path,
			strings.Replace(path, "_", " ", -1),
		}

		Routes["/pages/"+path] = func(w http.ResponseWriter, r *http.Request) {
			ExecuteTemplate(w, "html/pages/"+file.Name(), nil)
		}
	}

	// Page navigator

	Routes["/partials/navigator"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/partials/navigator.html", pfiles)
	}

	return Routes
}

func ExecuteTemplate(w http.ResponseWriter, path string, data any) {
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, data)
}
