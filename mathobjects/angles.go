package mathobjects

import "math"

func AngleDiff(a, b float64) float64 {
	diff := a - b
	if diff > math.Pi {
		diff -= 2 * math.Pi
	} else if diff < -math.Pi {
		diff += 2 * math.Pi
	}
	return diff
}

func PolarAngle(start1, end1, start2, end2 [2]float64) float64 {
	// Вычисляем векторы
	vec1 := [2]float64{end1[0] - start1[0], end1[1] - start1[1]}
	vec2 := [2]float64{end2[0] - start2[0], end2[1] - start2[1]}

	// Вычисляем углы каждого вектора относительно оси X
	angle1 := math.Atan2(vec1[1], vec1[0])
	angle2 := math.Atan2(vec2[1], vec2[0])

	// Вычисляем разницу углов и нормализуем в [0, 2π)
	angle := AngleDiff(angle1, angle2)

	return angle
}
