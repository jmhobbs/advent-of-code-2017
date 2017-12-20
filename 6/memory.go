package main

type Memory []int

func Reallocate(mem Memory) Memory {
	lb := FindLargestBlock(mem)
	ctr := mem[lb]

	mem[lb] = 0
	for i := 1; i <= ctr; i++ {
		j := (i + lb) % len(mem)
		mem[j] += 1
	}

	return mem
}

func FindLargestBlock(mem Memory) int {
	idx := 0
	value := 0

	for i, v := range mem {
		if v > value {
			idx = i
			value = v
		}
	}

	return idx
}
