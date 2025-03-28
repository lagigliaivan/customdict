package dictionary_test

import (
	"testing"

	"github.com/lagigliaivan/customdic/dictionary"
)

// Add Word: Add a word to the dictionary.
// Search Word: Check if a word exists in the dictionary.
// Auto-Complete: Given a prefix, return a list of words that start with that prefix.
// Remove Word: Remove a word from the dictionary.
// Additional items

// Create unit tests to showcase every functionality.
// Assume the dictionary would work in a concurrent environment.

func TestAddWordToDictionary(t *testing.T) {
	dic := dictionary.New()

	dic.Add("hello")
	ok := dic.IsPresent("hello")

	if !ok {
		t.Errorf("Expected %v but got %v", true, ok)
	}

}
