package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	Children    map[string]*Trie
	IsEndOfWord bool
}

func _____Constructor() Trie {
	return Trie{
		Children:    make(map[string]*Trie),
		IsEndOfWord: true,
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for i, char := range word {
		tmp := &Trie{
			Children:    make(map[string]*Trie),
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
}

func (this *Trie) Search(word string) bool {
	current := this
	for i, char := range word {
		if next, ok := current.Children[string(char)]; ok {
			if i == len(word)-1 {
				return next.IsEndOfWord
			}
			current = next
		} else {
			return false
		}
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	fmt.Println(this.String())
	current := this
	for _, char := range prefix {
		if next, ok := current.Children[string(char)]; ok {
			current = next
		} else {
			return false
		}
	}

	return true
}

func (t *Trie) String() string {
	return t.stringWithIndent(0)
}

func (t *Trie) stringWithIndent(level int) string {
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
