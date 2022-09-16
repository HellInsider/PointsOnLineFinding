package main

import (
	"fmt"
)

type Point struct {
	x, y, ind int
}

type kCoef struct { //y = kx+b
	a, b int // a/b
}

type bCoef struct {
	a, b int // a/b
}

func main() {
	findMPointsOnLine(3, getCoefs(getPoints()))
}

func IntAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func GCD(a, b int) int {
	for b != 0 {
		a %= b
		if a == 0 {
			return IntAbs(b)
		}
		b %= a
	}
	return IntAbs(a)
}

func getPoints() (int, []Point) {
	var n, x, y int
	var points []Point
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		points = append(points, Point{x: x, y: y, ind: i + 1})
	}
	return n, points
}

// I made third map to implement set for points instead of array
func getCoefs(n int, points []Point) map[kCoef]map[bCoef]map[Point]struct{} {
	var dX, dY, gcd int
	var k kCoef
	var b bCoef
	coefMap := make(map[kCoef]map[bCoef]map[Point]struct{}) //map(k, map(b, map(Points, void)))

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dX = points[i].x - points[j].x
			dY = points[i].y - points[j].y
			if dY < 0 {
				dY *= -1
				dX *= -1
			}

			if dX == 0 {
				k = kCoef{1, 0} // if points are vertical, I write (1, 0) as k coef and their x offset to b coef
				b = bCoef{points[i].x, 1}
			} else if dY == 0 {
				k = kCoef{0, 1} // if points are horizontal, I write (0, 1) as k coef and their y offset to b coef
				b = bCoef{points[i].y, 1}
			} else {
				gcd = GCD(dY, IntAbs(dX))
				k = kCoef{dY / gcd, dX / gcd}

				t := points[i].y*k.b - k.a*points[i].x
				gcd = GCD(t, k.b)
				b = bCoef{t / gcd, k.b / gcd}
			}

			//fmt.Println("(", points[i].x, ", ", points[i].y, ") -> ", "(", points[j].x, ", ", points[j].y, "): ", k, b)

			if coefMap[k] == nil {
				coefMap[k] = make(map[bCoef]map[Point]struct{})
			}

			if coefMap[k][b] == nil {
				coefMap[k][b] = make(map[Point]struct{})
				coefMap[k][b][points[i]] = struct{}{}
			}

			coefMap[k][b][points[j]] = struct{}{}

			/*fmt.Print(k, b, ": ")
			for i, _ := range coefMap[k][b] {
				fmt.Print(i.ind, " ")
			}
			fmt.Println()
			*/
		}
	}
	return coefMap
}

func findMPointsOnLine(m int, coefMap map[kCoef]map[bCoef]map[Point]struct{}) {

	//fmt.Println(coefMap)

	for _, i := range coefMap {
		for _, j := range i {
			if len(j) >= m {
				//fmt.Print(k, b, ": ")
				for q, _ := range j {
					fmt.Print(q.ind, " ")
				}
				fmt.Print("\n")
			}
		}
	}
}
