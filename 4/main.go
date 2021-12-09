package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var cards []card
var calls []int
var found map[int]bool

type card struct {
	lines []struct {
		numbers []bingoNr
	}
}

type bingoNr struct {
	number int
	hit    bool
}

func main() {
	foundFirst := false
	for _, call := range calls {
		hit(call)
		fSum, lSum := check()
		if fSum != 0 && !foundFirst {
			foundFirst = true
			fmt.Println("Part Juan:", fSum*call)
		}
		if lSum != 0 {
			fmt.Println("Part Too:", lSum*call)
			return
		}
	}

}

func hit(call int) {
	for i, c := range cards {
		for j, l := range c.lines {
			for k, bn := range l.numbers {
				if bn.number == call {
					cards[i].lines[j].numbers[k].hit = true
				}
			}
		}
	}
}

func check() (int, int) {
	firstSum := 0

	for cardIndex, c := range cards {
		for _, l := range c.lines {
			count := 0
			for _, bn := range l.numbers {
				if bn.hit {
					count++
				}
			}

			if count == 5 && firstSum == 0 {
				firstSum = uncalled(c)
			}
			_, keyExists := found[cardIndex]

			if count == 5 && len(found) == (len(cards)-1) && !keyExists {
				fmt.Println(cardIndex)
				fmt.Println(found)
				fmt.Println(c)
				return firstSum, uncalled(c)
			}

			if count == 5 {
				found[cardIndex] = true
			}
		}

		for i := 0; i < 5; i++ {
			cou := 0

			for _, line := range c.lines {
				if line.numbers[i].hit {
					cou++
				}
			}

			if cou == 5 && firstSum == 0 {
				firstSum = uncalled(c)
			}

			_, keyExists := found[cardIndex]

			if cou == 5 && len(found) == (len(cards)-1) && !keyExists {
				fmt.Println(cardIndex)
				fmt.Println(found)
				fmt.Println(c)
				return firstSum, uncalled(c)
			}

			if cou == 5 {
				found[cardIndex] = true
			}
		}
	}

	return firstSum, 0
}

func uncalled(c card) (sum int) {
	for _, l := range c.lines {
		for _, bn := range l.numbers {
			if !bn.hit {
				sum += bn.number
			}
		}
	}
	return
}

func init() {
	found = make(map[int]bool)

	calls = []int{
		26, 55, 7, 40, 56, 34, 58, 90, 60, 83, 37, 36, 9, 27, 42, 19, 46, 18, 49, 52, 75, 17, 70, 41, 12, 78, 15, 64, 50, 54, 2, 77, 76, 10, 43, 79, 22, 32, 47, 0, 72, 30, 21, 82, 6, 95, 13, 59, 16, 89, 1, 85, 57, 62, 81, 38, 29, 80, 8, 67, 20, 53, 69, 25, 23, 61, 86, 71, 68, 98, 35, 31, 4, 33, 91, 74, 14, 28, 65, 24, 97, 88, 3, 39, 11, 93, 66, 44, 45, 96, 92, 51, 63, 84, 73, 99, 94, 87, 5, 48,
	}
	cards = make([]card, 0)

	input, err := ioutil.ReadFile("./4/input.txt")
	if err != nil {
		panic(err)
	}

	cs := strings.Split(string(input), "\n\n")

	for _, c := range cs {
		var card card
		ls := strings.Split(string(c), "\n")

		for _, l := range ls {
			s := strings.Split(string(l), ",")

			line := make([]bingoNr, 0, 5)

			for _, num := range s {
				n, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}

				line = append(line, bingoNr{
					number: n,
				})
			}

			card.lines = append(card.lines, struct {
				numbers []bingoNr
			}{
				numbers: line,
			})
		}

		cards = append(cards, card)
	}
}
