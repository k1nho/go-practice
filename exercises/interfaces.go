package main

import (
	"fmt"
	"math"
  "strconv"
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

// implemeting print for rectangle
func (r Rectangle) String() string {
  return fmt.Sprintf("Rectangle: [width: %d, height: %d]", r.width, r.height)
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


// Stringers ex
type IPAddr [4]byte

func (ip IPAddr) String() string {
  var s string
  for _ , val := range ip {
    s += strconv.Itoa(int(val)) + "."
  }

  return s[:len(s)-1]
}


func main() {
  var shape Shape

  shape = Triangle{1,2,3}
  shape.Area()
  checkType(shape)
  shape= Rectangle{10,4}
  shape.Area()
  checkType(shape)

  // the empty interface specifies 0 methods is known as the empty interface (it may hold any type)

  // Stringer interface is used to implement print of values to strings
  rect := Rectangle{10,40}
  fmt.Println(rect)

  hosts := map[string]IPAddr{
    "loopback": {127,0,0,1},
    "googleDNS" :{8,8,8,8},
  }

  for k,v := range hosts {
    fmt.Printf("%v, %v\n", k, v)
  }
}
