package program

import "math"

type MapPointFloat struct {
	X, Y float64
}

func subPathByCircle(path []MapPointFloat, pos MapPointFloat, radius float64) []MapPointFloat {
	if len(path) < 2 {
		return path
	}

	radius2 := radius * radius
	startOutside := dist2(path[0], pos) > radius2

	var result []MapPointFloat
	result = append(result, path[0])

	for i := 0; i < len(path)-1; i++ {
		p1 := path[i]
		p2 := path[i+1]

		// 检查线段是否和圆相交
		if intersect, interPt := lineCircleIntersection(p1, p2, pos, radius); intersect {
			// 保留到交点为止
			if startOutside {
				result = append(result, interPt)
				return result
			} else {
				// 丢掉到交点之前，保留之后
				result = []MapPointFloat{interPt}
				result = append(result, path[i+1:]...)
				return result
			}
		}
		if startOutside {
			result = append(result, p2)
		}
	}

	return result
}

// 欧式距离平方
func dist2(a, b MapPointFloat) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

// 判断线段是否与圆相交，并返回第一个交点
func lineCircleIntersection(p1, p2, center MapPointFloat, r float64) (bool, MapPointFloat) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	fx := p1.X - center.X
	fy := p1.Y - center.Y

	a := dx*dx + dy*dy
	b := 2 * (fx*dx + fy*dy)
	c := fx*fx + fy*fy - r*r

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false, MapPointFloat{}
	}

	discriminant = math.Sqrt(discriminant)
	t1 := (-b - discriminant) / (2 * a)
	t2 := (-b + discriminant) / (2 * a)

	if t1 >= 0 && t1 <= 1 {
		return true, MapPointFloat{p1.X + t1*dx, p1.Y + t1*dy}
	}
	if t2 >= 0 && t2 <= 1 {
		return true, MapPointFloat{p1.X + t2*dx, p1.Y + t2*dy}
	}
	return false, MapPointFloat{}
}
