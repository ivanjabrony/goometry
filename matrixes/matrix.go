package matrixes

import "errors"

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
