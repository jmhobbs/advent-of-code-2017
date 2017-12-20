package main

import "regexp"

type Node struct {
	Name     string
	Weight   int
	Children []string
}

var matcher *regexp.Regexp

func init() {
	matcher = regexp.MustCompile("")
}

func NewNode(description string) Node {
	return Node{}
}
