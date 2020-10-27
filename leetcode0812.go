package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	var ans float64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				ans = math.Max(ans,getArea(points[i],points[j],points[k]))
			}
		}
	}
	return ans
}
func getArea(a,b,c []int) float64 {
	return 0.5 * math.Abs(float64(a[0]*b[1] + b[0]*c[1] + c[0]*a[1]-a[1]*b[0] - b[1]*c[0] - c[1]*a[0]))
}

func main() {
	points := [][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}
	fmt.Println(largestTriangleArea(points))
}