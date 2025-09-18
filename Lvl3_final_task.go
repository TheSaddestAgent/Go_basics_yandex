package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func getTopWords(wordMap map[string]int, n int) []string {
	if len(wordMap) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordMap[keys[i]] > wordMap[keys[j]]
	})

	if len(keys) > n {
		return keys[:n]
	}
	return keys
}
func AnalyzeText(text string) {
	cnt := make(map[string]int)
	isSeparator := func(r rune) bool {
		return unicode.IsSpace(r) || r == '.' || r == ',' || r == '!' || r == '?' || r == ';' || r == ':'
	}

	words := strings.FieldsFunc(text, isSeparator)
	for i, word := range words {
		words[i] = strings.ToLower(word)
		cnt[words[i]]++
	}
	sz := 0
	for key := range cnt {
		sz += cnt[key]
	}
	cnt_uni := 0
	for key := range cnt {
		if cnt[key] == 1 {
			cnt_uni++
		}
	}
	tops := getTopWords(cnt, 5)
	fmt.Printf("Количество слов: %d\nКоличество уникальных слов: %d\nСамое часто встречающееся слово: \"%s\" (встречается %d раз)\nТоп-5 самых часто встречающихся слов:\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n", sz, len(cnt), tops[0], cnt[tops[0]], tops[0], cnt[tops[0]], tops[1], cnt[tops[1]], tops[2], cnt[tops[2]], tops[3], cnt[tops[3]], tops[4], cnt[tops[4]])
}
