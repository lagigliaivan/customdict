package dictionary_test

import (
	"sync"
	"testing"

	"github.com/lagigliaivan/customdict/dictionary"
	"github.com/stretchr/testify/assert"
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
	ok := dic.Contains("hello")

	if !ok {
		t.Errorf("Expected %v but got %v", true, ok)
	}
}

func TestSearchWordInDictionary(t *testing.T) {
	dic := dictionary.New()

	dic.Add("hello")
	dic.Add("world")
	dic.Add("hello world")
	dic.Add("world hell")
	dic.Add("world hello")
	dic.Add("foo")
	dic.Add("bar")

	ok := dic.Contains("hello")
	if !ok {
		t.Errorf("Expected %v but got %v", true, ok)
	}

	ok = dic.Contains("hell")
	if ok {
		t.Errorf("Expected %v but got %v", false, ok)
	}
}

func TestAutoComplete(t *testing.T) {
	dic := dictionary.New()

	dic.Add("hello")
	dic.Add("world")
	dic.Add("hello world")
	dic.Add("world hell")
	dic.Add("world hello")
	dic.Add("foo")
	dic.Add("bar")

	words := dic.Prefix("hell")
	if words.Len() != 4 {
		t.Errorf("Expected %v but got %v", 4, words.Len())
	}

	ok := words.Contains("hello")
	if !ok {
		t.Errorf("Expected %v but got %v", true, ok)
	}

	ok = words.Contains("hello world")
	if !ok {
		t.Errorf("Expected %v but got %v", false, ok)
	}

	ok = words.Contains("world hello")
	if !ok {
		t.Errorf("Expected %v but got %v", false, ok)
	}

	ok = words.Contains("world hell")
	if !ok {
		t.Errorf("Expected %v but got %v", false, ok)
	}
}

func TestRemoveWordFromDictionary(t *testing.T) {
	dic := dictionary.New()

	dic.Add("hello")
	dic.Add("world")
	dic.Add("hello world")
	dic.Add("world hell")
	dic.Add("world hello")
	dic.Add("foo")
	dic.Add("bar")

	dic.Remove("hello")
	ok := dic.Contains("hello")
	if ok {
		t.Errorf("Expected %v but got %v", false, ok)
	}
}

func TestRaceCondition(t *testing.T) {
	var present bool

	wg := sync.WaitGroup{}
	wg.Add(3)

	dictionary := dictionary.New()

	dictionary.Add("hello")

	go func() {
		dictionary.Add("hello1")
		wg.Done()
	}()
	go func() {
		present = dictionary.Contains("hello")
		wg.Done()
	}()
	go func() {
		dictionary.Remove("hello1")
		wg.Done()
	}()

	wg.Wait()
	assert.True(t, present)
}
