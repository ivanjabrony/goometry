package hw2_n2

import (
	"fmt"
	"gogeom/mathobjects"
	"gogeom/matrixes"
	"math"
)

const (
	v = 50.0
	c = 1.0
	w = 3.0
	R = 300.0
	k = 8.0

	width, height, hasNumbers = 1000, 1000, true
)

func Hw2_n2_a_b_c() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	// spiral
	f1, f2 := mathobjects.GetSpiralArchimedFunctions(v, c, w)
	for t := 0.0; t < 3*math.Pi; t += 0.05 {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, t)
		x0, y0, x1, y1 := f1(t)-tx/30, f2(t)-ty/30, f1(t)+tx/30, f2(t)+ty/30
		dc.DrawLine(x0, y0, x1, y1)
		dc.Stroke()
	}

	t0 := 2.0
	tx, ty := mathobjects.GetTangetVectorParam(f1, f2, t0)
	nx, ny := mathobjects.GetNormalVectorParam(f1, f2, t0)

	curvature_spiral := mathobjects.GetCurvature(f1, f2, t0)
	x0, y0, x1, y1 := f1(t0)-tx, f2(t0)-ty, f1(t0)+tx, f2(t0)+ty
	xn0, yn0, xn1, yn1 := f1(t0)-nx, f2(t0)-ny, f1(t0)+nx, f2(t0)+ny
	dc.DrawLine(x0, y0, x1, y1)
	dc.Stroke()
	dc.DrawLine(xn0, yn0, xn1, yn1)
	dc.Stroke()
	dc.SavePNG("hw/hw2/n2/spiral/task2_a.png")

	fmt.Println("SPIRAL")
	fmt.Print("Equation of spiral(Archimed):\n { x(t) = (c + v*t) * cos(w*t)\n { y(t) = (c + v*t) * sin(w*t)\n\n")
	fmt.Print("Equation of derivative of spiral(Archimed):\n { x'(t) = v * cos(w*t) - w * (c + v*t) * sin(w*t)\n { y'(t) = v * sin(w*t) + w * (c + v*t) * cos(w*t)\n\n")
	fmt.Print("Equation of tangent of spiral(Archimed):\n { x(s) = x(t_0) + s * x'(t_0)\n { y(s) = y(t_0) + s * y'(t_0)\n\n")
	fmt.Print("Equation of normal of spiral(Archimed):\n { x(s) = x(t_0) - s * y'(t_0)\n { y(s) = y(t_0) + s * x'(t_0)\n\n")
	fmt.Printf("Curvature at t_0 = %v: %v\n\n", t0, curvature_spiral)

	// hypocicloida
	dc = matrixes.NewDrawingContext(width, height, hasNumbers)

	f1, f2 = mathobjects.GetHypoCicloidFunctions(R, k)
	for t := 0.0; t < 3*math.Pi; t += 0.05 {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, t)
		x0, y0, x1, y1 := f1(t)-tx/20, f2(t)-ty/20, f1(t)+tx/20, f2(t)+ty/20
		dc.DrawLine(x0, y0, x1, y1)

		dc.Stroke()
	}

	t0 = 2.0
	tx, ty = mathobjects.GetTangetVectorParam(f1, f2, t0)
	nx, ny = mathobjects.GetNormalVectorParam(f1, f2, t0)

	curvature_cicloida := mathobjects.GetCurvature(f1, f2, t0)
	x0, y0, x1, y1 = f1(t0)-tx, f2(t0)-ty, f1(t0)+tx, f2(t0)+ty
	xn0, yn0, xn1, yn1 = f1(t0)-nx, f2(t0)-ny, f1(t0)+nx, f2(t0)+ny
	dc.DrawLine(x0, y0, x1, y1)
	dc.DrawLine(xn0, yn0, xn1, yn1)
	dc.Stroke()

	dc.SavePNG("hw/hw2/n2/cicloida/task2_a.png")

	fmt.Println("CICLOIDA")
	fmt.Print("Equation of hypocicloida:\n { x(t) = r * (k-1) * cos(t) + r * cos((k-1)*t)\n { y(t) = r * (k-1) * sin(t) - r * sin((k-1)*t)\n\n")
	fmt.Print("Equation of derivative of hypocicloida:\n { x'(t) = -r * (k-1) * sin(t) - r * (k-1) * sin((k-1)*t)\n { y'(t) =  r * (k-1) * cos(t) - r * (k-1) * cos((k-1)*t)\n\n")
	fmt.Print("Equation of tangent of hypocicloida:\n { x(s) = x(t_0) + s * x'(t_0)\n { y(s) = y(t_0) + s * y'(t_0)\n\n")
	fmt.Print("Equation of normal of hypocicloida:\n { x(s) = x(t_0) - s * y'(t_0)\n { y(s) = y(t_0) + s * x'(t_0)\n\n")
	fmt.Printf("Curvature at t_0 = %v: %v\n\n", t0, curvature_cicloida)
}
