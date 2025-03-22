package main

//import "gonum.org/v1/gonum/mat"

import (
	"fmt"
	"gogeom/mathobjects"
	"gogeom/matrixes"
)

func main() {

	const width, height = 1000, 1000 // Размер изображения
	dc := matrixes.NewDrawingContext(width, height)

	lines := [][4]float64{
		{1, 1, 2, 5},
		{2, 7, 5, -1},
		{3, 4, 0, 5},
		{3, 1, 2, 6},
		{4, 8, 5, 0}}

	for _, v := range lines {

		dc.DrawLine(v[0]*50, v[1]*50, v[2]*50, v[3]*50)
		dc.Stroke()
	}

	for i, _ := range lines {
		for j := 0; j < i; j++ {
			ans1 := mathobjects.IsIntersectSimple(
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3])

			ans2 := mathobjects.IsIntersectWithPlacement(
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3])

			fmt.Printf("Line 1: (%v, %v)-(%v, %v) Line 2: (%v, %v)-(%v, %v) == %v and %v \n ",
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3],
				ans1,
				ans2)
		}
	}

	//println(mathobjects.IsIntersectSimple(1, 1, 2, 5, 2, 7, 5, -1))

	dc.SavePNG("coordinate_plane.png")
}
