package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func Distance(p1 *Point, p2 *Point) float64 {
	dx := math.Abs(p1.x - p2.x)
	dy := math.Abs(p1.y - p2.y)

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := Point{x: 2, y: 10}
	point2 := Point{x: 12, y: -6}
	distance := Distance(&point1, &point2)

	fmt.Println(distance)
}
