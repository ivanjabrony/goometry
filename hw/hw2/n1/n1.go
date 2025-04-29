package hw2_n1

import (
	"fmt"
	"gogeom/mathobjects"
	"gogeom/matrixes"
	"math"
)

const (
	a             = 200.0
	b             = 100.0
	width, height = 1000, 1000 // Размер изображения
	hasNumbers    = true
)

func Hw2_n1_a() {
	fmt.Println("ELLIPSE AND HYPERBOLA")
	fmt.Print("Equation for ellipse: \n { x(t) = a * cos(t)\n { y(t) = b * sin(t)\n\n")
	fmt.Print("Equation for hyperbola: \n { x(t) = a * ch(t)\n { y(t) = b * sh(t)\n\n")
}

func Hw2_n1_b() {
	fmt.Print("Equation of tangent for ellipse: \n { x(s) = a * cos(t_0) - a * sin(t_0) * s\n { y(s) = b * sin(t_0) + b * cos(t_0) * s\n\n")
	fmt.Print("Equation of tangent for hyperbola: \n { x(s) = a * ch(t_0) + a * sh(t_0) * s\n { y(s) = b * sh(t_0) + b * ch(t_0) * s\n\n")
}

func Hw2_n1_c() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	n := 40.0

	f1, f2 := mathobjects.GetEllipseFunctions(a, b)
	for t := float64(0); t < 2*math.Pi; t += 2 * math.Pi / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx, f2(t)-ty, f1(t)+tx, f2(t)+ty)
		dc.Stroke()
	}
	dc.SavePNG("hw/hw2/n1/ellipse/task1_c.png")

	n = 50.0
	dc = matrixes.NewDrawingContext(width, height, hasNumbers)
	f1, f2 = mathobjects.GetHyperbolicFunctions(a, b)
	for t := -6.0; t < 6.0; t += 6.28 / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx, f2(t)-ty, f1(t)+tx, f2(t)+ty)
		dc.Stroke()
		dc.DrawLine(-f1(t)-tx/5, -f2(t)-ty/5, -f1(t)+tx/5, -f2(t)+ty/5)
		dc.Stroke()
	}

	dc.SavePNG("hw/hw2/n1/hyperbola/task1_c.png")
}

func Hw2_n1_d() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	n := 40.0
	offset := n/3 - 1

	f1, f2 := mathobjects.GetEllipseFunctions(a, b)
	for t := float64(0); t < 2*math.Pi; t += 2 * math.Pi / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx/offset, f2(t)-ty/offset, f1(t)+tx/offset, f2(t)+ty/offset)
		dc.Stroke()
	}
	dc.SavePNG("hw/hw2/n1/ellipse/task1_d.png")

	n = 50.0
	dc = matrixes.NewDrawingContext(2*width, height, hasNumbers)
	f1, f2 = mathobjects.GetHyperbolicFunctions(a, b)
	for t := -6.0; t < 6.0; t += 2 * math.Pi / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx/15, f2(t)-ty/15, f1(t)+tx/15, f2(t)+ty/15)
		dc.Stroke()
		dc.DrawLine(-f1(t)-tx/15, -f2(t)-ty/15, -f1(t)+tx/15, -f2(t)+ty/15)
		dc.Stroke()
	}
	dc.SavePNG("hw/hw2/n1/hyperbola/task1_d.png")
}

func Hw2_n1_e() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	n := 40.0
	offset := n/3 - 1

	f1, f2 := mathobjects.GetEllipseFunctions(a, b)
	for t := float64(0); t < 2*math.Pi; t += 2 * math.Pi / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx/offset, f2(t)-ty/offset, f1(t)+tx/offset, f2(t)+ty/offset)
		dc.Stroke()
	}

	for t := float64(0); t < 2*math.Pi; t += 2 * math.Pi / 10000 {
		x, y := mathobjects.GetEvolute(f1, f2, t)
		dc.DrawPoint(x, y, 1)
		dc.Stroke()
	}
	dc.SavePNG("hw/hw2/n1/ellipse/task1_e.png")
	fmt.Print("Equation of evoluta of ellipse:\n { x(t) = (a^2 - b^2)/a * cos^3(t)\n { y(t) = (b^2 - a^2)/b * sin^3(t)\n\n")

	n = 10000.0
	dc = matrixes.NewDrawingContext(2*width, height, hasNumbers)
	f1, f2 = mathobjects.GetHyperbolicFunctions(a, b)
	for t := -6.0; t < 6.0; t += 2 * math.Pi / n {
		tx, ty := mathobjects.GetTangetVectorParam(f1, f2, float64(t))
		dc.DrawLine(f1(t)-tx/offset, f2(t)-ty/offset, f1(t)+tx/offset, f2(t)+ty/offset)
		dc.Stroke()
		dc.DrawLine(-f1(t)-tx/offset, -f2(t)-ty/offset, -f1(t)+tx/15, -f2(t)+ty/offset)
		dc.Stroke()
	}
	for t := 0.0; t < 2*math.Pi; t += 2 * math.Pi / n {
		x, y := mathobjects.GetEvolute(f1, f2, t)
		if x <= 10000 && x >= -10000 && y <= 10000 && y >= -10000 {
			dc.DrawPoint(x, y, 1)
			dc.Stroke()
			dc.DrawPoint(-x, -y, 1)
			dc.Stroke()
			dc.DrawPoint(-x, y, 1)
			dc.Stroke()
			dc.DrawPoint(x, -y, 1)
			dc.Stroke()
		}
	}
	dc.SavePNG("hw/hw2/n1/hyperbola/task1_e.png")
	fmt.Print("Equation of evoluta of hyperbola:\n { x(t) = (a^2 + b^2)/a * ch^3(t)\n { y(t) = -(a^2 + b^2)/b * sh^3(t)\n\n")
}
