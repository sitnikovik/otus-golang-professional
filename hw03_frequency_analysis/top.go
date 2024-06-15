package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const resultLen = 10

var validWordRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Я]|[^\w-]|^\d+`)

type keyValue struct {
	key   string
	value int
}

func Top10(s string) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return nil
	}

	keyValues := makeKeyValueSlice(words)
	sortKeyValuesByDesc(keyValues)

	return getFirstUpTo(keyValues, resultLen)
}

func makeKeyValueSlice(words []string) []keyValue {
	n := len(words)
	result := make([]keyValue, n)
	wordIndexes := make(map[string]int, n)
	excludedWords := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		word := prepareWord(words[i])

		if _, ok := excludedWords[word]; ok {
			continue
		}
		if !isWordValid(word) {
			excludedWords[word] = struct{}{}
			continue
		}

		if _, ok := wordIndexes[word]; ok {
			result[wordIndexes[word]].value++
			continue
		}

		wordIndexes[word] = i
		result[i] = keyValue{
			key:   word,
			value: 1,
		}
	}

	return result
}

func prepareWord(s string) string {
	return strings.TrimRight(strings.ToLower(s), ".")
}

func isWordValid(s string) bool {
	return validWordRegexp.MatchString(s)
}

func sortKeyValuesByDesc(kvs []keyValue) {
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].value > kvs[j].value || (kvs[i].value == kvs[j].value && kvs[i].key < kvs[j].key)
	})
}

func getFirstUpTo(kvs []keyValue, max int) []string {
	n := max
	if n > len(kvs) {
		n = len(kvs)
	}

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if kvs[i].value == 0 {
			continue
		}

		res = append(res, kvs[i].key)
	}

	return res
}
