package hw3_n2

import (
	"gogeom/matrixes"
	"math"
)

const (
	ax, ay                    = 100, 200
	width, height, hasNumbers = 1000, 1000, true
)

var (
	points_2nd_order = matrixes.CreateMatrix([][]float64{
		{-100, 200, 300},
		{-100, 0, 200},
		{1, 1, 1}})
	points_3rd_order = matrixes.CreateMatrix([][]float64{
		{-100, 200, 300, 150},
		{-100, 0, 200, 0},
		{1, 1, 1, 1}})
)

func Hw3_n2_a() {
	// Original
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	matrixes.DrawPoints(dc, 19, points_2nd_order)
	f1 := matrixes.QuadraticBezierFunction(points_2nd_order)

	var n float64 = 500

	for t := -math.Pi; t < math.Pi; t += math.Pi / n {
		point := f1(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	dc.SavePNG("hw/hw3/n2/bezier_2nd_order.png")

	// Transformed
	dc = matrixes.NewDrawingContext(width*3, height*2, hasNumbers)
	movement := matrixes.GetMovementMatrix(ax, ay)
	shear := matrixes.GetShearMatrixOx(2)
	rotation := matrixes.GetRotationMatrix(math.Pi / 4)

	trans_points, _ := matrixes.Multiply(movement, points_2nd_order)
	trans_points, _ = matrixes.Multiply(shear, trans_points)
	trans_points, _ = matrixes.Multiply(rotation, trans_points)

	matrixes.DrawPoints(dc, 19, trans_points)
	f2 := matrixes.QuadraticBezierFunction(trans_points)
	for t := -math.Pi; t < math.Pi; t += math.Pi / n {
		point := f2(t)
		dc.DrawPoint(point.Data[0][0]/3, point.Data[1][0], 1)
		dc.Stroke()
	}

	dc.SavePNG("hw/hw3/n2/bezier_2nd_order_transformed.png")
}

func Hw3_n2_b() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	matrixes.DrawPoints(dc, 19, points_3rd_order)
	f1 := matrixes.CubicBezierFunction(points_3rd_order)

	var n float64 = 500

	for t := -math.Pi; t < math.Pi; t += math.Pi / n {
		point := f1(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	dc.SavePNG("hw/hw3/n2/bezier_3rd_order.png")
}

func Hw3_n2_с() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	points1 := matrixes.CreateMatrix([][]float64{
		{-100, 200, 300, 200},
		{-100, 0, 200, 200},
		{1, 1, 1, 1}})

	points2 := matrixes.CreateMatrix([][]float64{
		{200, 100, -100, 100},
		{200, 200, -100, -150},
		{1, 1, 1, 1}})

	points3 := matrixes.CreateMatrix([][]float64{
		{-100, -400, -200, 200},
		{-100, -200, -400, -300},
		{1, 1, 1, 1}})

	points4 := matrixes.CreateMatrix([][]float64{
		{200, 400, 100},
		{-300, -200, -150},
		{1, 1, 1}})

	matrixes.DrawPoints(dc, 10, points1)
	matrixes.DrawPoints(dc, 10, points2)
	matrixes.DrawPoints(dc, 10, points3)
	matrixes.DrawPoints(dc, 10, points4)

	var n float64 = 500

	f := matrixes.CubicBezierFunction(points1)
	for t := 0.0; t < 1; t += 1.0 / n {
		point := f(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	f = matrixes.CubicBezierFunction(points2)
	for t := 0.0; t < 1; t += 1.0 / n {
		point := f(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	f = matrixes.CubicBezierFunction(points3)
	for t := 0.0; t < 1; t += 1.0 / n {
		point := f(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	f = matrixes.QuadraticBezierFunction(points4)
	for t := 0.0; t < 1; t += 1.0 / n {
		point := f(t)
		dc.DrawPoint(point.Data[0][0], point.Data[1][0], 1)
		dc.Stroke()
	}

	dc.SavePNG("hw/hw3/n2/spline.png")
}
