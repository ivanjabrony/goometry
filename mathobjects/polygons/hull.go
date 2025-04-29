package polygons

import (
	"math"
	"sort"

	"github.com/golang-collections/collections/stack"
)

func BuildConvexHullJarvis(points [][2]float64) [][2]float64 {
	n := len(points)
	if n < 3 {
		return nil // выпуклая оболочка не существует
	}

	// Находим самую левую точку
	leftmost := 0
	for i := 1; i < n; i++ {
		if points[i][0] < points[leftmost][0] {
			leftmost = i
		} else if points[i][0] == points[leftmost][0] && points[i][1] < points[leftmost][1] {
			leftmost = i
		}
	}

	hull := make([][2]float64, 0)
	p := leftmost
	var q int

	for {
		// Добавляем текущую точку в оболочку
		hull = append(hull, points[p])

		q = (p + 1) % n
		for i := 0; i < n; i++ {
			// Если точка i лежит "левее" pq, обновляем q
			if slice_orientation(points[p], points[i], points[q]) == 2 {
				q = i
			} else if slice_orientation(points[p], points[i], points[q]) == 0 { // Если точки коллинеарны, выбираем самую дальнюю
				if distSq(points[p], points[i]) > distSq(points[p], points[q]) {
					q = i
				}
			}
		}

		p = q

		// Завершаем, когда возвращаемся к начальной точке
		if p == leftmost {
			break
		}
	}

	return hull
}

func BuildConvexHullGraham(points [][2]float64) [][2]float64 {
	if len(points) < 3 {
		return nil
	}

	// Находим точку с минимальной y-координатой
	minY := points[0]
	minIndex := 0
	for i, p := range points {
		if p[1] < minY[1] || (p[1] == minY[1] && p[0] < minY[0]) {
			minY = p
			minIndex = i
		}
	}

	// Перемещаем стартовую точку в начало
	points[0], points[minIndex] = points[minIndex], points[0]
	p0 := points[0]
	points = points[1:]

	// Сортируем точки по полярному углу
	sort.Slice(points, func(i, j int) bool {
		angleI := math.Atan2(points[i][1]-p0[1], points[i][0]-p0[0])
		angleJ := math.Atan2(points[j][1]-p0[1], points[j][0]-p0[0])
		if angleI == angleJ {
			distI := math.Pow(points[i][0]-p0[0], 2) + math.Pow(points[i][1]-p0[1], 2)
			distJ := math.Pow(points[j][0]-p0[0], 2) + math.Pow(points[j][1]-p0[1], 2)
			return distI < distJ
		}
		return angleI < angleJ
	})

	// Инициализируем стек
	st := stack.New()
	st.Push(p0)
	if len(points) > 0 {
		st.Push(points[0])
	}
	if len(points) > 1 {
		st.Push(points[1])
	}

	// Обрабатываем остальные точки
	for i := 2; i < len(points); i++ {
		for st.Len() >= 2 {
			// Получаем три верхние точки
			c := points[i]
			b := st.Pop().([2]float64)
			a := st.Peek().([2]float64)

			// Векторы ab и ac
			ab := [2]float64{b[0] - a[0], b[1] - a[1]}
			ac := [2]float64{c[0] - a[0], c[1] - a[1]}

			// Векторное произведение
			cross := ab[0]*ac[1] - ab[1]*ac[0]

			if cross > 0 {
				st.Push(b) // Возвращаем точку обратно
				break
			}
			// Иначе точка b удаляется из стека
		}
		st.Push(points[i])
	}

	// Преобразуем стек в слайс
	result := make([][2]float64, st.Len())
	for i := st.Len() - 1; i >= 0; i-- {
		result[i] = st.Pop().([2]float64)
	}

	return result
}

// IntersectConvexHulls находит пересечение двух выпуклых многоугольников
func IntersectConvexHulls(hullA, hullB [][2]float64) [][2]float64 {
	// Проверка на минимальное количество точек
	if len(hullA) < 3 || len(hullB) < 3 {
		return nil
	}

	// Инициализируем результат как первый многоугольник
	result := make([][2]float64, len(hullA))
	copy(result, hullA)

	// Последовательно отсекаем результат гранями второго многоугольника
	for i := 0; i < len(hullB); i++ {
		next := (i + 1) % len(hullB)
		edgeStart := hullB[i]
		edgeEnd := hullB[next]

		result = clipPolygon(result, edgeStart, edgeEnd)
		if len(result) < 3 {
			return nil // Пересечение пусто
		}
	}

	return result
}

func clipPolygon(polygon [][2]float64, edgeStart, edgeEnd [2]float64) [][2]float64 {
	var output [][2]float64
	n := len(polygon)
	if n == 0 {
		return output
	}

	// Вычисляем нормаль к ребру, указывающую ВНУТРЬ второго многоугольника
	edgeVecX := edgeEnd[0] - edgeStart[0]
	edgeVecY := edgeEnd[1] - edgeStart[1]
	normalX := -edgeVecY // Изменено направление нормали!
	normalY := edgeVecX  // по сравнению с предыдущим кодом

	for i := 0; i < n; i++ {
		current := polygon[i]
		next := polygon[(i+1)%n]

		// Векторы от начала ребра к точкам
		vecCurX := current[0] - edgeStart[0]
		vecCurY := current[1] - edgeStart[1]
		vecNextX := next[0] - edgeStart[0]
		vecNextY := next[1] - edgeStart[1]

		// Скалярные произведения с нормалью
		dotCur := normalX*vecCurX + normalY*vecCurY
		dotNext := normalX*vecNextX + normalY*vecNextY

		// Точка внутри полуплоскости (включая границу)
		if dotCur >= 0 { // Изменено условие с <= на >=
			output = append(output, current)
		}

		// Пересечение ребра многоугольника с границей полуплоскости
		if (dotCur > 0 && dotNext < 0) || (dotCur < 0 && dotNext > 0) {
			t := dotCur / (dotCur - dotNext)
			intersectX := current[0] + t*(next[0]-current[0])
			intersectY := current[1] + t*(next[1]-current[1])
			output = append(output, [2]float64{intersectX, intersectY})
		}
	}

	return output
}
