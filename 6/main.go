package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	m := Memory{}

	for _, s := range strings.Fields(string(content)) {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, i)
	}

	i := 0
	seen := map[string]bool{}
	var s string
	for {
		i += 1
		m = Reallocate(m)
		s = fmt.Sprintf("%v", m) // I'm lazy.
		if _, ok := seen[s]; ok {
			break
		}
		seen[s] = true
	}

	log.Println("Part A:", i)

	ia := i

	for {
		i += 1
		m = Reallocate(m)
		sb := fmt.Sprintf("%v", m) // I'm lazy.
		if sb == s {
			break
		}
	}

	log.Println("Part B:", i-ia)
}
