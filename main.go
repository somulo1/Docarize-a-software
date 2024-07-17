package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web/Web"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	// http.HandleFunc("/", Web.FormHandler)
	http.HandleFunc("/", Web.SubmitFormHandler)
	http.HandleFunc("/generate", Web.SubmitFormHandler)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
