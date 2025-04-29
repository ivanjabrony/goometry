package hw3_n1

import "gogeom/matrixes"

const (
	ax, ay                    = -10.0, -7.0
	bx, by                    = 15.0, 23.0
	cx, cy                    = 30.0, 0.0
	width, height, hasNumbers = 100, 100, false
)

func Hw3_n1_a_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	matrixes.DrawLineBerz(dc, ax, ay, bx, by, true)
	matrixes.DrawLineBerz(dc, bx, by, cx, cy, true)
	matrixes.DrawLineBerz(dc, cx, cy, ax, ay, true)
	dc.SavePNG("hw/hw3/n1/bezier.png")

	dc = matrixes.NewDrawingContext(width, height, hasNumbers)
	matrixes.DrawLineStupid(dc, ax, ay, bx, by, true)
	matrixes.DrawLineStupid(dc, bx, by, cx, cy, true)
	matrixes.DrawLineStupid(dc, ax, ay, cx, cy, true)

	dc.SavePNG("hw/hw3/n1/simple.png")
}
