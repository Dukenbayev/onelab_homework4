package top_words

import (
	"log"
	"regexp"
	"sort"
	"strings"
)
type wordCount struct {
	word  string
	count int
}

func MaxCountString(s string,n int) []string{
	if n<=0{
		return []string{}
	}
	reg, err := regexp.Compile("[^A-Za-z]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString(s, " ") //cleaned strings
	array := strings.Split(newStr," ")//

	field := make([]wordCount, 0)

	for i:=0; i<len(array);i++{
		word := string(array[i])
		count:=strings.Count(s,word)
		s = strings.Replace(s, word, "", 	-1)
		// результат: "Чистим полностью стринг
		field = append(field, wordCount{word: word, count: count})
	}
	// Сортировка по количеству слов
	sort.SliceStable(field, func(i, j int) bool {
		return field[i].count < field[j].count
	})

	result := make([]string, 0)

	for i := 0; i < n; i++ {
		result = append(result, field[i].word)
	}
	return result
}
