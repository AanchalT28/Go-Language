package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

func main() {

	files := []string{"file1.txt", "file2.txt"}
	wordCount := make(map[string]int)

	var mu sync.Mutex
	var wg sync.WaitGroup

	//process each file concurrently
	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()

			//read the entire file into memory
			data, err := os.ReadFile(file)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file, err)
				return
			}

			words := splitWords(string(data))

			//update word counts
			mu.Lock()
			for _, word := range words {
				wordCount[word]++
			}
			mu.Unlock()
		}(file)
	}

	wg.Wait()

	//collect and sort words alphabetically
	var keys []string
	for key := range wordCount {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("Word: %s | Count: %d\n", key, wordCount[key])
	}
}

// splitWords splits a string into words, converting them to lowercase
func splitWords(text string) []string {
	text = strings.ToLower(text)

	//regex to find all words(sequences of alphabetic characters)
	re := regexp.MustCompile(`[a-z]+`)
	return re.FindAllString(text, -1)
}
