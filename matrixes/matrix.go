package matrixes

import (
	"errors"
	"fmt"
	"strings"
)

type Matrix struct {
	Length int
	Width  int
	Data   [][]float64
}

func CreateMatrix(data [][]float64) Matrix {
	if data[0] == nil {
		return Matrix{}
	}
	l := len(data)
	w := len(data[0])
	m := Matrix{l, w, data}
	return m
}

func CreateEmptyMatrix(length, width int) (Matrix, error) {
	if length <= 0 || width <= 0 {
		return Matrix{}, errors.New("invalid sizes")
	}
	data := make([][]float64, length)
	for i := range length {
		data[i] = make([]float64, width)
	}
	return Matrix{length, width, data}, nil
}

func Sum(m1, m2 Matrix) (Matrix, error) {
	if m1.Length != m2.Length || m1.Width != m2.Width {
		return Matrix{}, errors.New("incompatible sizes")
	}
	sum, _ := CreateEmptyMatrix(m1.Length, m1.Width)
	for i := 0; i < m1.Length; i++ {
		for j := 0; j < m1.Width; j++ {
			sum.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
		}
	}
	return sum, nil
}

func Multiply(m1, m2 Matrix) (Matrix, error) {
	if m1.Width != m2.Length {
		return Matrix{}, errors.New("incompatible sizes")
	}

	mult, _ := CreateEmptyMatrix(m1.Length, m2.Width)
	n := m1.Length
	m := m1.Width
	p := m2.Width
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			sum := float64(0)
			for k := 0; k < m; k++ {
				sum = sum + m1.Data[i][k]*m2.Data[k][j]
			}
			mult.Data[i][j] = sum
		}
	}
	return mult, nil
}

func (m Matrix) String() string {
	if m.Length == 0 || m.Width == 0 {
		return "[] (empty matrix)"
	}

	// First pass: find maximum width needed for each column
	colWidths := make([]int, m.Width)
	for _, row := range m.Data {
		for j, val := range row {
			// Format the number and measure its string length
			s := fmt.Sprintf("%.2f", val)
			if len(s) > colWidths[j] {
				colWidths[j] = len(s)
			}
		}
	}

	var sb strings.Builder
	for _, row := range m.Data {
		sb.WriteString("[ ")
		for j, val := range row {
			// Format with padding to match column width
			format := fmt.Sprintf("%%%d.2f", colWidths[j])
			sb.WriteString(fmt.Sprintf(format, val))
			if j < len(row)-1 {
				sb.WriteString("  ") // Add spacing between columns
			}
		}
		sb.WriteString(" ]\n")
	}
	return sb.String()
}

func (m Matrix) ConcatHorizontal(others ...Matrix) (Matrix, error) {
	for _, other := range others {
		if m.Length != other.Length {
			return Matrix{}, errors.New("matrices must have same number of rows for horizontal concatenation")
		}
	}

	totalWidth := m.Width
	for _, other := range others {
		totalWidth += other.Width
	}

	result, _ := CreateEmptyMatrix(m.Length, totalWidth)

	for i := range m.Data {
		copy(result.Data[i], m.Data[i])
	}

	currentCol := m.Width
	for _, other := range others {
		for i := range other.Data {
			copy(result.Data[i][currentCol:], other.Data[i])
		}
		currentCol += other.Width
	}

	return result, nil
}
