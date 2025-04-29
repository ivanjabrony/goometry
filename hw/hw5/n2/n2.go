package hw5_n2

import (
	"fmt"
	"gogeom/mathobjects/polygons"
	"gogeom/matrixes"
	"image/color"
	"math/rand/v2"
)

const (
	bottom_range = 0
	top_range    = 400
	n            = 20

	width, height, hasNumbers = 1000, 1000, true
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func Hw5_n2_a_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	points1 := make([][2]float64, n)
	points2 := make([][2]float64, n)

	for i := range n {
		points1[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
		dc.DrawPoint(points1[i][0], points1[i][1], 3)
		dc.Stroke()

		points2[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
		dc.DrawPoint(points2[i][0], points2[i][1], 3)
		dc.Stroke()
	}

	hull1 := polygons.BuildConvexHullJarvis(points1)
	hull2 := polygons.BuildConvexHullGraham(points2)
	matrix_polygon1 := make([]matrixes.Matrix, len(hull1))
	matrix_polygon2 := make([]matrixes.Matrix, len(hull2))

	for i, v := range hull1 {
		matrix_polygon1[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}
	for i, v := range hull2 {
		matrix_polygon2[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}

	matrixes.DrawPolygon(dc, matrix_polygon1)
	matrixes.DrawPolygon(dc, matrix_polygon2)

	intersection := polygons.IntersectConvexHulls(hull1, hull2)

	for i := range n {
		if polygons.IsPointInsidePolygon(intersection, points1[i]) {
			dc.SetColor(color.RGBA{40, 100, 200, 100})
			dc.DrawPoint(points1[i][0], points1[i][1], 5)
			dc.Stroke()
		}
		if polygons.IsPointInsidePolygon(intersection, points2[i]) {
			dc.SetColor(color.RGBA{40, 100, 200, 100})
			dc.DrawPoint(points2[i][0], points2[i][1], 5)
			dc.Stroke()
		}
	}

	for _, p := range points1 {
		if polygons.IsPointInsidePolygon(intersection, p) {
			fmt.Printf("point of E1 - (%v, %v) is inside intersection\n", p[0], p[1])
		}
	}
	for _, p := range points2 {
		if polygons.IsPointInsidePolygon(intersection, p) {
			fmt.Printf("point of E2 - (%v, %v) is inside intersection\n", p[0], p[1])
		}
	}

	dc.SavePNG("hw/hw5/n2/task_a_b.png")
}
