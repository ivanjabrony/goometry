package matrixes

import (
	"math"
)

func GetMovementMatrix(x, y float64) Matrix {
	data := [][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetRotationMatrix(phi float64) Matrix {
	data := [][]float64{
		{math.Cos(phi), math.Sin(phi), 0},
		{-math.Sin(phi), math.Cos(phi), 0},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetFullRotationMatrix(x, y float64, phi float64) Matrix {
	s := math.Sin(phi)
	c := math.Cos(phi)
	data := [][]float64{
		{c, -s, -c*x + s*y + x},
		{s, c, -s*x - c*y + y},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetOsSimmetryMatrix(x, y bool) Matrix {
	var data [][]float64
	if x && !y { //Ox
		data = [][]float64{
			{1, 0, 0},
			{0, -1, 0},
			{0, 0, 1}}
	} else if !x && y { // Oy
		data = [][]float64{
			{-1, 0, 0},
			{0, 1, 0},
			{0, 0, 1}}
	} else if x && y { // x = y
		data = [][]float64{
			{0, 1, 0},
			{1, 0, 0},
			{0, 0, 1}}
	} else { // x = -y
		data = [][]float64{
			{0, -1, 0},
			{-1, 0, 0},
			{0, 0, 1}}
	}
	return CreateMatrix(data)
}

func GetFullSymmetryMatrix(x1, y1, x2, y2 float64) Matrix {
	a := y2 - y1
	b := x1 - x2
	c := x2*y1 - x1*y2
	d := a*a + b*b

	data := [][]float64{
		{1 - (2*a*a)/d, (-2 * a * b) / d, (-2 * a * c) / d},
		{(-2 * a * b) / d, 1 - (2*b*b)/d, (-2 * b * c) / d},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetStretchMatrix(x, y float64) Matrix {
	data := [][]float64{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetShearMatrixOx(k float64) Matrix {
	data := [][]float64{
		{1, k, 0},
		{0, 1, 0},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetShearMatrixOy(k float64) Matrix {
	data := [][]float64{
		{1, 0, 0},
		{k, 1, 0},
		{0, 0, 1}}
	return CreateMatrix(data)
}

func GetHomotethetyMatrix(x, y, k float64) Matrix {
	data := [][]float64{
		{k, 0, (1 - k) * x},
		{0, k, (1 - k) * y},
		{0, 0, 1}}
	return CreateMatrix(data)
}
