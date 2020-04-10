package mergesort

func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	middle := len(slice) / 2

	left := mergeSort(slice[0:middle])
	right := mergeSort(slice[middle:len(slice)])
	return merge(left, right)
}

func merge(left, right []int) []int {
	p1, p2 := 0, 0

	var res []int
	for p1 < len(left) && p2 < len(right) {
		if left[p1] > right[p2] {
			res = append(res, right[p2])
			p2++
			continue
		}
		res = append(res, left[p1])
		p1++
	}
	res = append(res, left[p1:len(left)]...)
	res = append(res, right[p2:len(right)]...)

	return res
}
