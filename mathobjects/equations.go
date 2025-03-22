package mathobjects

import "math"

func GetLineEquationParametric(x1, y1, x2, y2 float64) func(t float64) (float64, float64) {
	// Вектор направления (dx, dy)
	dx := x2 - x1
	dy := y2 - y1

	// Функция, вычисляющая координаты точки на прямой для заданного t
	return func(t float64) (float64, float64) {
		x := x1 + t*dx
		y := y1 + t*dy
		return x, y
	}
}

func GetEllipse(a, b, t float64) (float64, float64) {
	x := a * math.Cos(t)
	y := b * math.Sin(t)
	return x, y
}

func GetEllipseFunctions(a, b float64) (func(float64) float64, func(float64) float64) {
	f1 := func(t float64) float64 { return a * math.Cos(t) }
	f2 := func(t float64) float64 { return b * math.Sin(t) }
	return f1, f2
}

func GetHyperbolic(a, b, t float64) (float64, float64) {
	x := a * math.Cosh(t)
	y := b * math.Sinh(t)
	return x, y
}

func GetSpiralArchimedFunctions(v, c, w float64) (func(float64) float64, func(float64) float64) {
	f1 := func(t float64) float64 { return (v*t + c) * math.Cos(w*t) }
	f2 := func(t float64) float64 { return (v*t + c) * math.Sin(w*t) }
	return f1, f2
}

func GetCicloidFunctions(a float64) (func(float64) float64, func(float64) float64) {
	f1 := func(t float64) float64 { return a * (t - math.Sin(t)) }
	f2 := func(t float64) float64 { return a * (1 - math.Cos(t)) }
	return f1, f2
}

func GetHypoCicloidFunctions(R, k float64) (func(float64) float64, func(float64) float64) {
	r := R / k
	f1 := func(t float64) float64 { return r * (k - 1) * (math.Cos(t) + math.Cos((k-1)*t)/(k-1)) }
	f2 := func(t float64) float64 { return r * (k - 1) * (math.Sin(t) - math.Sin((k-1)*t)/(k-1)) }
	return f1, f2
}

func deriviate(f func(float64) float64, x, h float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

// Численное вычисление второй производной
func secondDerivative(f func(float64) float64, t, h float64) float64 {
	return (f(t+h) - 2*f(t) + f(t-h)) / (h * h)
}

// Функция вычисления кривизны
func GetCurvature(xFunc, yFunc func(float64) float64, t float64) float64 {
	h := 1e-5 // Маленький шаг для численного дифференцирования

	// Производные
	x1 := deriviate(xFunc, t, h)
	y1 := deriviate(yFunc, t, h)
	x2 := secondDerivative(xFunc, t, h)
	y2 := secondDerivative(yFunc, t, h)

	// Формула кривизны
	numerator := math.Abs(x1*y2 - y1*x2)
	denominator := math.Pow(x1*x1+y1*y1, 1.5)

	if denominator == 0 {
		return 0 // Кривизна не определена (обычно это прямая)
	}

	return numerator / denominator
}

func GetTangetVectorParam(f1, f2 func(float64) float64, t float64) (float64, float64) {
	tx := deriviate(f1, t, 1e-5)
	ty := deriviate(f2, t, 1e-5)
	return tx, ty
}

func GetNormalVectorParam(xFunc, yFunc func(float64) float64, t float64) (float64, float64) {
	h := 1e-5
	x1 := deriviate(xFunc, t, h)
	y1 := deriviate(yFunc, t, h)
	return -y1, x1
}

func getTangentNormalFunctions(xFunc, yFunc func(float64) float64, t float64) (func(float64) float64, func(float64) float64, func(float64) float64, func(float64) float64) {
	h := 1e-5

	// Точка (x0, y0)
	x0 := xFunc(t)
	y0 := yFunc(t)

	// Производные
	x1 := deriviate(xFunc, t, h)
	y1 := deriviate(yFunc, t, h)

	// Вектор нормали
	nx, ny := GetNormalVectorParam(xFunc, yFunc, t)

	// Функции для касательной
	tangentX := func(s float64) float64 {
		return x0 + s*x1
	}
	tangentY := func(s float64) float64 {
		return y0 + s*y1
	}

	// Функции для нормали
	normalX := func(s float64) float64 {
		return x0 + s*nx
	}
	normalY := func(s float64) float64 {
		return y0 + s*ny
	}

	return tangentX, tangentY, normalX, normalY
}

func GetEvolute(xFunc, yFunc func(float64) float64, t float64) (float64, float64) {
	h := 1e-5 // Малый шаг для численного дифференцирования

	// Первая производная
	x1 := deriviate(xFunc, t, h)
	y1 := deriviate(yFunc, t, h)

	// Вторая производная
	x2 := deriviate(func(t float64) float64 { return deriviate(xFunc, t, h) }, t, h)
	y2 := deriviate(func(t float64) float64 { return deriviate(yFunc, t, h) }, t, h)

	// Вычисление знаменателя (избегаем деления на ноль)
	denominator := x1*y2 - y1*x2
	if math.Abs(denominator) < 1e-9 {
		return math.NaN(), math.NaN()
	}

	// Координаты эволюты (центр кривизны)
	X := xFunc(t) - (x1*x1+y1*y1)*y1/denominator
	Y := yFunc(t) + (x1*x1+y1*y1)*x1/denominator

	return X, Y
}
