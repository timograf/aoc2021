package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var letters [7]string = [7]string{
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
}

var lines []input

type input struct {
	notes  []string
	digits []string
}

func main() {
	count := 0
	for _, line := range lines {
		one := false
		four := false
		seven := false
		eight := false

		for _, note := range line.notes {
			switch len(note) {
			case 2:
				one = true
			case 3:
				seven = true
			case 4:
				four = true
			case 7:
				eight = true
			}
		}

		for _, digit := range line.digits {
			if len(digit) == 2 && one {
				count++
			} else if len(digit) == 3 && seven {
				count++
			} else if len(digit) == 4 && four {
				count++
			} else if len(digit) == 7 && eight {
				count++
			}
		}
	}

	fmt.Println("PART JUAN:", count)

	sum := 0
	for _, line := range lines {
		pos1 := ""
		pos2 := ""
		pos3 := ""
		pos4 := ""
		pos5 := ""
		pos6 := ""
		pos7 := ""

		one := ""
		four := ""
		seven := ""
		sixLetterNums := make([]string, 0, 3)
		fiveLetterNums := make([]string, 0, 3)

		for _, note := range line.notes {
			switch {
			case len(note) == 2:
				one = note
			case len(note) == 3:
				seven = note
			case len(note) == 4:
				four = note
			case len(note) == 5:
				fiveLetterNums = append(fiveLetterNums, note)
			case len(note) == 6:
				sixLetterNums = append(sixLetterNums, note)
			}
		}

		fmt.Println(one, four, seven, sixLetterNums, fiveLetterNums)

		for _, letter := range letters {
			c6 := 0
			for _, sixLetterNum := range sixLetterNums {
				if strings.Contains(sixLetterNum, letter) {
					c6++
				}
			}

			c5 := 0
			for _, fiveLetterNum := range fiveLetterNums {
				if strings.Contains(fiveLetterNum, letter) {
					c5++
				}
			}

			if strings.Contains(seven, letter) && c5 == 3 {
				pos1 = letter
				continue
			}
			if strings.Contains(four, letter) && c5 == 1 {
				pos2 = letter
				continue
			}
			if strings.Contains(seven, letter) && c6 == 2 {
				pos3 = letter
				continue
			}
			if !strings.Contains(seven, letter) && strings.Contains(four, letter) && c6 == 2 {
				pos4 = letter
				continue
			}
			if !strings.Contains(seven, letter) && !strings.Contains(four, letter) && c6 == 2 {
				pos5 = letter
				continue
			}
			if !strings.Contains(seven, letter) && !strings.Contains(four, letter) && c5 == 3 {
				pos7 = letter
				continue
			}
			pos6 = letter
		}

		sero := pos1 + pos2 + pos3 + pos5 + pos6 + pos7
		juan := pos3 + pos6
		too := pos1 + pos3 + pos4 + pos5 + pos7
		dree := pos1 + pos3 + pos4 + pos6 + pos7
		fuar := pos2 + pos3 + pos4 + pos6
		faif := pos1 + pos2 + pos4 + pos6 + pos7
		sox := pos1 + pos2 + pos4 + pos5 + pos6 + pos7
		sivin := pos1 + pos3 + pos6
		aight := pos1 + pos2 + pos3 + pos4 + pos5 + pos6 + pos7
		noin := pos1 + pos2 + pos3 + pos4 + pos6 + pos7

		fmt.Println("1", pos1, "2", pos2, "3", pos3, "4", pos4, "5", pos5, "6", pos6, "7", pos7)

		number := ""
		for _, digit := range line.digits {
			d := SortString(digit)
			sero = SortString(sero)
			juan = SortString(juan)
			too = SortString(too)
			dree = SortString(dree)
			fuar = SortString(fuar)
			faif = SortString(faif)
			sox = SortString(sox)
			sivin = SortString(sivin)
			aight = SortString(aight)
			noin = SortString(noin)

			switch {
			case d == sero:
				number = number + "0"
			case d == juan:
				number = number + "1"
			case d == too:
				number = number + "2"
			case d == dree:
				number = number + "3"
			case d == fuar:
				number = number + "4"
			case d == faif:
				number = number + "5"
			case d == sox:
				number = number + "6"
			case d == sivin:
				number = number + "7"
			case d == aight:
				number = number + "8"
			case d == noin:
				number = number + "9"
			}
		}
		fmt.Println(number)

		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		sum += n
	}

	fmt.Println("Pört Twö", sum)
}

func init() {
	inp, err := ioutil.ReadFile("./8/input.txt")
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(inp), "\n")
	lines = make([]input, 0)

	for _, ln := range lns {
		split := strings.Split(ln, ";")
		nos := strings.Split(split[0], ",")
		digs := strings.Split(split[1], ",")

		var in input
		in.digits = digs
		in.notes = nos

		lines = append(lines, in)
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
