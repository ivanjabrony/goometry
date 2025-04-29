package main

import (
	ostovtree "gogeom/mathobjects/ostovTree"
	"gogeom/matrixes"
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	const width, height = 1000, 1000 // Размер изображения
	dc := matrixes.NewDrawingContext(width, height)

	n := 20
	//scale := 1.0
	bottom_range := -400
	top_range := 400
	points := make([][2]float64, n)

	for i := range n {
		points[i] = [2]float64{
			float64(randRange(bottom_range, top_range)),
			float64(randRange(bottom_range, top_range))}
		dc.DrawPoint(points[i][0], points[i][1], 3)
		dc.Stroke()
	}

	tree := ostovtree.FindOstovTree(points)

	for _, e := range tree {
		dc.DrawLine(e[0][0], e[0][1], e[1][0], e[1][1])
		dc.Stroke()
	}

	dc.SavePNG("coordinate_plane.png")
}
