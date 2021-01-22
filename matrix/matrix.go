package matrix

// 矩阵置零

// 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
func setZeroes(matrix [][]int) {
	rows, cols := make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for k := range rows {
		for i := 0; i < len(matrix[k]); i++ {
			matrix[k][i] = 0
		}
	}
	for i := 0; i < len(matrix); i++ {
		for k := range cols {
			matrix[i][k] = 0
		}
	}
}

// 螺旋矩阵

// 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	const (
		right = 1
		down  = 2
		left  = 3
		up    = 4
	)
	direction := right
	h1, h2 := 0, len(matrix)-1
	w1, w2 := 0, len(matrix[0])-1
	i, j := 0, 0

	for h1 <= h2 && w1 <= w2 {
		res = append(res, matrix[i][j])
		switch direction {
		case right:
			if j == w2 {
				direction = down
				h1++
				i++
			} else {
				j++
			}
		case down:
			if i == h2 {
				direction = left
				w2--
				j--
			} else {
				i++
			}
		case left:
			if j == w1 {
				direction = up
				h2--
				i--
			} else {
				j--
			}
		case up:
			if i == h1 {
				direction = right
				w1++
				j++
			} else {
				i--
			}
		}
	}
	return res
}

// 旋转图像

// 给定一个 n × n 的二维矩阵表示一个图像。
// 将图像顺时针旋转 90 度。
// 说明：
// 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	height, width := len(matrix), len(matrix[0])

	newMatrix := make([][]int, width)
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, height)
		for j := 0; j < len(newMatrix[i]); j++ {
			newMatrix[i][j] = matrix[height-j-1][i]
		}
	}

	if len(matrix) >= len(newMatrix) {
		matrix = matrix[:len(newMatrix)]
	} else {
		matrix = append(matrix, make([][]int, len(newMatrix)-len(matrix))...)
	}

	for i := 0; i < len(newMatrix); i++ {
		for j := 0; j < len(newMatrix[i]); j++ {
			if j >= len(matrix[i]) {
				matrix[i] = append(matrix[i], newMatrix[i][j])
			} else {
				matrix[i][j] = newMatrix[i][j]
			}
			if j == len(newMatrix[i])-1 && len(matrix[i]) > len(newMatrix[i]) {
				matrix[i] = matrix[i][:j+1]
			}
		}
	}
}

// 单词搜索

// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
func exist(board [][]byte, word string) bool {
	m := make([][]bool, len(board))
	for i := 0; i < len(m); i++ {
		m[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if exist2(board, word, m, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func exist2(board [][]byte, word string, m [][]bool, i, j, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j > len(board[0]) || m[i][j] || board[i][j] != word[index] {
		return false
	}

	m[i][j] = true
	index++
	res := exist2(board, word, m, i+1, j, index) ||
		exist2(board, word, m, i-1, j, index) ||
		exist2(board, word, m, i, j+1, index) ||
		exist2(board, word, m, i, j-1, index)

	if !res {
		m[i][j] = false
	}
	return res
}
