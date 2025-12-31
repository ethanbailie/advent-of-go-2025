package main

import (
	"os"
	"strconv"
	"strings"
)

func check_vals(value_range string) {
	values := strings.Split(value_range, "-")
	start_str := values[0]
	end_str := values[1]

	start_mod := len(start_str)
	end_mod := len(end_str)

	start_val, err := strconv.Atoi(start_str)
	if err != nil {
		panic(err)
	}
	end_val, err := strconv.Atoi(end_str)
	if err != nil {
		panic(err)
	}

	for start_mod <= end_mod {
		if start_mod%2 != 0 {
			start_mod += 1
			continue
		}

		start_mod += 2
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	vals := strings.SplitSeq(string(data), ",")
	for val := range vals {
	}
}
