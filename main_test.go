package main

import (
	"os"
	"testing"
)

const testFilePath = "./War_and_Peace.txt"

func BenchmarkWordCount(b *testing.B) {
	file, err := os.Open(testFilePath)
	if err != nil {
		b.Fatalf("failed reading file: %s", err)
	}
	defer file.Close()
	for i := 0; i < b.N; i++ {
		_, err := file.Seek(0, 0)
		if err != nil {
			b.Fatalf("failed seeking file: %s", err)
		}

		_, err = wordCounts(file)
		if err != nil {
			b.Fatalf("failed counting words: %s", err)
		}
	}
}
