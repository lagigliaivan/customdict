package dictionary

type Dictionary struct {
	words map[string]bool
}

func New() Dictionary {
	return Dictionary{
		words: make(map[string]bool),
	}
}
