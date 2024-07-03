package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))

	// firstRowLength := 0
	for i, r := range rows {
		cols := strings.Split(strings.TrimSpace(r), " ")
		matrix[i] = make([]int, len(cols))

		// if firstRowLength == 0 {
		// 	firstRowLength = len(r)
		// }

		// if len(r) != firstRowLength {
		// 	return nil, errors.New("error Invalid matrix")
		// }

		for j, c := range cols {
			val, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}

			matrix[i][j] = val
		}
	}

	return matrix, matrix.validate()
}

func (m Matrix) validate() error {
	length := 0
	for _, row := range m {
		if length == 0 {
			length = len(row)
		}

		if len(row) != length {
			return errors.New("error Invalid Matrix")
		}
	}
	return nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	newMatrix := make(Matrix, len(m[0]))
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, len(m))
		for j := 0; j < len(newMatrix[i]); j++ {
			newMatrix[i][j] = m[j][i]
		}
	}
	return newMatrix
}

func (m Matrix) Rows() [][]int {
	newMatrix := make(Matrix, len(m))
	for i, row := range m {
		newMatrix[i] = make([]int, len(m[i]))
		copy(newMatrix[i], row)
	}
	return newMatrix
}

func (m Matrix) Set(row, col, val int) bool {
	if m == nil {
		return false
	}

	if row < 0 || row >= len(m) {
		return false
	}

	if col < 0 || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}
