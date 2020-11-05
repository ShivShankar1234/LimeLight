package LimeLight

import "github.com/ShivShankar1234/LimeLight/trie"
/**
The structural basis for LimeLight search is the Aho-Corasick Automaton.
https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
Additional changes to the algorithm have been made here to allow for concurrency.
 */

func extractKeywords(t *trie.Trie, sentence string) []string {
	matches := make([]string, 0)
	currentTrie := t;

	idx:= 0
	sentenceLen := len(sentence)

	for idx < sentenceLen {
		char := string(sentence[idx])
		if isWordBoundary(trie.Character(char)){
			idx2, longestSequenceFound := checkIfMatch(currentTrie, sentence, idx)
			idx = idx2
			if longestSequenceFound != "" {
				matches = append(matches, longestSequenceFound)
			}
			currentTrie = t
		} else if insideTrie, _ := currentTrie.Retrieve(trie.Character(char)); insideTrie != nil {
			currentTrie = insideTrie
		} else {
			currentTrie = t
			idy := idx + 1
			for idy < sentenceLen {
				char := sentence[idy]
				if isWordBoundary(trie.Character(char)) {
					break
				}
				idy += 1
			}
			idx = idy
		}
		if idx + 1 >= sentenceLen {
			if currentTrie.IsKeyword() {
				matches = append(matches, currentTrie.IndexedWord)
			}
		}
		idx += 1
	}
	return matches
}

type Keywords struct {
	t *trie.Trie
}

func NewKeywords() Keywords {
	return Keywords{trie.NewTrie()}
}

func (x Keywords) Extract(sentence string) []string {
	return extractKeywords(x.t, sentence)
}

func (x Keywords) Add(w string) {
	x.t.Index(trie.Keyword(w))
}

func isWordBoundary(c trie.Character) bool {
	return c == "" || c == " " || c == "\t" || c == "\n"
}


