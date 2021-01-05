package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mask struct {
	off int64
	on  int64
}

func (mask *Mask) ApplyOn(value int64) int64 {
	value = value | mask.on
	value = value & mask.off
	return value
}

func (mask *Mask) Process(input string) {
	maskOr := strings.ReplaceAll(input, "X", "0")
	maskAnd := strings.ReplaceAll(input, "X", "1")

	mask.off, _ = strconv.ParseInt(maskAnd, 2, 64)
	mask.on, _ = strconv.ParseInt(maskOr, 2, 64)
}

func sumMemory() int64 {
	var sum int64 = 0
	for _, v := range memory {
		sum = sum + v
	}
	return sum
}

func resetMemory() {
	memory = make(map[int]int64)
}

var (
	memory = make(map[int]int64)
	mask   Mask
)

func main() {
	TaskOne()
	resetMemory()
	TaskTwo()
}

func TaskOne() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {

		var (
			addr      int
			value     int64
			maskInput string
		)
		input := scanner.Text()
		if n, _ := fmt.Sscanf(input, "mask = %s", &maskInput); n == 1 {
			mask.Process(maskInput)
		} else if n, _ := fmt.Sscanf(input, "mem[%d] = %d", &addr, &value); n == 2 {
			memory[addr] = mask.ApplyOn(value)
		}
	}

	fmt.Println("Task 1 -> Sum of all values in memory:", sumMemory())
}

func TaskTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var (
		addr             int
		value            int64
		maskInput        string
		floatingBitIndex = []int{}
	)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		input := scanner.Text()
		if n, _ := fmt.Sscanf(input, "mask = %s", &maskInput); n == 1 {

		} else if n, _ := fmt.Sscanf(input, "mem[%d] = %d", &addr, &value); n == 2 {
			for i, bit := range maskInput {
				if bit == '1' {
					addr |= 1 << i
				} else if bit == 'X' {
					floatingBitIndex = append(floatingBitIndex, i)
				}
			}

			addrVariants := []int{addr}
			for _, i := range floatingBitIndex {
				for _, addr := range addrVariants {
					bit1Variant := addr | 1<<i
					bit0Variant := addr ^ 1<<i
					addrVariants = append(addrVariants, bit1Variant, bit0Variant)
				}
			}

			for _, addr := range addrVariants {
				memory[addr] = value
			}
		}
	}

	fmt.Println("Task 2 -> Sum of all values in memory:", sumMemory())
}
