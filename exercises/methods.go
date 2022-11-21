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

func (p *Point) IncreaseSlope(x int) {
  p.x += x
  p.y += x
}

func main() {
  p := Point{4,2}
  p1 := Point{1, 1}
  fmt.Println(p.Slope(p1))
  p1.IncreaseSlope(1)
  fmt.Printf("x: %d, y: %d\n", p1.x, p1.y)
}
