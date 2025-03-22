package mathobjects

func IsIntersectSimple(x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	xp, yp, isIntersect := findLineIntersection(x1, y1, x2, y2, x3, y3, x4, y4)
	if !isIntersect {
		return false
	}

	xa, xb := min(x1, x2), max(x1, x2)
	xc, xd := min(x3, x4), max(x3, x4)
	ya, yb := min(y1, y2), max(y1, y2)
	yc, yd := min(y3, y4), max(y3, y4)

	return (xa <= xp && xp <= xb) &&
		(xc <= xp && xp <= xd) &&
		(ya <= yp && yp <= yb) &&
		(yc <= yp && yp <= yd)
}

func findLineIntersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (float64, float64, bool) {
	// Вычисляем коэффициенты системы уравнений
	A1 := x2 - x1
	B1 := -(x4 - x3)
	C1 := x3 - x1

	A2 := y2 - y1
	B2 := -(y4 - y3)
	C2 := y3 - y1

	// Вычисляем детерминант
	D := A1*B2 - A2*B1

	// Проверка на параллельность
	if D == 0 {
		return 0, 0, false // Прямые параллельны или совпадают
	}

	// Находим t1
	t1 := (C1*B2 - C2*B1) / D

	// Вычисляем точку пересечения
	ix := x1 + t1*(x2-x1)
	iy := y1 + t1*(y2-y1)

	return ix, iy, true
}

func orientation(p1x, p1y, p2x, p2y, p3x, p3y float64) float64 {
	return (p2x-p1x)*(p3y-p1y) - (p2y-p1y)*(p3x-p1x)
}

// Проверка, лежит ли точка p3 между p1 и p2
func onSegment(p1x, p1y, p2x, p2y, p3x, p3y float64) bool {
	return p3x >= min(p1x, p2x) && p3x <= max(p1x, p2x) &&
		p3y >= min(p1y, p2y) && p3y <= max(p1y, p2y)
}

func IsIntersectWithPlacement(x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	o1 := orientation(x1, y1, x2, y2, x3, y3)
	o2 := orientation(x1, y1, x2, y2, x4, y4)
	o3 := orientation(x3, y3, x4, y4, x1, y1)
	o4 := orientation(x3, y3, x4, y4, x2, y2)

	// Основное условие пересечения
	if o1*o2 < 0 && o3*o4 < 0 {
		return true
	}

	// Проверка коллинеарных случаев (отрезки лежат на одной прямой)
	if o1 == 0 && onSegment(x1, y1, x2, y2, x3, y3) {
		return true
	}
	if o2 == 0 && onSegment(x1, y1, x2, y2, x4, y4) {
		return true
	}
	if o3 == 0 && onSegment(x3, y3, x4, y4, x1, y1) {
		return true
	}
	if o4 == 0 && onSegment(x3, y3, x4, y4, x2, y2) {
		return true
	}

	return false
}
