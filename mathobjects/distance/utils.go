package distance

import "math"

func CalculateDistance2Dots(a [2]float64, b [2]float64) float64 {
	dist := math.Sqrt((a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1]))
	return dist
}
