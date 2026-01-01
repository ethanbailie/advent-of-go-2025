package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func exp(power int) int {
	return int(math.Pow10(power))
}

func checkVals(valueRange string) []int {
	values := strings.Split(valueRange, "-")
	startStr := strings.TrimSpace(values[0])
	endStr := strings.TrimSpace(values[1])

	startMod := len(startStr)

	var res []int
	var curr int

	startVal, err := strconv.Atoi(startStr)
	if err != nil {
		panic(err)
	}
	endVal, err := strconv.Atoi(endStr)
	if err != nil {
		panic(err)
	}

	if startMod%2 != 0 {
		startVal = exp(startMod)
		startMod += 1
	}

	for startVal <= endVal {
		curr = startVal % exp(startMod/2)
		if curr == ((startVal - curr) / exp(startMod/2)) {
			res = append(res, startVal)
		}
		startVal += 1

		if startVal == exp(startMod) {
			startMod += 2
			startVal *= 10
		}
	}

	return res
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var nums []int
	var res int

	vals := strings.SplitSeq(string(data), ",")
	for val := range vals {
		nums = append(nums, checkVals(val)...)
	}

	for _, num := range nums {
		res += num
	}
	fmt.Println(res)
}
