package main

import (
	"fmt"
	"math"
)

type ShapeAreaRes interface {
  int | float64
}

type Shape interface {
  Area()
}


type Rectangle struct {
  width int;
  height int;
}

type Triangle struct {
  A, B, C int
}

// implement interface
func (r Rectangle) Area() {
  fmt.Printf( "area %d",  r.width*r.height)
}


func (t Triangle) Area() {
  s := (t.A + t.B + t.C)
  
  fmt.Printf("area: %f", math.Sqrt(float64(s*(s-t.A)*(s-t.B)*(s-t.C))))
}

func checkType(shape Shape) {
  fmt.Printf("Type of this is %T\n", shape)
}

func main() {
  var shape Shape

  shape = Triangle{1,2,3}
  shape.Area()
  checkType(shape)
  shape= Rectangle{10,4}
  shape.Area()
  checkType(shape)
}
