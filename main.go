package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	fmt.Println(&arr)

	fmt.Println(&arr[0])
	// fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	arr = append([]int{}, arr[0], arr[2])
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
}
