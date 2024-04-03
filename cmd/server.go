package main

import (
    "fmt"
    "net/http"
    "github.com/maxvanasten/go-htmx-blog/pkg/router"
)

func main() {
    for path, handlerFunc := range router.GetRoutes() {
        fmt.Println("Adding path", path, "to server")
        http.HandleFunc(path, handlerFunc)
    } 

    fmt.Printf("\nhttp://localhost%s", ":3000")
    http.ListenAndServe(":3000", nil)
}
