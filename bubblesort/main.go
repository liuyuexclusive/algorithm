package main

func bubblesort(slice []int) []int {
	length, isMove := len(slice), true
	for isMove {
		isMove = false
		for i := 0; i < length-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				isMove = true
			}
		}
		length--
	}
	return slice
}
