// https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/770/
package rotate_image

func rotate(matrix [][]int) {
	length := len(matrix)
	rotateFrom := make([][][]int, length)
	for i, eachLine := range rotateFrom {
		eachLine = make([][]int, length)
		rotateFrom[i] = eachLine
		for j := 0; j < length; j++ {
			rotateFrom[i][j] = []int{i - j, length - 1 - i - j}
		}
	}

	for i := 0; i < length-1; i++ {
		for j := i; j < length-1-i; j++ {
			x := i
			y := j
			temp := matrix[y][x]
			for r := 0; r < 3; r++ {
				x, y = innerRotate(matrix, rotateFrom, x, y)
			}
			matrix[y][x] = temp
		}
	}
}

func innerRotate(matrix [][]int, rotateFrom [][][]int, x, y int) (nextX, nextY int) {
	nextX = x + rotateFrom[y][x][0]
	nextY = y + rotateFrom[y][x][1]
	matrix[y][x] = matrix[nextY][nextX]
	return
}
