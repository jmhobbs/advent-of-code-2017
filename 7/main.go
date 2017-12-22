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

	root := BuildTree(nodes)

	log.Println("Part A:", root.Name)
	_, bal := GetUnbalancedNodeWeights(root, -1)
	log.Println("Part B:", bal)
}

func GetUnbalancedNodeWeights(current *Node, expected_total_weight int) (unbalanced, corrected int) {
	type Weighed struct {
		Node   *Node
		Weight int
		Seen   int
	}

	weights := [2]Weighed{Weighed{nil, -1, 0}, Weighed{nil, -1, 0}}
	for _, child := range current.ChildNodes {
		weight := child.TotalWeight()

		if weights[0].Node == nil {
			weights[0] = Weighed{child, weight, 1}
			continue
		} else if weights[1].Node == nil && weights[0].Weight != weight {
			weights[1] = Weighed{child, weight, 1}
			continue
		}

		if weight == weights[0].Weight {
			weights[0].Seen += 1
		} else {
			weights[1].Seen += 1
		}
	}

	if weights[0].Seen == 1 {
		if len(weights[0].Node.ChildNodes) > 0 {
			return GetUnbalancedNodeWeights(weights[0].Node, weights[1].Weight)
		}
		return weights[0].Weight, weights[1].Weight
	} else if weights[1].Seen == 1 {
		if len(weights[1].Node.ChildNodes) > 0 {
			return GetUnbalancedNodeWeights(weights[1].Node, weights[0].Weight)
		}
		return weights[1].Weight, weights[0].Weight
	}

	return current.Weight, current.Weight - (current.TotalWeight() - expected_total_weight)
}
