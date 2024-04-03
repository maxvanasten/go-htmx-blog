package router

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type BlogPost struct {
	Path string
	Name string
}

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

    // TODO: Get directories in 'pages' directory, perform below functionality for every directory.

	// Get list of files in 'pages' directory
	files, err := os.ReadDir("html/pages")
	if err != nil {
		log.Fatal(err)
	}

	BlogPosts := make(map[string]BlogPost)
	for _, file := range files {
		path := strings.Replace(file.Name(), ".html", "", -1)

        // Add blog post to BlogPosts for the navigator
		BlogPosts[path] = BlogPost{
			path,
			strings.Replace(path, "_", " ", -1),
		}

        // Add routes for every blog post
		Routes["/pages/"+path] = func(w http.ResponseWriter, r *http.Request) {
			ExecuteTemplate(w, "html/pages/"+file.Name(), nil)
		}
	}

	// Page navigator
	Routes["/partials/navigator"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/partials/navigator.html", BlogPosts)
	}

	return Routes
}

func ExecuteTemplate(w http.ResponseWriter, path string, data any) {
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, data)
}
