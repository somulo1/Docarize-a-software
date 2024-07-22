package lib

import (
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
