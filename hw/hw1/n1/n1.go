package hw1_n1

//import "gonum.org/v1/gonum/mat"

import (
	"gogeom/mathobjects/distance"
	"gogeom/matrixes"
	"image/color"

	"github.com/fogleman/gg"
)

const (
	x1, y1        = 50.0, 100.0
	x2, y2        = 150.0, 100.0
	x3, y3        = 300.0, 200.0
	width, height = 1000, 1000 // Размер изображения
	hasNumbers    = true
)

func Hw1_n1_a() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})
	matrixes.DrawPolygonFromMatrix(dc, triangle)
	dc.SavePNG("hw/hw1/n1/task1_a.png")
}

func Hw1_n1_b_move() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	matrixes.DrawPolygonFromMatrix(dc, triangle)

	a_vec := matrixes.CreateMatrix([][]float64{{100}, {-100}, {0}})
	movement_matrix := matrixes.GetMovementMatrix(a_vec.Data[0][0], a_vec.Data[1][0])

	moved_triangle, _ := matrixes.Multiply(movement_matrix, triangle)

	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, moved_triangle)

	dc.SavePNG("hw/hw1/n1/task1_b_move.png")
}

func Hw1_n1_b_rotate() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)

	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	centerX := (triangle.Data[0][0] + triangle.Data[0][1] + triangle.Data[0][2]) / 3
	centerY := (triangle.Data[1][0] + triangle.Data[1][1] + triangle.Data[1][2]) / 3

	dc.DrawPoint(centerX, centerY, 3)
	phi := gg.Radians(60)

	matrixes.DrawPolygonFromMatrix(dc, triangle)

	rotation_matrix := matrixes.GetFullRotationMatrix(centerX, centerY, phi)

	rotated_triangle, _ := matrixes.Multiply(rotation_matrix, triangle)

	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, rotated_triangle)

	dc.SavePNG("hw/hw1/n1/task1_b_rotate.png")
}

func Hw1_n1_b_symmetry() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	line := [2][2]float64{{0, 0}, {300, 50}}
	dc.DrawLine(line[0][0], line[0][1], line[1][0], line[1][1])
	dc.Stroke()

	matrixes.DrawPolygonFromMatrix(dc, triangle)

	symmetry_matrix := matrixes.GetFullSymmetryMatrix(line[0][0], line[0][1], line[1][0], line[1][1])
	sym_triangle, _ := matrixes.Multiply(symmetry_matrix, triangle)

	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, sym_triangle)

	dc.SavePNG("hw/hw1/n1/task1_b_symmetry.png")
}

func Hw1_n1_b_homotetia() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	k := 2.0

	matrixes.DrawPolygonFromMatrix(dc, triangle)

	homotetia_matrix := matrixes.GetHomotethetyMatrix(0, 0, k)
	homo_triangle, _ := matrixes.Multiply(homotetia_matrix, triangle)

	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, homo_triangle)

	dc.SavePNG("hw/hw1/n1/task1_b_homotetia.png")
}

func Hw1_n1_b_homotetia2() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	m := 2.0

	matrixes.DrawPolygonFromMatrix(dc, triangle)
	dist1 := distance.CalculateDistance2Dots([2]float64{x1, y1}, [2]float64{x2, y2})
	dist2 := distance.CalculateDistance2Dots([2]float64{x2, y2}, [2]float64{x3, y3})
	dist3 := distance.CalculateDistance2Dots([2]float64{x1, y1}, [2]float64{x3, y3})

	var middlepoint [2]float64
	switch {
	case dist1 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x1 + x2) / 2, (y1 + y2) / 2}
	case dist2 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x2 + x3) / 2, (y2 + y3) / 2}
	case dist3 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x1 + x3) / 2, (y1 + y3) / 2}
	}

	dc.DrawPoint(middlepoint[0], middlepoint[1], 3)
	dc.Stroke()

	homotetia2_matrix := matrixes.GetHomotethetyMatrix(middlepoint[0], middlepoint[1], m)
	homo_triangle2, _ := matrixes.Multiply(homotetia2_matrix, triangle)

	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, homo_triangle2)

	dc.SavePNG("hw/hw1/n1/task1_b_homotetia2.png")
}

func Hw1_n1_с_composition1() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	line := [2][2]float64{{0, 0}, {300, 50}}
	k := 2.0
	dc.DrawLine(line[0][0], line[0][1], line[1][0], line[1][1])
	dc.Stroke()

	matrixes.DrawPolygonFromMatrix(dc, triangle)

	symmetry_matrix := matrixes.GetFullSymmetryMatrix(line[0][0], line[0][1], line[1][0], line[1][1])
	homotetia_matrix := matrixes.GetHomotethetyMatrix(0, 0, k)

	triangle, _ = matrixes.Multiply(symmetry_matrix, triangle)
	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, triangle)

	triangle, _ = matrixes.Multiply(homotetia_matrix, triangle)
	dc.SetColor(color.RGBA{50, 40, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, triangle)

	dc.SavePNG("hw/hw1/n1/task1_с_composition1.png")
}

func Hw1_n1_с_composition2() {
	dc := matrixes.NewDrawingContext(width, height, hasNumbers)
	triangle := matrixes.CreateMatrix([][]float64{
		{x1, x2, x3},
		{y1, y2, y3},
		{1, 1, 1}})

	phi := gg.Radians(180)
	m := 2.0

	matrixes.DrawPolygonFromMatrix(dc, triangle)
	dist1 := distance.CalculateDistance2Dots([2]float64{x1, y1}, [2]float64{x2, y2})
	dist2 := distance.CalculateDistance2Dots([2]float64{x2, y2}, [2]float64{x3, y3})
	dist3 := distance.CalculateDistance2Dots([2]float64{x1, y1}, [2]float64{x3, y3})

	var middlepoint [2]float64
	switch {
	case dist1 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x1 + x2) / 2, (y1 + y2) / 2}
	case dist2 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x2 + x3) / 2, (y2 + y3) / 2}
	case dist3 == min(dist1, dist2, dist3):
		middlepoint = [2]float64{(x1 + x3) / 2, (y1 + y3) / 2}
	}

	dc.DrawPoint(middlepoint[0], middlepoint[1], 3)
	dc.Stroke()

	homotetia_matrix := matrixes.GetHomotethetyMatrix(middlepoint[0], middlepoint[1], m)
	rotation_matrix := matrixes.GetFullRotationMatrix(middlepoint[0], middlepoint[1], phi)

	triangle, _ = matrixes.Multiply(homotetia_matrix, triangle)
	dc.SetColor(color.RGBA{50, 40, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, triangle)

	triangle, _ = matrixes.Multiply(rotation_matrix, triangle)
	dc.SetColor(color.RGBA{100, 0, 0, 100})
	matrixes.DrawPolygonFromMatrix(dc, triangle)

	dc.SavePNG("hw/hw1/n1/task1_с_composition2.png")
}
