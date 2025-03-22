package matrixes

func QuadraticBezierFunction(points Matrix) func(float64) Matrix {

	return func(t float64) Matrix {
		// B(t) = (1-t)^2 * P0 + 2*(1-t)*t * P1 + t^2 * P2
		oneMinusT := 1 - t
		oneMinusTSquared := oneMinusT * oneMinusT
		tSquared := t * t

		x0, y0 := points.Data[0][0], points.Data[1][0]
		x1, y1 := points.Data[0][1], points.Data[1][1]
		x2, y2 := points.Data[0][2], points.Data[1][2]

		x := oneMinusTSquared*x0 + 2*oneMinusT*t*x1 + tSquared*x2
		y := oneMinusTSquared*y0 + 2*oneMinusT*t*y1 + tSquared*y2

		return CreateMatrix([][]float64{{x}, {y}, {1}})
	}
}

func CubicBezierFunction(points Matrix) func(float64) Matrix {

	return func(t float64) Matrix {
		// B(t) = (1-t)^3 * P0 + 3*(1-t)^2*t * P1 + 3*(1-t)*t^2 * P2 + t^3 * P3
		oneMinusT := 1 - t
		oneMinusT2 := oneMinusT * oneMinusT
		oneMinusT3 := oneMinusT2 * oneMinusT
		t2 := t * t
		t3 := t2 * t

		x0, y0 := points.Data[0][0], points.Data[1][0]
		x1, y1 := points.Data[0][1], points.Data[1][1]
		x2, y2 := points.Data[0][2], points.Data[1][2]
		x3, y3 := points.Data[0][3], points.Data[1][3]

		x := oneMinusT3*x0 + 3*oneMinusT2*t*x1 + 3*oneMinusT*t2*x2 + t3*x3
		y := oneMinusT3*y0 + 3*oneMinusT2*t*y1 + 3*oneMinusT*t2*y2 + t3*y3

		return CreateMatrix([][]float64{{x}, {y}, {1}})
	}
}
