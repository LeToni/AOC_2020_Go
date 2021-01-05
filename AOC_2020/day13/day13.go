package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	earliest int
	buslines []int
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	i := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		if i == 0 {
			earliest, _ = strconv.Atoi(scanner.Text())
			i++
		} else {
			for _, s := range strings.Split(scanner.Text(), ",") {
				if s != "x" {
					line, _ := strconv.Atoi(s)
					buslines = append(buslines, line)
				} else {
					buslines = append(buslines, -1)
				}
			}
		}
	}
	favLine := 0
	waitTime := math.MaxInt64
	for _, l := range buslines {
		if l != -1 {
			if t := l - earliest%l; t < waitTime {
				favLine = l
				waitTime = t
			}
		}
	}
	fmt.Println("Part 1 -> Product of earliest busId and minutes to wait for bus", favLine*waitTime)

	t, step := 0, 1
	for i, l := range buslines {
		if l == -1 {
			continue
		}
		for (t+i)%l != 0 {
			t += step
		}
		step *= l
	}

	fmt.Println("Part 2-> Earliest time where all bus departs with offset:", t)
}
