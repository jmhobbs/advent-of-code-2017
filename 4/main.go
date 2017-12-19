package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if validPassphrase(scanner.Text()) {
			valid += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part A:", valid)
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
