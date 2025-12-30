package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := 50
	var res int
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		direction := string(line[0])
		shift := 1
		move, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == "L" {
			shift = -1
		}
		num = (num + move*shift) % 100
		if num < 0 {
			num += 100
		} else if num == 0 {
			res += 1
		}
	}

	fmt.Println(res)
}
