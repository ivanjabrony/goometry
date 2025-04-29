package hw4_n2

import (
	"fmt"
	"gogeom/mathobjects/polygons"
	"gogeom/matrixes"
	"math/rand/v2"
)

const (
	n                         = 10
	width, height, hasNumbers = 1000, 1000, true
)

var (
	polygon = [][2]float64{
		{-300, 300},
		{-100, 100},
		{100, 300},
		{300, 100},
		{100, -300},
		{-100, -100},
		{-200, -200},
	}
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func Hw4_n2_a() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	points := make([][2]float64, n)
	for i := range n {
		points[i] = [2]float64{
			float64(randRange(-width/2, width/2)),
			float64(randRange(-height/2, height/2))}
	}

	matrix_polygon := make([]matrixes.Matrix, len(polygon))

	for i, v := range polygon {
		matrix_polygon[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}

	matrixes.DrawPolygon(dc, matrix_polygon)

	fmt.Println("SUMM OF ANGLES")
	for _, p := range points {
		dc.DrawPoint(p[0], p[1], 3)
		dc.Stroke()
		matrixes.WritePointCoordinates(dc, p[0], p[1])
		fmt.Printf("For point (%f, %f) answer is: %v (angle summ)\n", p[0], p[1], polygons.IsPointInsidePolygonWinding(polygon, p))
	}

	dc.SavePNG("hw/hw4/n2/winding.png")
}

func Hw4_n2_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	points := make([][2]float64, n)
	for i := range n {
		points[i] = [2]float64{
			float64(randRange(-width/2, width/2)),
			float64(randRange(-height/2, height/2))}
	}

	matrix_polygon := make([]matrixes.Matrix, len(polygon))

	for i, v := range polygon {
		matrix_polygon[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}

	matrixes.DrawPolygon(dc, matrix_polygon)

	fmt.Println("RAYCAST")
	for _, p := range points {
		dc.DrawPoint(p[0], p[1], 3)
		dc.Stroke()
		matrixes.WritePointCoordinates(dc, p[0], p[1])
		fmt.Printf("For point (%f, %f) answer is: %v (raycast)\n", p[0], p[1], polygons.IsPointInsidePolygon(polygon, p))
	}

	dc.SavePNG("hw/hw4/n2/raycast.png")
}
