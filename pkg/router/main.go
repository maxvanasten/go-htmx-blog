package router

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type BlogPost struct {
	Path     string
	Name     string
	Category string
}

func GetRoutes() map[string]http.HandlerFunc {
	// Initialize 'Routes' map
	Routes := make(map[string]http.HandlerFunc)

	// Main page
	Routes["/"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/index.html", nil)
	}

	// Posts
	// Get list of directories in 'posts' directory
	directories, err := os.ReadDir("html/posts")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize list of categories
	categories := []string{}
	// Iterate over 'directories'
	for _, category := range directories {
		// Look for directories
		if category.IsDir() {
			// Add category name to 'categories' list
			category_name := category.Name()
			categories = append(categories, category_name)
			// Get all files in this category
			files, err := os.ReadDir("html/posts/" + category_name)
			if err != nil {
				log.Fatal(err)
			}
			// Initialize 'BlogPosts' list
			BlogPosts := make(map[string]BlogPost)
			// Iterate over 'files' in this category
			for _, file := range files {
				// Remove '.html' from filename
				path := strings.Replace(file.Name(), ".html", "", -1)

				// Add blog post to BlogPosts for the navigator
				BlogPosts[path] = BlogPost{
					path,
					strings.Replace(path, "_", " ", -1),
					category_name,
				}

				// Add routes for every blog post
				Routes["/posts/"+category_name+"/"+path] = func(w http.ResponseWriter, r *http.Request) {
					ExecuteTemplate(w, "html/posts/"+category_name+"/"+file.Name(), nil)
				}
			}
			// Add route for every category, returns a list of blog posts
			Routes["/partials/posts_navigator/"+category_name] = func(w http.ResponseWriter, r *http.Request) {
				ExecuteTemplate(w, "html/partials/posts_navigator.html", BlogPosts)
			}
		}
	}

	// Category navigator
	Routes["/partials/category_navigator"] = func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, "html/partials/category_navigator.html", categories)
	}

	return Routes
}

func ExecuteTemplate(w http.ResponseWriter, path string, data any) {
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(w, data)
}
