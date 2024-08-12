package lib

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestHandleWords(t *testing.T) {
	file, _ := os.ReadFile("../banner-files/standard.txt") // Read standard.txt
	content := strings.Split(string(file), "\n")           // Split file line-by-line

	subject := []string{"Hello"}
	got := HandleWords(content, subject)
	expected := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestHandleWords Failed!")
	}
}

func TestHandleCharacters(t *testing.T) {
	file, _ := os.ReadFile("../banner-files/standard.txt") // Read standard.txt
	content := strings.Split(string(file), "\n")           // Split file line-by-line

	subject := "Hello"
	got := HandleCharacters("", subject, content)
	expected := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"

	// Compare got & expected
	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestHandleCharaters Failed!")
	}
}

func TestEscapeSequence(t *testing.T) {
	subject := []string{"\\a", "\\b", "\\t", "\\v", "\\f"}

	for _, str := range subject {
		if !EscapeSequence(str) {
			t.Errorf("Got: %t", false)
			t.Errorf("Expected: %t", true)
			t.Errorf("TestEscapeSequence Failed!")
			t.FailNow()
		}
	}
}

func TestIsPrintable(t *testing.T) {
	subject := []string{"\\n", "\\r", "~", " ", "g"} // Include valid printable characters

	for _, str := range subject {
		if !IsPrintable(str) {
			t.Errorf("Got: %t", false)
			t.Errorf("Expected: %t", true)
			t.Errorf("TestEscapeSequence Failed!")
			t.FailNow()
		}
	}
}

func TestValidFile(t *testing.T) {
	subject := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"} // Include all valid file names

	// Compare if all names in subject are valid
	for _, str := range subject {
		if !IsPrintable(str) {
			t.Errorf("Got: %t", false)
			t.Errorf("Expected: %t", true)
			t.Errorf("TestValidFile Failed!")
			t.FailNow()
		}
	}
}

func TestFileIntegrity(t *testing.T) {
	file, _ := os.ReadFile("../banner-files/standard.txt")

	//Generate hash for banner file
	hash := sha256.New()
	hash.Write(file)
	got := fmt.Sprintf("%x", hash.Sum(nil))

	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"

	// Compare hashes
	if got != standard {
		t.Errorf("Got: %t", false)
		t.Errorf("Expected: %t", true)
		t.Errorf("TestFileIntegrity Failed!")
	}
}

