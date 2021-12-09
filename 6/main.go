package main

import "fmt"

var fish []int
var fish0 int
var fish1 int
var fish2 int
var fish3 int
var fish4 int
var fish5 int
var fish6 int
var fish7 int
var fish8 int

func main() {
	for i := 0; i < 80; i++ {
		for j, f := range fish {
			if f == 0 {
				fish = append(fish, 8)
				f = 6
			} else {
				f -= 1
			}
			fish[j] = f
		}
	}
	fmt.Println("one", len(fish))

	for i := 0; i < 256; i++ {
		tempFish := fish0
		fish0 = fish1
		fish1 = fish2
		fish2 = fish3
		fish3 = fish4
		fish4 = fish5
		fish5 = fish6
		fish6 = fish7 + tempFish
		fish7 = fish8
		fish8 = tempFish
	}

	fmt.Println("two", fish0+fish1+fish2+fish3+fish4+fish5+fish6+fish7+fish8)
}

func init() {
	fish = []int{
		1, 1, 1, 1, 2, 1, 1, 4, 1, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 3, 1, 1, 1, 5, 1, 3, 1, 4, 1, 2, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 4, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 4, 1, 1, 1, 1, 3, 5, 1, 1, 2, 1, 1, 1, 1, 4, 4, 1, 1, 1, 4, 1, 1, 4, 2, 4, 4, 5, 1, 1, 1, 1, 2, 3, 1, 1, 4, 1, 5, 1, 1, 1, 3, 1, 1, 1, 1, 5, 5, 1, 2, 2, 2, 2, 1, 1, 2, 1, 1, 1, 1, 1, 3, 1, 1, 1, 2, 3, 1, 5, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 3, 2, 1, 1, 1, 4, 3, 1, 1, 4, 1, 5, 4, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 4, 5, 1, 1, 1, 1, 5, 4, 1, 3, 1, 1, 1, 1, 4, 3, 3, 3, 1, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 5, 1, 3, 1, 4, 3, 1, 3, 1, 5, 1, 1, 1, 1, 3, 1, 5, 1, 2, 4, 1, 1, 4, 1, 4, 4, 2, 1, 2, 1, 3, 3, 1, 4, 4, 1, 1, 3, 4, 1, 1, 1, 2, 5, 2, 5, 1, 1, 1, 4, 1, 1, 1, 1, 1, 1, 3, 1, 5, 1, 2, 1, 1, 1, 1, 1, 4, 4, 1, 1, 1, 5, 1, 1, 5, 1, 2, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2, 4, 1, 1, 2, 1, 1, 3, 2,
	}

	for _, f := range fish {
		switch f {
		case 1:
			fish1++
		case 2:
			fish2++
		case 3:
			fish3++
		case 4:
			fish4++
		case 5:
			fish5++
		}
	}
}
