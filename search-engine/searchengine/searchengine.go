package searchengine

import (
	"strings"
	"unicode"
)

var stopWords = map[string]bool{"is": true, "the": true, "a": true, "this": true, "with": true, "and": true, "that": true}

func BuildInvertedIndex(documents []string) map[string][]int {
	invertedIndex := make(map[string][]int)

	for i, doc := range documents {
		words := Tokenize(doc)
		for _, word := range words {
			invertedIndex[word] = append(invertedIndex[word], i)
		}
	}
	return invertedIndex
}

func Tokenize(text string) []string {
	var tokens []string
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		cleaned := CleanWord(word)
		if cleaned != "" && !stopWords[cleaned] {
			tokens = append(tokens, cleaned)
		}
	}
	return tokens
}

func CleanWord(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, word)
}

func IntersectIndices(result [][]int) []int {
	counter := make(map[int]int)
	for _, list := range result {
		seen := make(map[int]bool)
		for _, index := range list {
			if !seen[index] {
				counter[index]++
				seen[index] = true
			}
		}
	}
	var intersection []int
	for k, v := range counter {
		if v == len(result) {
			intersection = append(intersection, k)
		}
	}
	return intersection
}

func Search(query string, index map[string][]int) []int {
	tokens := Tokenize(query)
	if len(tokens) == 0 {
		return nil
	}
	var results [][]int
	for _, token := range tokens {
		if idxList, ok := index[token]; ok {
			results = append(results, idxList)
		} else {
			return nil
		}
	}
	return IntersectIndices(results)
}
