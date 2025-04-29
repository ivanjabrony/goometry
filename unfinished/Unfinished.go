package unfinished

import (
	"fmt"
	"gogeom/mathobjects"
	"gogeom/matrixes"
	"image/color"
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	const width, height = 1000, 1000 // Размер изображения
	dc := matrixes.NewDrawingContext(width, height)

	n := 50
	scale := 1.0
	bottom_range := 0
	top_range := 400
	points := make([][2]float64, n)

	for i := range n {
		points[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
	}

	hull := mathobjects.BuildConvexHullJarvis(points)

	matrix_polygon := make([]matrixes.Matrix, len(hull))

	for i, v := range hull {
		matrix_polygon[i] = matrixes.CreateMatrix([][]float64{
			{v[0] * scale},
			{v[1] * scale},
			{1}})
	}

	matrixes.DrawPolygon(dc, matrix_polygon)

	for _, p := range hull {
		dc.SetColor(color.RGBA{100, 50, 200, 100})
		dc.DrawPoint(p[0]*scale, p[1]*scale, 3)
		dc.Stroke()
	}

	k := 5

	k_points := make([][2]float64, k)

	for i := range k {
		k_points[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
	}

	for _, p := range k_points {
		dc.SetColor(color.RGBA{100, 50, 200, 100})
		dc.DrawPoint(p[0]*scale, p[1]*scale, 3)
		dc.Stroke()

		isInside := mathobjects.IsPointInsidePolygonWinding(hull, p[0], p[1])
		fmt.Printf("Is Point(%v, %v) inside? -%v\n", p[0], p[1], isInside)
	}

	dc.SavePNG("coordinate_plane.png")
}
