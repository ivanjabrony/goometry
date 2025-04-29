package polygons

import (
	"gogeom/mathobjects"
	"math"
)

// Проверка, находится ли точка внутри многоугольника по Winding Number
func IsPointInsidePolygonWinding(polygon [][2]float64, point [2]float64) bool {
	windingNumber := 0.0
	px := point[0]
	py := point[1]

	for i := 0; i < len(polygon); i++ {
		x1, y1 := polygon[i][0], polygon[i][1]
		x2, y2 := polygon[(i+1)%len(polygon)][0], polygon[(i+1)%len(polygon)][1]

		theta1 := math.Atan2(y1-py, x1-px)
		theta2 := math.Atan2(y2-py, x2-px)

		windingNumber += mathobjects.AngleDiff(theta2, theta1)
	}

	// Если сумма углов ≈ 2π или -2π, точка внутри
	return math.Abs(windingNumber) > 1e-6
}

func FindPerimeter(polygon [][2]float64) float64 {
	n := len(polygon)
	ans := mathobjects.FindDistance(polygon[n-1], polygon[0])

	for i := range polygon {
		if i != 0 {
			ans += mathobjects.FindDistance(polygon[i], polygon[i-1])
		}
	}

	return ans
}

func FindPolygonArea(points [][2]float64) float64 {
	n := len(points)
	if n < 3 {
		return 0.0 // Не многоугольник
	}

	var area float64

	// Применяем формулу шнурка
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i][0] * points[j][1]
		area -= points[i][1] * points[j][0]
	}

	return math.Abs(area) / 2.0
}

func IsPointInsidePolygon(polygon [][2]float64, point [2]float64) bool {
	n := len(polygon)
	if n < 3 {
		return false // Не многоугольник
	}

	intersections := 0
	const infinity = 1e10

	// Создаем луч (горизонтальный вправо)
	rayEnd := [2]float64{infinity, point[1]}

	for i := 0; i < n; i++ {
		j := (i + 1) % n
		edgeStart := polygon[i]
		edgeEnd := polygon[j]

		if edgeStart[1] == edgeEnd[1] {
			continue
		}

		// Проверяем пересечение луча с текущим ребром
		if doIntersect(point, rayEnd, edgeStart, edgeEnd) {
			// Если точка лежит на ребре
			if onSegment(edgeStart, point, edgeEnd) {
				return true
			}

			// Учитываем особые случаи для горизонтальных ребер и вершин
			if edgeStart[1] == point[1] && edgeEnd[1] >= point[1] {
				continue
			}
			if edgeEnd[1] == point[1] && edgeStart[1] >= point[1] {
				continue
			}

			intersections++
		}
	}

	// Если число пересечений нечетное - точка внутри
	return intersections%2 == 1
}

// doIntersect проверяет пересекаются ли два отрезка
func doIntersect(p1, p2, p3, p4 [2]float64) bool {
	// Ориентации для всех комбинаций точек
	o1 := slice_orientation(p1, p2, p3)
	o2 := slice_orientation(p1, p2, p4)
	o3 := slice_orientation(p3, p4, p1)
	o4 := slice_orientation(p3, p4, p2)

	// Общий случай пересечения
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Специальные случаи коллинеарности
	if o1 == 0 && onSegment(p1, p3, p2) {
		return true
	}
	if o2 == 0 && onSegment(p1, p4, p2) {
		return true
	}
	if o3 == 0 && onSegment(p3, p1, p4) {
		return true
	}
	if o4 == 0 && onSegment(p3, p2, p4) {
		return true
	}

	return false
}
