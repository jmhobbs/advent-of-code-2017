package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
}

type Row []int
type Spreadsheet []Row

func parse(input io.Reader) Spreadsheet {
	ss := Spreadsheet{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		r := Row{}
		for _, c := range strings.Fields(scanner.Text()) {
			i, err := strconv.Atoi(c)
			if err == nil {
				r = append(r, i)
			}
		}
		ss = append(ss, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ss
}

func checksum(s Spreadsheet) int {
	return 0
}
