package main

import (
	"log"
	"net/http"
	"web/Web"
)

func main() {
	// // Saves the input from the user
	// inputs := os.Args[1:]

	// // Exits the program if the arguments passed are greater than 5
	// if len(inputs) == 0 || len(inputs) > 5 {
	// 	Lib.PrintError()
	// }

	// // Check what user input contains and returns required variables
	// color1, color2, reset, mainString, subString, fileName, outputFile := Lib.CheckInput(inputs)

	http.HandleFunc("/", Web.FormHandler)
	http.HandleFunc("/submit-ascii-art", Web.SubmitFormHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8000", nil))

	// // Call the AsciiArt function to handle input
	// output := Lib.AsciiArt(color1, color2, reset, mainString, subString, fileName)

	// //Write ascii art to outputFile
	// if len(outputFile) > 0 {
	// 	os.WriteFile(outputFile, []byte(output), 0666)
	// } else {
	// 	fmt.Print(output)
	// }
}
