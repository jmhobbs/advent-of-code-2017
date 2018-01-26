package main

func main() {}

type HexGrid struct {
	Q int
	R int
}

func New() *HexGrid {
	return &HexGrid{0, 0}
}

func (hg *HexGrid) Move(direction string) {
	// TODO
	return
}

func (hg *HexGrid) DistanceFromOrigin() int {
	// TODO
	return 0
}
