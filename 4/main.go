package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid := 0
	anagramValid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passphrase := scanner.Text()
		if validPassphrase(passphrase) {
			valid += 1
			if anagramValidPassphrase(passphrase) {
				anagramValid += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part A:", valid)
	log.Println("Part B:", anagramValid)
}

func validPassphrase(pp string) bool {
	seen := map[string]bool{}

	for _, word := range strings.Fields(pp) {
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = true
	}

	return true
}

func anagramValidPassphrase(pp string) bool {
	seen := map[string]bool{}

	for _, word := range strings.Fields(pp) {
		// given that the passphrases are all ASCII, I'm ok with this.
		sortable := []byte(word)
		sort.Slice(sortable, func(i, j int) bool {
			return sortable[i] < sortable[j]
		})
		sorted := string(sortable)

		if _, ok := seen[sorted]; ok {
			return false
		}
		seen[sorted] = true
	}

	return true
}
