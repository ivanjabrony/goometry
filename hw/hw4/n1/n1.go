package hw4_n1

import (
	"fmt"
	"gogeom/mathobjects"
	"gogeom/matrixes"
)

const (
	width, height, hasNumbers = 1000, 1000, true
)

var lines = [][4]float64{
	{100, 100, 200, 400},
	{-200, 350, 450, -100},
	{-200, -400, 0, 400},
	{300, 100, -200, 200},
	{-200, -300, -200, 0}}

func Hw4_n1_a() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	for _, v := range lines {
		dc.DrawLine(v[0], v[1], v[2], v[3])
		dc.Stroke()
	}

	fmt.Println("PARAMETRISATION")
	for i := range lines {
		for j := 0; j < i; j++ {
			xp, yp, ans1 := mathobjects.IsIntersectSimple(
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3])

			if xp != nil && yp != nil {
				fmt.Printf("  Line 1: (%v, %v):(%v, %v) Line 2: (%v, %v):(%v, %v) == %v in dot(%f, %f) (parametrisation)\n",
					lines[i][0],
					lines[i][1],
					lines[i][2],
					lines[i][3],
					lines[j][0],
					lines[j][1],
					lines[j][2],
					lines[j][3],
					ans1,
					*xp,
					*yp)
			} else {
				fmt.Printf("  Line 1: (%v, %v):(%v, %v) Line 2: (%v, %v):(%v, %v) == %v (parametrisation)\n",
					lines[i][0],
					lines[i][1],
					lines[i][2],
					lines[i][3],
					lines[j][0],
					lines[j][1],
					lines[j][2],
					lines[j][3],
					ans1)
			}
		}
	}
	dc.SavePNG("hw/hw4/n1/task1_a.png")
}

func Hw4_n1_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	for _, v := range lines {
		dc.DrawLine(v[0], v[1], v[2], v[3])
		dc.Stroke()
	}

	fmt.Println("PLACEMENT")
	for i := range lines {
		for j := 0; j < i; j++ {
			ans1 := mathobjects.IsIntersectWithPlacement(
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3])

			fmt.Printf("  Line 1: (%v, %v):(%v, %v) Line 2: (%v, %v):(%v, %v) == %v (placement)\n",
				lines[i][0],
				lines[i][1],
				lines[i][2],
				lines[i][3],
				lines[j][0],
				lines[j][1],
				lines[j][2],
				lines[j][3],
				ans1)
		}
	}
	dc.SavePNG("hw/hw4/n1/task1_b.png")
}

func Hw4_n1_с() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	for _, v := range lines {
		dc.DrawLine(v[0], v[1], v[2], v[3])
		dc.Stroke()
	}

	crossedPairs := mathobjects.FindIntersectionsOfNLines(lines)

	fmt.Println("SWEEP LINE")
	for _, v := range crossedPairs {
		fmt.Printf("  Line (%v, %v):(%v, %v) crosses with (%v, %v):(%v, %v) (Bently-Ottman algorithm)\n",
			v[0][0], v[0][1], v[0][2], v[0][3],
			v[1][0], v[1][1], v[1][2], v[1][3])
	}
	dc.SavePNG("hw/hw4/n1/task1_c.png")
}
