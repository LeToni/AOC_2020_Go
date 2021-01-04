package main

const (
	N int = iota
	E
	S
	W
)

type Instruction struct {
	action int
	value  int
}

type Ship struct {
	facing int
	x, y   int
}

func main() {

}
