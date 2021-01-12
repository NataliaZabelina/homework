package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

type wordCounterStruct struct {
	Key   string
	Value int
}

func Top10(input string) []string {
	if input == "" {
		return []string{}
	}

	expression := regexp.MustCompile(`\s{2,}`)
	inputText := expression.ReplaceAllString(input, " ")
	text := strings.Split(inputText, " ")

	wordCounterMap := make(map[string]int)
	for _, element := range text {
		_, ok := wordCounterMap[element]
		if ok {
			wordCounterMap[element]++
		} else {
			wordCounterMap[element] = 1
		}
	}

	delete(wordCounterMap, "")
	if len(wordCounterMap) == 0 {
		return []string{}
	}

	words := []wordCounterStruct{}
	for k, v := range wordCounterMap {
		words = append(words, wordCounterStruct{k, v})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].Value > words[j].Value
	})

	size := int(math.Min(float64(len(wordCounterMap)), 10))
	wordsSlice := make([]string, size)
	for i, l := range words {
		if i >= size {
			break
		}
		wordsSlice[i] = l.Key
	}

	return wordsSlice
}
