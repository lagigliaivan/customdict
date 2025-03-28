package dictionary

import "strings"

type Dictionary struct {
	words map[string]bool
}

func New() Dictionary {
	return Dictionary{
		words: make(map[string]bool),
	}
}

func (d *Dictionary) Add(word string) {
	d.words[word] = true
}

func (d *Dictionary) IsPresent(word string) bool {
	return d.words[word]
}

func (d *Dictionary) Prefix(prefix string) Dictionary {
	result := New()
	for i := range d.words {
		if strings.Contains(i, prefix) {
			result.Add(i)
		}
	}

	return result
}

func (d *Dictionary) Remove(word string) Dictionary {
	delete(d.words, word)
	return *d
}

func (d *Dictionary) Len() int {
	return len(d.words)
}
