package main

import "fmt"


type Point struct {
  x int
  y int
}

func (p Point) Slope(p1 Point) float64 {
  num := p.x - p1.x
  dem := p.y - p1.y
  if dem == 0 {return 0.0}
  return float64(num/dem)
}

func main() {
  p := Point{4,2}
  p1 := Point{1, 1}
  fmt.Println(p.Slope(p1))
}
