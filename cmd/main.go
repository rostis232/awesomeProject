package main

import (
	"awesomeProject/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main application.
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
