package hw1_n2

import (
	"fmt"
	"gogeom/matrixes"
	"image/color"
	"math"
)

const (
	width, height = 1500, 1500 // Размер изображения
	hasNumbers    = true
)

var (
	length = 100.0
	k      = matrixes.CreateMatrix([][]float64{{0}, {0}, {1}})
	l      = matrixes.CreateMatrix([][]float64{{length}, {0}, {1}})
	m      = matrixes.CreateMatrix([][]float64{{length}, {length}, {1}})
	n      = matrixes.CreateMatrix([][]float64{{0}, {length}, {1}})

	trans = matrixes.CreateMatrix([][]float64{
		{2, 0, 3 * length},
		{2 * math.Sqrt(3) / 3, 2, 3 * length},
		{0, 0, 1}})

	trans_inv = matrixes.CreateMatrix([][]float64{
		{0.5, 0, -3 * length / 2},
		{-0.5 / math.Sqrt(3), 0.5, length * (math.Sqrt(3) - 3) / 2},
		{0, 0, 1}})

	rectangle, _ = k.ConcatHorizontal(l, m, n)
)

func Hw1_n2_a() {
	fmt.Printf("square:\n%v\n", rectangle)

	shear_matrix := matrixes.GetShearMatrixOy(1 / math.Sqrt(3))
	stretch_matrix := matrixes.GetStretchMatrix(2, 2)
	move_matrix := matrixes.GetMovementMatrix(300, 300)

	trans_composition, _ := matrixes.Multiply(stretch_matrix, shear_matrix)
	trans_composition, _ = matrixes.Multiply(move_matrix, trans_composition)

	fmt.Printf("matrix F_composition:\n%v\n", trans_composition)
	fmt.Printf("Equation is: \n { x' = 2x + 3 * %v\n { y' = 2x/sqrt(3) + 2y + 3 * %v\n\n", length, length)
}

func Hw1_n2_b() {
	fmt.Printf("matrix F:\n%v\n", trans)
	fmt.Printf("matrix F^-1\n%v\n", trans_inv)
}

func Hw1_n2_c() {
	rectangle, _ := k.ConcatHorizontal(l, m, n)

	affected_rectangle, _ := matrixes.Multiply(trans, rectangle)

	reverted_rectangle, _ := matrixes.Multiply(trans_inv, affected_rectangle)

	fmt.Printf("matrix F(KLMN):\n%v\n", affected_rectangle)
	fmt.Printf("matrix F^-1(ABCD)\n%v\n", reverted_rectangle)
}

func Hw1_n2_d() {
	dc := matrixes.NewDrawingContext(width, height, false)

	dc.Push()
	//KLMN
	matrixes.DrawPolygon(dc, []matrixes.Matrix{k, l, m, n})

	matrixes.DrawTextAtPoint(dc, "K", k.Data[0][0], k.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "L", l.Data[0][0], l.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "M", m.Data[0][0], m.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "N", n.Data[0][0], n.Data[1][0])

	a, _ := matrixes.Multiply(trans, k)
	b, _ := matrixes.Multiply(trans, l)
	c, _ := matrixes.Multiply(trans, m)
	d, _ := matrixes.Multiply(trans, n)

	//F(KLMN) = ABCD
	matrixes.DrawPolygon(dc, []matrixes.Matrix{a, b, c, d})

	matrixes.DrawTextAtPoint(dc, "A", a.Data[0][0], a.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "B", b.Data[0][0], b.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "C", c.Data[0][0], c.Data[1][0])
	matrixes.DrawTextAtPoint(dc, "D", d.Data[0][0], d.Data[1][0])

	dc.SavePNG("hw/hw1/n2/task1_d1_rotate.png")

	matrixes.DrawPolygon(dc, []matrixes.Matrix{a, b, c, d})

	a, _ = matrixes.Multiply(trans_inv, a)
	b, _ = matrixes.Multiply(trans_inv, b)
	c, _ = matrixes.Multiply(trans_inv, c)
	d, _ = matrixes.Multiply(trans_inv, d)

	//после обратного преобразования
	//KLMN = F^-1(ABCD)
	dc.SetColor(color.RGBA{100, 0, 100, 100})
	matrixes.DrawPolygon(dc, []matrixes.Matrix{a, b, c, d})
	dc.SavePNG("hw/hw1/n2/task1_d2_rotate.png")
}

func Hw1_n2_e() {

	rectangle, _ := k.ConcatHorizontal(l, m, n)

	affected_rectangle, _ := matrixes.Multiply(trans, rectangle)
	reverted_rectangle, _ := matrixes.Multiply(trans_inv, affected_rectangle)

	println("Comparison of matrixes")
	fmt.Printf("matrix KLMN:\n%v\n", rectangle)
	fmt.Printf("matrix F(KLMN):\n%v\n", affected_rectangle)
	fmt.Printf("matrix F^-1(ABCD)\n%v\n", reverted_rectangle)
}
