package matrix

// Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it in-
func setZeroes(matrix [][]int) {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				m1[i] = true
				m2[j] = true
			}
		}
	}

	for v := range m1 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v][i] = 0
		}
	}

	for v := range m2 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}
}

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	direction := 1
	i, j := 0, 0
	var res []int
loop:
	for left <= right && up <= down {
		res = append(res, matrix[i][j])
		switch direction {
		case 1:
			if j == right {
				direction = 2
				up++
				i++
				continue loop
			}
			j++
		case 2:
			if i == down {
				direction = 3
				right--
				j--
				continue loop
			}
			i++
		case 3:
			if j == left {
				direction = 4
				down--
				i--
				continue loop
			}
			j--
		case 4:
			if i == up {
				direction = 1
				left++
				j++
				continue loop
			}
			i--
		}
	}
	return res
}

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	target := make([][]int, len(matrix))
	for i := 0; i < len(target); i++ {
		target[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(target[0]); j++ {
			target[i][j] = matrix[i][j]
		}
	}

	row := len(target[0])
	col := len(target)
	if row <= len(matrix) {
		matrix = matrix[:row]
	} else {
		for i := 0; i < row-len(matrix); i++ {
			matrix = append(matrix, make([]int, len(matrix[0])))
		}
	}

	if col <= len(matrix[0]) {
		for i := 0; i < len(matrix); i++ {
			matrix[i] = matrix[i][:col]
		}
	} else {
		for i := 0; i < len(matrix); i++ {
			matrix[i] = append(matrix[i], make([]int, col-len(matrix[0]))...)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = target[len(matrix[i])-j-1][i]
		}
	}

	// fmt.Println(target)
	// fmt.Println(matrix)
}

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]

// Given word = "ABCCED", return true.
// Given word = "SEE", return true.
// Given word = "ABCB", return false.
func exist(board [][]byte, word string) bool {
	m := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		m[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if validate(m, board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func validate(m [][]bool, board [][]byte, word string, i, j, wordIndex int) bool {
	if wordIndex == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || m[i][j] || board[i][j] != word[wordIndex] {
		return false
	}

	m[i][j] = true
	wordIndex++

	res := validate(m, board, word, i+1, j, wordIndex) ||
		validate(m, board, word, i-1, j, wordIndex) ||
		validate(m, board, word, i, j+1, wordIndex) ||
		validate(m, board, word, i, j-1, wordIndex)
	if !res {
		m[i][j] = false
	}
	return res
}
