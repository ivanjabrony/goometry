package polygons

import "math"

// Функция для вычисления ориентации трех точек
func slice_orientation(p, q, r [2]float64) int {
	val := (q[1]-p[1])*(r[0]-q[0]) - (q[0]-p[0])*(r[1]-q[1])
	if val == 0 {
		return 0 // коллинеарны
	}
	if val > 0 {
		return 1 // по часовой стрелке
	}
	return 2 // против часовой стрелки
}

// Функция для вычисления квадрата расстояния между двумя точками
func distSq(p1, p2 [2]float64) float64 {
	return math.Pow(p1[0]-p2[0], 2) + math.Pow(p1[1]-p2[1], 2)
}

func onSegment(p1, p, p2 [2]float64) bool {
	// Проверяем, что p лежит на отрезке p1-p2
	cross := (p[0]-p1[0])*(p2[1]-p1[1]) - (p[1]-p1[1])*(p2[0]-p1[0])
	if math.Abs(cross) > 1e-12 { // Если не на линии
		return false
	}
	minX := math.Min(p1[0], p2[0])
	maxX := math.Max(p1[0], p2[0])
	minY := math.Min(p1[1], p2[1])
	maxY := math.Max(p1[1], p2[1])
	return p[0] >= minX-1e-12 && p[0] <= maxX+1e-12 && p[1] >= minY-1e-12 && p[1] <= maxY+1e-12
}

func IsCCW(polygon [][2]float64) bool {
	sum := 0.0
	n := len(polygon)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		sum += (polygon[j][0] - polygon[i][0]) * (polygon[j][1] + polygon[i][1])
	}
	return sum < 0
}

func ReversePolygon(polygon [][2]float64) [][2]float64 {
	n := len(polygon)
	reversed := make([][2]float64, n)
	for i := 0; i < n; i++ {
		reversed[i] = polygon[n-1-i]
	}
	return reversed
}
