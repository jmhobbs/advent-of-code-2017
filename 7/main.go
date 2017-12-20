package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nodes := []*Node{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		node, err := NewNode(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nodes = append(nodes, node)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part A:", FindRootNode(nodes).Name)
}
