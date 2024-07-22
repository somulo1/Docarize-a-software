package lib

import "testing"

func TestAciiArt(t *testing.T) {
	subject := "Hello"
	err, got := AsciiArt(subject, "standard.txt")
	expected := " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"

	// Compare got & expected
	if err != "" && got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestAsciiArt Failed!")
	}
}
