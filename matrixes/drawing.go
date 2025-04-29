package matrixes

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fogleman/gg"
)

func DrawTextAtPoint(dc *gg.Context, text string, x, y float64) {
	dc.Push()
	dc.Scale(-1, 1)
	dc.Rotate(-math.Pi)
	dc.DrawString(text, x, -y)
	dc.Stroke()
	dc.Pop()
}

func WritePointCoordinates(dc *gg.Context, x, y float64) {
	DrawTextAtPoint(dc, fmt.Sprintf("(%v, %v)", x, y), x, y)
}

func NewDrawingContext(width, height int, hasNumbers bool) *gg.Context {
	dc := gg.NewContext(width, height)
	centerX, centerY := 0, 0

	dc.InvertY()
	dc.Translate(float64(width/2), float64(height/2))
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0) // Чёрный цвет
	dc.SetLineWidth(2)

	//оси
	dc.DrawLine(-float64(width), float64(centerY), float64(width), float64(centerY))
	dc.Stroke()

	dc.DrawLine(float64(centerX), -float64(height), float64(centerX), float64(height))
	dc.Stroke()

	for i := 0; i <= height/25; i++ {
		if i == 0 {
			if hasNumbers {
				DrawTextAtPoint(dc, "0", 1, 0)
			}
			continue // Пропускаем центр
		}

		// Метки по оси X
		x := float64(centerX + i*25)
		dc.DrawLine(x, float64(centerY-5), x, float64(centerY+5))
		dc.Stroke()

		dc.DrawLine(-x, float64(centerY-5), -x, float64(centerY+5))
		dc.Stroke()

		if hasNumbers {
			text := strconv.Itoa(int(x))
			DrawTextAtPoint(dc, text, x, 0)
			text = strconv.Itoa(int(-x))
			DrawTextAtPoint(dc, text, -x, 0)
		}

		// Метки по оси Y
		y := float64(centerY + i*25)
		dc.DrawLine(float64(centerX-5), y, float64(centerX+5), y)
		dc.Stroke()

		dc.DrawLine(float64(centerX-5), -y, float64(centerX+5), -y)
		dc.Stroke()

		if hasNumbers {
			text := strconv.Itoa(int(y))
			DrawTextAtPoint(dc, text, 1, y)
			text = strconv.Itoa(int(-y))
			DrawTextAtPoint(dc, text, 1, -y)
		}
	}

	return dc
}

func DrawPolygon(ctx *gg.Context, points []Matrix) {
	ctx.Push()
	for i := range points[:len(points)-1] {
		x1 := points[i].Data[0][0]
		y1 := points[i].Data[1][0]
		x2 := points[i+1].Data[0][0]
		y2 := points[i+1].Data[1][0]
		ctx.DrawLine(x1, y1, x2, y2)
		ctx.Stroke()
	}
	ctx.DrawLine(
		points[0].Data[0][0],
		points[0].Data[1][0],
		points[len(points)-1].Data[0][0],
		points[len(points)-1].Data[1][0])
	ctx.Stroke()
	ctx.Pop()
}

func DrawPolygonFromMatrix(ctx *gg.Context, points Matrix) {
	ctx.Push()
	n := len(points.Data[0])
	for i := range n {
		x := points.Data[0][i]
		y := points.Data[1][i]
		k := points.Data[2][i]
		x = x / k
		y = y / k

		ctx.DrawPoint(x, y, 5)
		ctx.Stroke()
	}

	for i := range n {
		if i == 0 {
			continue
		}

		x1 := points.Data[0][i-1]
		y1 := points.Data[1][i-1]

		x2 := points.Data[0][i]
		y2 := points.Data[1][i]

		k1 := points.Data[2][i-1]
		k2 := points.Data[2][i]

		x1 = x1 / k1
		y1 = y1 / k1

		x2 = x2 / k2
		y2 = y2 / k2

		ctx.DrawLine(x1, y1, x2, y2)
		ctx.Stroke()

		if i == n-1 {
			x1 := points.Data[0][0]
			y1 := points.Data[1][0]

			x2 := points.Data[0][i]
			y2 := points.Data[1][i]

			ctx.DrawLine(x1, y1, x2, y2)
			ctx.Stroke()
		}
	}
	ctx.Pop()
}

func DrawLineStupid(ctx *gg.Context, xa, ya, xb, yb int, isInverted bool) {
	k := float64(yb-ya) / float64(xb-xa)
	b := float64(ya) - k*float64(xa)
	ctx.Push()

	if xa > xb {
		xa, xb = xb, xa
	}
	for xi := xa; xi < xb; xi++ {
		yi := int(k*float64(xi) + b)
		if isInverted {
			ctx.SetPixel(ctx.Width()/2+xi, ctx.Height()/2-yi)
		} else {
			ctx.SetPixel(ctx.Width()/2+xi, ctx.Height()/2+yi)
		}
		ctx.Stroke()
	}
	ctx.Pop()
}

func DrawLineBerz(ctx *gg.Context, x0, y0, x1, y1 float64, isInverted bool) {
	x0, y0, x1, y1 = math.Round(x0), math.Round(y0), math.Round(x1), math.Round(y1)
	w_offset := ctx.Width() / 2
	h_offset := ctx.Height() / 2

	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)
	er := dx - dy
	var sx, sy float64

	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}

	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}

	for {
		if isInverted {
			ctx.SetPixel(w_offset+int(x0), h_offset-int(y0))
		} else {
			ctx.SetPixel(w_offset+int(x0), h_offset+int(y0))
		}

		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * er
		if e2 >= -dy {
			er -= dy
			x0 += sx
		}
		if e2 < dx {
			er += dx
			y0 += sy
		}
	}
}

func DrawPoints(ctx *gg.Context, thickness int, points Matrix) {
	ctx.Push()
	n := len(points.Data[0])
	for i := range n {
		x := points.Data[0][i]
		y := points.Data[1][i]
		k := points.Data[2][i]
		x = x / k
		y = y / k

		ctx.DrawPoint(x, y, 5)
		ctx.Stroke()
	}
	ctx.Pop()
}

func DrawFunc(ctx *gg.Context, f func(float64) float64, t1, t2 float64) {
	ctx.Push()
	for t := float64(t1); t <= t2; t += (t2 - t1) / 10 {
		ctx.DrawPoint(t, f(t1), 3)

		ctx.Stroke()
	}
	ctx.Pop()
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
