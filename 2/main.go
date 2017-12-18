package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := parse(file)

	log.Println("Checksum A", checksum(s))
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
	acc := 0
	for _, row := range s {
		acc = acc + checksumRow(row)
	}
	return acc
}

func checksumRow(r Row) int {
	min := r[0]
	max := r[0]
	for _, i := range r {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}
	return max - min
}

func evenChecksum(s Spreadsheet) int {
	acc := 0
	for _, row := range s {
		acc = acc + evenRowChecksum(row)
	}
	return acc
}

func evenRowChecksum(r Row) int {
	return 0
}
