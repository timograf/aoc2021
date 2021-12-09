package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./2/input.txt")
	if err != nil {
		panic(err)
	}

	ln := strings.Split(string(input), "\n")

	h, d := 0, 0

	for _, c := range ln {
		command := strings.Split(c, " ")
		m, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case "forward":
			h += m
		case "up":
			d -= m
		case "down":
			d += m
		default:
			panic("unknown command")
		}
	}
	fmt.Printf("Part one: %d", d*h)

	h, a, d := 0, 0, 0

	for _, c := range ln {
		command := strings.Split(c, " ")
		m, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err)
		}

		switch command[0] {
		case "forward":
			h += m
			d += a * m
		case "up":
			a -= m
		case "down":
			a += m
		default:
			panic("unknown command")
		}
	}
	fmt.Printf("Part two: %d", d*h)
}
