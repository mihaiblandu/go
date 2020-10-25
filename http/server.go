package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "Thetime is now " + time.Now().String()
	fmt.Fprintf(w, response)
}
func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}
func serveError(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
	fmt.Println("There is no way to do this.")
}
func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/dinamic", serveDynamic)
	http.HandleFunc("/error", serveError)
	http.ListenAndServe(Port, nil)
}
