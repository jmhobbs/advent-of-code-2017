package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Weight   int
	Children []string
}

var matcher *regexp.Regexp

func init() {
	matcher = regexp.MustCompile("([a-z]+) \\(([0-9]+)\\)( -> ([a-z, ]+))?")
}

func NewNode(description string) (*Node, error) {
	matches := matcher.FindAllStringSubmatch(description, -1)
	if matches == nil {
		return nil, fmt.Errorf("Description incorrectly formatted.")
	}

	n := Node{}
	n.Name = matches[0][1]

	var err error

	n.Weight, err = strconv.Atoi(matches[0][2])
	if err != nil {
		return nil, fmt.Errorf("Error converting weight: %s", err)
	}

	if matches[0][4] != "" {
		n.Children = strings.Split(matches[0][4], ", ")
	} else {
		n.Children = []string{}
	}

	return &n, nil
}

func FindRootNode(nodes []*Node) *Node {
	children := []string{}
	possible := []*Node{}
	for _, node := range nodes {
		if len(node.Children) > 0 {
			children = append(children, node.Children...)
			possible = append(possible, node)
		}
	}

	for _, maybe_root := range possible {
		is_root := true
		for _, child := range children {
			if maybe_root.Name == child {
				is_root = false
				break
			}
		}
		if is_root {
			return maybe_root
		}
	}
	return nil
}
