package main

import (
	"strings"
)

type WordDictionary struct {
	Children    map[string]*WordDictionary
	IsEndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		Children:    make(map[string]*WordDictionary),
		IsEndOfWord: true,
	}
}

func (this *WordDictionary) AddWord(word string) {
	current := this
	for i, char := range word {
		tmp := &WordDictionary{
			Children:    make(map[string]*WordDictionary),
			IsEndOfWord: false,
		}

		if next, ok := current.Children[string(char)]; ok {
			current = next
		} else {
			current.Children[string(char)] = tmp
			current = tmp
		}

		if i == len(word)-1 {
			current.IsEndOfWord = true
		}
	}
	// fmt.Println("addWord")
	// fmt.Println(this.String())
}

func (this *WordDictionary) Search(word string) bool {
	// fmt.Println("search word", word)
	var s func(w *WordDictionary, wordIndex int) bool
	s = func(w *WordDictionary, wordIndex int) bool {
		// fmt.Printf("s: %d\n%s", wordIndex, w.String(),)
		if wordIndex > len(word) {
			// fmt.Println("長すぎ")
			return false
		}

		if wordIndex == len(word) {
			// fmt.Println("同じ長さ: ", w.IsEndOfWord)
			return w.IsEndOfWord
		}
		if word[wordIndex] == '.' {
			for _, v := range w.Children {
				if s(v, wordIndex+1) {
					return true
				}
			}
		}
		if next, ok := w.Children[string(word[wordIndex])]; ok {
			return s(next, wordIndex+1)
		}

		return false
	}

	return s(this, 0)
}

func (t *WordDictionary) String() string {
	return t.stringWithIndent(0)
}

func (t *WordDictionary) stringWithIndent(level int) string {
	indent := strings.Repeat("  ", level)
	var result string

	for char, child := range t.Children {
		result += indent + char
		if child.IsEndOfWord {
			result += "*"
		}
		result += "\n"
		result += child.stringWithIndent(level + 1)
	}

	return result
}
