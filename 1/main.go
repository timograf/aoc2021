package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./1/input.txt")
	if err != nil {
		panic(err)
	}

	ln := strings.Split(string(input), "\n")

	count := 0

	for i, n := range ln {
		d, err := strconv.Atoi(string(n))
		if err != nil {
			panic(err)
		}

		if i > 2 {
			prevD, err := strconv.Atoi(string(ln[i-3]))
			if err != nil {
				panic(err)
			}

			if d > prevD {
				count++
			}
		}
	}

	fmt.Println(count)
}
