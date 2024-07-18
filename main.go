package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web/Web"
)

func main() {

	// Kill server when arguments other than the program name are added
	if len(os.Args) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.HandleFunc("/", Web.SubmitFormHandler) // Serve home page
	http.HandleFunc("/ascii-art", Web.SubmitFormHandler) // Serve /ascii-art at POST requests
	log.Fatal(http.ListenAndServe(":8080", nil)) // Defines host port
}
