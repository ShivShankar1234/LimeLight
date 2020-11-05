package trie

import "errors"

type Character string
type Keyword string

type Trie struct {
	children map[Character] *Trie
	IndexedWord string
}

func NewTrie() *Trie {
	return &Trie{
		children: make(map[Character] *Trie),
		IndexedWord: "",
	}
}

/* Returns TRUE if the node is not the root node in the Trie. */
func (t *Trie) IsKeyword() bool {
	return t.IndexedWord != ""
}

/* Retrieves a word from the Trie by checking if the corresponding Character */
func (t *Trie) Retrieve(c Character) (*Trie, error){
	if value, ok := t.children[c]; ok {
		return value, nil
	}
	return nil, errors.New("The word you are looking for is not in the Trie")
}

/* Insert the word into the Trie, assumimg its not already present. */
func (t *Trie) Index(word Keyword) error {
	var currentTrieNode = t;
	for _, char := range word {
		nextTrieNode, err := currentTrieNode.Retrieve(Character (char))
		if err == nil {
			currentTrieNode = nextTrieNode;
		} else {
			//lock
			currentTrieNode.children[Character(char)] = NewTrie()
			currentTrieNode = currentTrieNode.children[Character(char)]
		}
	}
	currentTrieNode.IndexedWord = string(word)
	return nil
}

/* Checks if the current TrieNode has C as one of its children. */
func (t *Trie) HasChar(c Character) bool {
	_, present := t.children[c]
	return present
}
