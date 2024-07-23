package lib

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"strings"
)

// Prints new lines and calls HandleCharacters() on each word
func HandleWords(slices []string, words []string) string {
	var output string
	countSpaces := 0

	for _, word := range words {
		if word == "" {
			countSpaces++
			if countSpaces < len(words) {
				output += "\n"
			}
		} else {
			output = HandleCharacters(output, word, slices)
		}
	}
	return output
}

// Prints the ASCII ART for each character in each word and updates the output string
func HandleCharacters(output, word string, slices []string) string {
	var startIndex int

	for i := 0; i < 8; i++ {
		for _, ch := range word {
			startIndex = int(ch-32)*9 + 1
			output += slices[startIndex+i]
		}
		output += "\n"
	}
	return output
}

// Chaeck if input contains non-printable ascii characters with escape sequences (except '\n')
func EscapeSequence(input string) bool {
	allowedEscapes := []string{"\\a", "\\b", "\\t", "\\v", "\\f"}
	for _, unprint := range allowedEscapes {
		if strings.Contains(input, unprint) {
			return true
		}
	}
	return false
}

// Check if input contains characters outside of the banner file (except '\n')
func IsPrintable(input string) bool {
	var status bool
	for _, v := range input {
		if (v >= ' ' && v <= '~') || v == '\n' || v == '\r' {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
}

// Check if the banner file is valid
func ValidFile(file string) bool {
	fileNames := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
	var status bool

	for _, flName := range fileNames {
		if file == flName {
			status = true
		}
	}
	return status
}

func fileIntegrity(file []byte) bool {
	var status bool

	//Generate hash for banner file
	hash := sha256.New()
	hash.Write(file)
	curr := fmt.Sprintf("%x", hash.Sum(nil))

	//Hashes of various banner files when unadultarated
	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	checksums := []string{standard, shadow, thinkertoy}

	//Compare whether current hash matches with any of the checksums above
	for _, hsh := range checksums {
		if curr == hsh {
			status = true
		}
	}

	return status
}

func checkOS() []string {
	var sep []string
	os := runtime.GOOS // Identify operating system

	// Set separators according to operating system
	if os == "windows" {
		sep = []string{"\n", "\r\n"}
	} else {
		sep = []string{"\r\n", "\n"}
	}

	return sep
}
