package main

import (
	"bufio"
	"fmt"
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

	regs := make(Registers)

	max := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ins, err := ParseInstruction(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		regs.Execute(ins)
		m := regs.Max()
		if m > max {
			max = m
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Part A:", regs.Max())
	log.Println("Part B:", max)
}

type Registers map[string]int

type Comparison int

const (
	LESS_THAN Comparison = iota
	GREATER_THAN
	LESS_THAN_EQUAL
	GREATER_THAN_EQUAL
	EQUAL
	NOT_EQUAL
)

type Instruction struct {
	Subject           string
	Increment         bool
	Amount            int
	ComparisonSubject string
	Comparator        Comparison
	ComparisonValue   int
}

func ParseInstruction(line string) (Instruction, error) {
	ins := Instruction{"", true, 0, "", LESS_THAN, 0}
	split := strings.Fields(line)
	if len(split) != 7 {
		return ins, fmt.Errorf("Incorrect number of values for instruction.")
	}

	ins.Subject = split[0]
	ins.Increment = split[1] == "inc"
	v, err := strconv.Atoi(split[2])
	if err != nil {
		return ins, err
	}
	ins.Amount = v
	ins.ComparisonSubject = split[4]
	ins.Comparator = ComparisonFromString(split[5])
	v, err = strconv.Atoi(split[6])
	if err != nil {
		return ins, err
	}
	ins.ComparisonValue = v

	return ins, nil
}

func ComparisonFromString(in string) Comparison {
	switch in {
	case "<":
		return LESS_THAN
	case ">":
		return GREATER_THAN
	case "<=":
		return LESS_THAN_EQUAL
	case ">=":
		return GREATER_THAN_EQUAL
	case "==":
		return EQUAL
	case "!=":
		return NOT_EQUAL
	}
	panic(fmt.Sprintf("Invalid Comparison: %s", in))
}

func (r Registers) Execute(in Instruction) {
	v := r[in.ComparisonSubject]
	ok := false
	switch in.Comparator {
	case LESS_THAN:
		ok = v < in.ComparisonValue
	case GREATER_THAN:
		ok = v > in.ComparisonValue
	case LESS_THAN_EQUAL:
		ok = v <= in.ComparisonValue
	case GREATER_THAN_EQUAL:
		ok = v >= in.ComparisonValue
	case EQUAL:
		ok = v == in.ComparisonValue
	case NOT_EQUAL:
		ok = v != in.ComparisonValue
	}

	if ok {
		if in.Increment {
			r[in.Subject] += in.Amount
		} else {
			r[in.Subject] -= in.Amount
		}
	}
}

func (r Registers) Max() int {
	max := 0
	for _, v := range r {
		if v > max {
			max = v
		}
	}
	return max
}
