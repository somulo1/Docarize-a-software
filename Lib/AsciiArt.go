package Lib

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input, bnStyle string) (string, string) {
	var err string
	// Exits program if input is empty
	if input == "" {
		return "", err
	}

	// Exits program if input contains characters outside of the banner file (except '\n')
	if !IsPrintable(input) {
		err = `Input should only contain PRINTABLE ASCII characters or '\n'`
		return "", err
	}

	// Exits program if input contains non-printable ascii characters with escape sequences (except '\n')
	if EscapeSequence(input) {
		err = `Input should only contain PRINTABLE ASCII characters or '\n'`
		return "", err
	}

	// Split input into printable lines at the '\n' character
	words := strings.Split(input, "\r\n")

	// Set and check if banner file is valid
	if !ValidFile(bnStyle) {
		err = "Incorrect file name, program only accepts thinkertoy.txt, standard.txt or shadow.txt"
		return "", err
	}

	// Read banner file
	content, fileErr := os.ReadFile("banner-files/" + bnStyle)
	if fileErr != nil {
		err = fmt.Sprintf("Error reading %s\n", bnStyle)
		return "", err
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
	return HandleWords(slices, words), err
}
