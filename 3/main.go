package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./3/input.txt")
	if err != nil {
		panic(err)
	}

	ln := strings.Split(string(input), "\n")

	bools := make([][]bool, 0, 1000)
	for i, bitString := range ln {
		bits := strings.Split(bitString, "")
		boolSlice := make([]bool, 0, 12)
		bools = append(bools, boolSlice)

		for _, bit := range bits {
			num, err := strconv.ParseBool(bit)
			bools[i] = append(bools[i], num)
			if err != nil {
				panic(err)
			}
		}
	}

	g := ""
	e := ""

	for i := 0; i < 12; i++ {
		tc := 0

		for _, bit := range bools {
			if bit[i] {
				tc++
			}
		}

		if tc > len(bools)/2 {
			g = g + "1"
			e = e + "0"
		} else {
			g = g + "0"
			e = e + "1"
		}
	}

	gamma, err := strconv.ParseInt(g, 2, 64)
	epsilon, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part one:", gamma*epsilon)

	oxBools := make([][]bool, len(bools))
	coBools := make([][]bool, len(bools))
	var ox, co2 int64
	copy(oxBools, bools)
	copy(coBools, bools)

	for i := 0; i < 12; i++ {
		tc := 0
		fc := 0

		for _, bit := range coBools {
			if bit[i] {
				tc++
			} else {
				fc++
			}
		}

		if tc >= fc {
			newCoBools := make([][]bool, 0)

			for j, b := range coBools {
				if !b[i] {
					newCoBools = append(newCoBools, coBools[j])
				}
			}

			coBools = newCoBools
		} else {
			newCoBools := make([][]bool, 0)

			for j, b := range coBools {
				if b[i] {
					newCoBools = append(newCoBools, coBools[j])
				}
			}

			coBools = newCoBools
		}

		if len(coBools) == 1 {
			fmt.Println("CO2:", coBools[0])
			co2, err = toInt(coBools[0])
			if err != nil {
				panic(err)
			}
			break
		}
	}

	for i := 0; i < 12; i++ {
		tc := 0
		fc := 0

		for _, bit := range oxBools {
			if bit[i] {
				tc++
			} else {
				fc++
			}
		}

		if tc >= fc {
			newOxBools := make([][]bool, 0)

			for j, bit := range oxBools {
				if bit[i] {
					newOxBools = append(newOxBools, oxBools[j])
				}
			}
			oxBools = newOxBools
		} else {
			newOxBools := make([][]bool, 0)

			for j, bit := range oxBools {
				if !bit[i] {
					newOxBools = append(newOxBools, oxBools[j])
				}
			}
			oxBools = newOxBools
		}

		if len(oxBools) == 1 {
			fmt.Println("Oxygen:", oxBools[0])
			ox, err = toInt(oxBools[0])
			if err != nil {
				panic(err)
			}
			break
		}
	}

	fmt.Println("Part two:", ox*co2)
}

func toInt(binary []bool) (int64, error) {
	s := ""
	for _, bit := range binary {
		if bit {
			s += "1"
		} else {
			s += "0"
		}
	}

	return strconv.ParseInt(s, 2, 64)
}
