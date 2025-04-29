package distance

import (
	"math"
	"sort"
)

func ClosestPairBruteForce(points [][2]float64) ([2]float64, [2]float64, float64) {
	minDist := math.MaxFloat64
	var p1, p2 [2]float64

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := math.Abs(math.Pow(points[i][0]-points[j][0], 2) +
				math.Pow(points[i][1]-points[j][1], 2))
			if dist < minDist {
				minDist = dist
				p1, p2 = points[i], points[j]
			}
		}
	}
	return p1, p2, math.Sqrt(minDist)
}

func ClosestPairDAC(points [][2]float64) ([2]float64, [2]float64, float64) {
	// Предварительная сортировка по X координате
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	return closestPairRec(points)
}

func closestPairRec(points [][2]float64) ([2]float64, [2]float64, float64) {
	if len(points) <= 3 {
		return ClosestPairBruteForce(points)
	}

	mid := len(points) / 2
	leftPair, rightPair := points[:mid], points[mid:]

	l1, l2, dl := closestPairRec(leftPair)
	r1, r2, dr := closestPairRec(rightPair)

	var minPair1, minPair2 [2]float64
	minDist := math.Min(dl, dr)
	if dl < dr {
		minPair1, minPair2 = l1, l2
	} else {
		minPair1, minPair2 = r1, r2
	}

	// Проверка полосы вокруг разделяющей линии
	midX := points[mid][0]
	strip := make([][2]float64, 0)
	for _, p := range points {
		if math.Sqrt(math.Pow(p[0]-midX, 2)) < minDist {
			strip = append(strip, p)
		}
	}

	// Сортировка полосы по Y координате
	sort.Slice(strip, func(i, j int) bool {
		return strip[i][1] < strip[j][1]
	})

	// Проверка соседей в полосе
	for i := 0; i < len(strip); i++ {
		for j := i + 1; j < len(strip) && math.Sqrt(math.Pow(strip[j][1]-strip[i][1], 2)) < minDist; j++ {
			dist := math.Pow(strip[i][0]-strip[j][0], 2) +
				math.Pow(strip[i][1]-strip[j][1], 2)
			dist = math.Sqrt(dist)
			if dist < minDist {
				minDist = dist
				minPair1, minPair2 = strip[i], strip[j]
			}
		}
	}

	return minPair1, minPair2, minDist
}
