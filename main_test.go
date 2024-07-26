package main_test

import (
	"banner-art-web/ascii"
	"testing"
	// "ascii-art-web-export-file/ascii"
)

func TestLoadBanner(t *testing.T) {
	testCases := []struct {
		filename      string
		expectedChars []rune
	}{
		{"standard", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}},
		{"thinkertoy", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}},
		{"shadow", []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '+', '`', '~', '?', ',', '\'', '"', ';', '=', '_', '/', '\\', '.', '<', '>', '|', '[', ']', '{', '}', ':'}},
	}
	for _, tc := range testCases {
		banner, err := ascii.LoadBanner(tc.filename)
		if err != nil {
			t.Errorf("For file %q: Expected no error, but got %v", tc.filename, err)
		}
		// Check if the loaded banner is not empty
		if len(banner) == 0 {
			t.Errorf("For file %q: Expected banner to load characters, but got none", tc.filename)
		}
		// Check if all expected characters are in the loaded banner
		for _, char := range tc.expectedChars {
			if _, ok := banner[char]; !ok {
				t.Errorf("For file %q: Expected to find character '%c' in banner, but it was not found", tc.filename, char)
			}
		}
	}
}
