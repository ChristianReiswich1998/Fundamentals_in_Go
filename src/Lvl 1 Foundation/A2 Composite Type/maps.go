package main

import "fmt"

func NewShortener() map[string]string {
	// Your code here
	shortMap := make(map[string]string)

	return shortMap
}

func Save(store map[string]string, slug, url string) {
	// Your code here
	if len(url) != 0 {
		store[slug] = url
	}

	for slug, url := range store {
		if _, has := store[url]; has {
			delete(store, slug)
		}
	}
}

func Resolve(store map[string]string, slug string) string {
	// Your code here
	result := ""

	if k, v := store[slug]; v {
		result = k
	}

	return result
}

func main() {
	s := NewShortener()
	Save(s, "gh", "https://github.com")
	Save(s, "docs", "https://go.dev/doc")
	fmt.Println(Resolve(s, "gh"))
	fmt.Println(Resolve(s, "docs"))
}
