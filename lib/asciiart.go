package lib

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input, bnStyle string) (string, string) {
	var err string

	// Exits program if input contains characters outside of the banner file (except '\n')
	if !IsPrintable(input) {
		err = `Input should only contain PRINTABLE ASCII characters`
		return "", err
	}

	// Exits program if input contains non-printable ascii characters with escape sequences (except '\n')
	if EscapeSequence(input) {
		err = `Input should only contain PRINTABLE ASCII characters`
		return "", err
	}

	// Split input into printable lines at the '\n' character
	words := strings.Split(input, "\r\n")

	// Read banner file
	content, fileErr := os.ReadFile("banner-files/" + bnStyle)
	if fileErr != nil {
		err = fmt.Sprintf("Error reading %s\n", bnStyle)
		return "", err
	}

	// Check contents of file for adulteration
	if !fileIntegrity(content) {
		err += fmt.Sprintf("%s is modified or corrupted", bnStyle)
		return "", err
	}

	// Split and store the file contents line-by-line in a slice string depending on the banner file
	var slices []string

	// Lines in thinkertoy.txt banner file are separated by "\r\n" or "\n" depending on operating system
	if bnStyle == "thinkertoy.txt" {
		slices = strings.Split(string(content), "\r\n")
	} else {
		slices = strings.Split(string(content), "\n")
	}

	// Print ASCII ART and return output string for testing
	return HandleWords(slices, words), err
}
