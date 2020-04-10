package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// res := reverse(13356)
	// fmt.Println(res)

	fmt.Println(2 ^ 0)
}

func reverse(x int) int {
	isNagetive := x < 0
	var res int
	str := strconv.Itoa(int(math.Abs(float64(x))))
	for k, v := range str {
		i, _ := strconv.Atoi(string(v))
		res += i * int(math.Pow10(k))
	}
	if res > math.MaxInt32 {
		return 0
	}
	if isNagetive {
		return -res
	}
	return res
}

func quickSort(slice []int, left, right int) {
	if left >= right {
		return
	}
	flag, start, end := left, left, right
	for start < end {
		for start < end {
			if slice[end] > slice[flag] {
				end--
				continue
			}
			slice[end], slice[flag] = slice[flag], slice[end]
			flag = end
			break
		}
		for start < end {
			if slice[start] < slice[flag] {
				start++
				continue
			}
			slice[start], slice[flag] = slice[flag], slice[start]
			flag = start
			break
		}
	}
	quickSort(slice, left, flag-1)
	quickSort(slice, flag+1, right)
}
