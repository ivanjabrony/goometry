package mathobjects

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type point struct {
	X, Y float64
}

// segment представляет отрезок с начальной и конечной точками
type segment struct {
	Start, End point
}

// event представляет событие для заметающей прямой
type event struct {
	X    float64
	Type string // "start" или "end"
	Seg  segment
}

func NewSegment(x1, y1, x2, y2 float64) segment {
	return segment{point{x1, y1}, point{x2, y2}}
}

func IsIntersectSimple(x1, y1, x2, y2, x3, y3, x4, y4 float64) (*float64, *float64, bool) {
	xp, yp, isIntersect := findLineIntersection(x1, y1, x2, y2, x3, y3, x4, y4)
	if !isIntersect {
		return nil, nil, false
	}

	xa, xb := min(x1, x2), max(x1, x2)
	xc, xd := min(x3, x4), max(x3, x4)
	ya, yb := min(y1, y2), max(y1, y2)
	yc, yd := min(y3, y4), max(y3, y4)

	isIntersect = (xa <= xp && xp <= xb) &&
		(xc <= xp && xp <= xd) &&
		(ya <= yp && yp <= yb) &&
		(yc <= yp && yp <= yd)

	if isIntersect {
		return &xp, &yp, true
	}

	return nil, nil, false
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
func onsegment(p1x, p1y, p2x, p2y, p3x, p3y float64) bool {
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
	if o1 == 0 && onsegment(x1, y1, x2, y2, x3, y3) {
		return true
	}
	if o2 == 0 && onsegment(x1, y1, x2, y2, x4, y4) {
		return true
	}
	if o3 == 0 && onsegment(x3, y3, x4, y4, x1, y1) {
		return true
	}
	if o4 == 0 && onsegment(x3, y3, x4, y4, x2, y2) {
		return true
	}

	return false
}

func doIntersectSegments(s1, s2 segment) bool {
	p1, q1 := s1.Start, s1.End
	p2, q2 := s2.Start, s2.End

	// Находим ориентации для всех комбинаций точек
	o1 := orientation(p1.X, p1.Y, q1.X, q1.Y, p2.X, p2.Y)
	o2 := orientation(p1.X, p1.Y, q1.X, q1.Y, q2.X, q2.Y)
	o3 := orientation(p2.X, p2.Y, q2.X, q2.Y, p1.X, p1.Y)
	o4 := orientation(p2.X, p2.Y, q2.X, q2.Y, q1.X, q1.Y)

	// Общий случай пересечения
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Специальные случаи (коллинеарность)
	if o1 == 0 && onsegment(p1.X, p1.Y, p2.X, p2.Y, q1.X, q1.Y) {
		return true
	}
	if o2 == 0 && onsegment(p1.X, p1.Y, q2.X, q2.Y, q1.X, q1.Y) {
		return true
	}
	if o3 == 0 && onsegment(p2.X, p2.Y, p1.X, p1.Y, q2.X, q2.Y) {
		return true
	}
	if o4 == 0 && onsegment(p2.X, p2.Y, q1.X, q1.Y, q2.X, q2.Y) {
		return true
	}

	return false
}

func findIntersectionsSegments(segments []segment) [][2]segment {
	// Создаем список событий
	events := []event{}
	for _, seg := range segments {
		start, end := seg.Start, seg.End
		// Убедимся, что start.X <= end.X
		if start.X > end.X {
			start, end = end, start
		}
		events = append(events, event{start.X, "start", seg})
		events = append(events, event{end.X, "end", seg})
	}

	// Сортируем события по X
	sort.Slice(events, func(i, j int) bool {
		return events[i].X < events[j].X
	})

	// Активные отрезки (используем красно-черное дерево)
	tree := redblacktree.NewWith(func(a, b interface{}) int {
		segA := a.(segment)
		segB := b.(segment)
		// Сравниваем отрезки по Y-координате в текущей позиции заметающей прямой
		yA := segA.Start.Y + (segA.End.Y-segA.Start.Y)*(events[0].X-segA.Start.X)/(segA.End.X-segA.Start.X)
		yB := segB.Start.Y + (segB.End.Y-segB.Start.Y)*(events[0].X-segB.Start.X)/(segB.End.X-segB.Start.X)
		if yA < yB {
			return -1
		} else if yA > yB {
			return 1
		}
		return 0
	})

	intersections := [][2]segment{}

	for _, event := range events {
		_, typ, seg := event.X, event.Type, event.Seg
		if typ == "start" {
			// Вставляем отрезок в дерево
			tree.Put(seg, true)
			// Получаем соседей
			lower, _ := tree.Floor(seg)
			upper, _ := tree.Ceiling(seg)
			// Проверяем пересечение с соседями
			if lower != nil {
				if doIntersectSegments(seg, lower.Key.(segment)) {
					intersections = append(intersections, [2]segment{seg, lower.Key.(segment)})
				}
			}
			if upper != nil {
				if doIntersectSegments(seg, upper.Key.(segment)) {
					intersections = append(intersections, [2]segment{seg, upper.Key.(segment)})
				}
			}
		} else {
			// Удаляем отрезок из дерева
			tree.Remove(seg)
		}
	}

	return intersections
}

func FindIntersectionsOfNLines(lines [][4]float64) [][2][4]float64 {
	segments := make([]segment, len(lines))

	for i, v := range lines {
		segments[i] = NewSegment(v[0], v[1], v[2], v[3])
	}

	crossedSegments := findIntersectionsSegments(segments)
	ans := make([][2][4]float64, len(crossedSegments))

	for i, v := range crossedSegments {
		ans[i] = [2][4]float64{
			{v[0].Start.X, v[0].Start.Y, v[0].End.X, v[0].End.Y},
			{v[1].Start.X, v[1].Start.Y, v[1].End.X, v[1].End.Y}}
	}

	return ans
}

func FindDistance(p1 [2]float64, p2 [2]float64) float64 {
	return (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
}
