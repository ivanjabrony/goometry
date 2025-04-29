package hw5_n1

import (
	"fmt"
	"gogeom/mathobjects/polygons"
	"gogeom/matrixes"
	"math/rand/v2"
)

const (
	bottom_range = 0
	top_range    = 400
	n            = 100

	width, height, hasNumbers = 1000, 1000, true
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func Hw5_n1_a_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	points := make([][2]float64, n)
	for i := range n {
		points[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
	}

	hull_jarvis := polygons.BuildConvexHullJarvis(points)
	hull_graham := polygons.BuildConvexHullGraham(points)
	matrix__jarvis := make([]matrixes.Matrix, len(hull_jarvis))
	matrix__graham := make([]matrixes.Matrix, len(hull_graham))

	for i, v := range hull_jarvis {
		matrix__jarvis[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}

	for i, v := range hull_graham {
		matrix__graham[i] = matrixes.CreateMatrix([][]float64{
			{v[0]},
			{v[1]},
			{1}})
	}

	matrixes.DrawPolygon(dc, matrix__jarvis)
	matrixes.DrawPolygon(dc, matrix__graham)

	for _, p := range points {
		dc.DrawPoint(p[0], p[1], 3)
		dc.Stroke()
	}

	perimeter_jarvis := polygons.FindPerimeter(hull_jarvis)
	area_jarvis := polygons.FindPolygonArea(hull_jarvis)
	perimeter_graham := polygons.FindPerimeter(hull_graham)
	area_graham := polygons.FindPolygonArea(hull_graham)

	fmt.Printf("Perimeter:%v, Area:%v (Jarvis) \n", perimeter_jarvis, area_jarvis)
	fmt.Printf("Perimeter:%v, Area:%v (Graham) \n", perimeter_graham, area_graham)
	dc.SavePNG("hw/hw5/n1/task_a_jarvis.png")
	dc.SavePNG("hw/hw5/n1/task_a_graham.png")
}
