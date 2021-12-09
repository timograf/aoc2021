package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var mapp [][]int
var baisinSum int
var alreadyCounted [][]bool

func main() {
	bigBabo1, bigBabo2, bigBabo3 := 0, 0, 0
	sum := 0
	for i, x := range mapp {
		for j, num := range x {
			cond1 := true
			cond2 := true
			cond3 := true
			cond4 := true

			xs, ys := surrounding(i, j)

			if len(xs) > 0 {
				cond1 = num < mapp[xs[0]][ys[0]]
			}
			if len(xs) > 1 {
				cond2 = num < mapp[xs[1]][ys[1]]
			}
			if len(xs) > 2 {
				cond3 = num < mapp[xs[2]][ys[2]]
			}
			if len(xs) > 3 {
				cond4 = num < mapp[xs[3]][ys[3]]
			}
			if cond1 && cond2 && cond3 && cond4 {
				sum++
				sum += num

				baisinSum = 0
				alreadyCounted = make([][]bool, 100)
				for i, _ := range mapp {
					alreadyCounted[i] = make([]bool, 100)
				}
				expand(i, j)
				if baisinSum > bigBabo1 {
					bigBabo3 = bigBabo2
					bigBabo2 = bigBabo1
					bigBabo1 = baisinSum
				} else if baisinSum > bigBabo2 {
					bigBabo3 = bigBabo2
					bigBabo2 = baisinSum
				} else if baisinSum > bigBabo3 {
					bigBabo3 = baisinSum
				}
			}

		}
	}
	fmt.Println("one", sum)
	fmt.Println("two", bigBabo1*bigBabo2*bigBabo3)
}

func surrounding(x, y int) (xs []int, ys []int) {
	if x != 0 {
		xs = append(xs, x-1)
		ys = append(ys, y)
	}
	if x != len(mapp)-1 {
		xs = append(xs, x+1)
		ys = append(ys, y)
	}
	if y != 0 {
		xs = append(xs, x)
		ys = append(ys, y-1)
	}
	if y != len(mapp[x])-1 {
		xs = append(xs, x)
		ys = append(ys, y+1)
	}

	return
}

func expand(x, y int) {
	if mapp[x][y] != 9 {
		baisinSum++
		alreadyCounted[x][y] = true
		xs, ys := surrounding(x, y)
		for i, x1 := range xs {
			if !alreadyCounted[x1][ys[i]] {
				expand(x1, ys[i])
			}
		}
	}
}

func init() {
	mapp = make([][]int, 100)
	for i, _ := range mapp {
		mapp[i] = make([]int, 0)
	}
	input, err := ioutil.ReadFile("./9/input.txt")
	if err != nil {
		panic(err)
	}

	lns := strings.Split(string(input), "\n")

	for i, ln := range lns {
		nums := strings.Split(ln, "")
		for _, num := range nums {
			num, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			mapp[i] = append(mapp[i], num)
		}
	}
}
