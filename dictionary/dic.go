package dictionary

import (
	"strings"
	"sync"
)

type Dictionary struct {
	words map[string]bool
	mutex sync.Mutex
}

func New() *Dictionary {
	return &Dictionary{
		words: make(map[string]bool),
		mutex: sync.Mutex{},
	}
}

func (d *Dictionary) Add(word string) {
	d.lock(func() {
		d.words[word] = true
	})
}

func (d *Dictionary) IsPresent(word string) bool {
	var rest bool
	d.lock(func() {
		rest = d.words[word]
	})

	return rest
}

func (d *Dictionary) Prefix(prefix string) *Dictionary {
	result := New()
	d.lock(func() {
		for i := range d.words {
			if strings.Contains(i, prefix) {
				result.Add(i)
			}
		}
	})

	return result
}

func (d *Dictionary) Remove(word string) {
	d.lock(func() {
		delete(d.words, word)
	})
}

func (d *Dictionary) Len() int {
	return len(d.words)
}

func (d *Dictionary) lock(f func()) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	f()
}
