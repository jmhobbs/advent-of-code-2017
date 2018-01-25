package main

func main() {}

type KnotHash struct {
	Size     int
	List     []int
	Position int
	SkipSize int
}

func New(size int) *KnotHash {
	kh := KnotHash{size, make([]int, size), 0, 0}
	for i := 0; i < size; i++ {
		kh.List[i] = i
	}
	return &kh
}

func (kh *KnotHash) Knot(length int) {
}
