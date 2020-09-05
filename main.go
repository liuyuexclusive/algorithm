package main

import "fmt"

func heapSort(slice []int) {
	end := len(slice) - 1
	for i := end / 2; i >= 0; i-- {
		rise(slice, i, end)
	}
	for end > 0 {
		slice[0], slice[end] = slice[end], slice[0]
		end--
		rise(slice, 0, end)
	}
}

func rise(slice []int, k, end int) {
	for {
		i := 2*k + 1
		if i > end {
			break
		}
		if i < end && slice[i+1] > slice[i] {
			i++
		}
		if slice[k] >= slice[i] {
			break
		}
		slice[i], slice[k] = slice[k], slice[i]
		k = i
	}
}

func insert(slice *[]int, item int) {
	*slice = append(*slice, item)
	end := len(*slice) - 1
	for i := (end - 1) / 2; i >= 0; i = (i - 2 + i%2) / 2 {
		rise(*slice, i, end)
	}
}

func main() {
	slice := []int{2, 1, 4, 3, 7, 6, 5}

	end := len(slice) - 1
	for i := end / 2; i >= 0; i-- {
		rise(slice, i, end)
	}

	fmt.Println(slice)

	// heapSort(slice)

	insert(&slice, 100)

	fmt.Println(slice)
}
