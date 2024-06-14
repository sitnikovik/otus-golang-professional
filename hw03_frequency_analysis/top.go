package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	k string
	v int
}

func Top10(s string) []string {
	words := strings.Fields(s)
	n := len(words)
	if n == 0 {
		return []string{}
	}

	kvs := makeKvs(words)
	sortKvs(kvs)

	return getFirstUpTo(kvs, 10)
}

func makeKvs(ss []string) []kv {
	n := len(ss)
	kvs := make([]kv, n)
	idxmap := make(map[string]int, n)
	excludedWords := make(map[string]struct{}, n)
	rexp := regexp.MustCompile(`^[a-zA-Zа-яА-Я]|[^\w-]|^\d+`)

	for i := 0; i < n; i++ {
		word := prepareWord(ss[i])

		if _, ok := excludedWords[word]; ok {
			continue
		}
		if !isStringValid(rexp, word) {
			excludedWords[word] = struct{}{}
			continue
		}

		if _, ok := idxmap[word]; ok {
			kvs[idxmap[word]].v++
			continue
		}

		idxmap[word] = i
		kvs[i] = kv{
			k: word,
			v: 1,
		}
	}

	return kvs
}

func prepareWord(s string) string {
	return strings.TrimRight(strings.ToLower(s), ".")
}

func isStringValid(rexp *regexp.Regexp, s string) bool {
	return rexp.MatchString(s)
}

func sortKvs(kvs []kv) {
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v > kvs[j].v || (kvs[i].v == kvs[j].v && kvs[i].k < kvs[j].k)
	})
}

func getFirstUpTo(kvs []kv, toIdx int) []string {
	n := toIdx
	if n > len(kvs) {
		n = len(kvs)
	}

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if kvs[i].v == 0 {
			continue
		}

		res = append(res, kvs[i].k)
	}

	return res
}
