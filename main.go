package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const filepath = "./War_and_Peace.txt"

// WordCount хранит слово и его частоту
type WordCount struct {
	Word  string
	Count int
}

// wordCounts считывает данные и подсчитывает частоту слов
func wordCounts(r io.Reader) (map[string]int, error) {
	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(strings.ToLower(line))
		for _, word := range words {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCount, nil
}

// sortByValue сортирует мапу по значениям и возвращает топ-10 слов
func sortByValue(m map[string]int) []WordCount {
	var counts []WordCount
	for word, count := range m {
		counts = append(counts, WordCount{Word: word, Count: count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	return counts
}

func main() {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	wordCountMap, err := wordCounts(file)
	if err != nil {
		log.Fatalf("failed counting words: %s", err)
	}

	sorted := sortByValue(wordCountMap)

	for i := 0; i < 10; i++ {
		fmt.Printf("word: \"%s\"\t count: %d\n", sorted[i].Word, sorted[i].Count)
	}

}
