package Lib

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input, bnStyle string) string {
	// Exits program if input is empty
	if input == "" {
		return ""
	}

	// Exits program if input contains characters outside of the banner file (except '\n')
	if !IsPrintable(input) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}

	// Exits program if input contains non-printable ascii characters with escape sequences (except '\n')
	if EscapeSequence(input) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}
	// // Make new line characters consistent
	// input = strings.ReplaceAll(input, "\\n", "\n")

	// Split input into printable lines at the '\n' character
	words := strings.Split(input, "\r\n")

	fmt.Printf("%q\n", words)

	// Set and check if banner file is valid
	if !ValidFile(bnStyle) {
		fmt.Println("Incorrect file name, program only accepts thinkertoy.txt, standard.txt or shadow.txt")
		return ""
	}

	// Read banner file
	content, err := os.ReadFile("banner-files/"+ bnStyle)
	if err != nil {
		fmt.Printf("Error reading %s\n", bnStyle)
		return ""
	}

	// Split and store the file contents line-by-line in a slice string depending on the banner file
	var slices []string
	if bnStyle == "thinkertoy.txt" {
		// Lines in thinkertoy.txt banner file are separated by "\r\n"
		slices = strings.Split(string(content), "\r\n")
	} else {
		// Lines in standard.txt and shadow.txt banner files are separated by "\n"
		slices = strings.Split(string(content), "\n") 
	}
	
	// Print ASCII ART and return output string for testing
	return HandleWords(slices, words)
}
