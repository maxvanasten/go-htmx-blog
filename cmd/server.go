package main

import (
    "fmt"
    "net/http"
    "github.com/maxvanasten/go-htmx-blog/pkg/router"
)

func main() {
    // Add all paths to http server
    for path, handlerFunc := range router.GetRoutes() {
        fmt.Println("Adding path", path, "to server")
        http.HandleFunc(path, handlerFunc)
    } 

    // Start http server
    fmt.Printf("\nhttp://localhost%s", ":3000\n")
    http.ListenAndServe(":3000", nil)
}
