package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var code []string

func main() {
	for _, line := range code {
		chars := strings.Split(line, "")
		fmt.Println(chars)

		for i := 0; i < len(chars); i++ {
			if chars[i] == "(" {
				count := 0
				start := i + 1
				for j := start; j < len(chars); j++ {
					if chars[j] == "(" {
						count++
					}
					if chars[j] == ")" && count == 0 {

						fmt.Println(chars[start:j])
						false := check(chars[start:j])
						fmt.Println(false)
						i = j
						break
					}

					if chars[j] == ")" {
						count--
					}
				}

			} else if chars[i] == "[" {
				count := 0
				start := i + 1
				for j := start; j < len(chars); j++ {
					if chars[j] == "[" {
						count++
					}
					if chars[j] == "]" && count == 0 {

						fmt.Println(chars[start:j])
						false := check(chars[start:j])
						fmt.Println(false)
						i = j
						break
					}

					if chars[j] == "]" {
						count--
					}
				}

			} else if chars[i] == "{" {
				count := 0
				start := i + 1
				for j := start; j < len(chars); j++ {
					if chars[j] == "{" {
						count++
					}
					if chars[j] == "}" && count == 0 {

						fmt.Println(chars[start:j])
						false := check(chars[start:j])
						fmt.Println(false)
						i = j
						break
					}

					if chars[j] == "}" {
						count--
					}
				}

			} else if chars[i] == "<" {
				count := 0
				start := i + 1
				for j := start; j < len(chars); j++ {
					if chars[j] == "<" {
						count++
					}
					if chars[j] == ">" && count == 0 {

						fmt.Println(chars[start:j])
						false := check(chars[start:j])
						fmt.Println(false)
						i = j
						break
					}

					if chars[j] == ">" {
						count--
					}
				}

			}
		}

		return
	}
}

func check(ss []string) string {
	count1, count2, count3, count4 := 0, 0, 0, 0

	for _, s := range ss {
		switch s {
		case "(":
			count1++
		case ")":
			count1--
		case "{":
			count2++
		case "}":
			count2--
		case "[":
			count3++
		case "]":
			count3--
		case "<":
			count4++
		case ">":
			count4--
		}
	}

	if count1 != 0 || count2 != 0 || count3 != 0 || count4 != 0 {
		return "false"
	}

	return ""
}

func init() {
	input, err := ioutil.ReadFile("./10/input.txt")
	if err != nil {
		panic(err)
	}

	code = strings.Split(string(input), "\n")
}
