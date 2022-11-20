package main

import "fmt"


type Constraint interface {
 int | bool
}

type List[T Constraint] struct {
  next *List[T]
  val T
}


func (head *List[T]) Add(node List[T]) {
  // traverse till the end of the LL
  v := head
  for v.next != nil {
    v = v.next
  }
  v.next = &node;
}

func (head List[T]) Print() {
  for head.next != nil {
    fmt.Printf("%v ", head.val)
    head = *head.next
  }
  fmt.Printf("%v\n", head.val)
}

func main() {
  linkedList := List[int]{nil, 0}
  linkedList.Add(List[int]{nil, 5})
  linkedList.Add(List[int]{nil, 7})
  linkedList.Add(List[int]{nil, 8})

  linkedList.Print()

  // The following will fail since we contraint our generic to be only int or bool
  //stringList := List[string]{nil, "5"}

  boolList := List[bool]{nil, true}
  boolList.Add(List[bool]{nil, false})
  boolList.Add(List[bool]{nil, true})
  boolList.Add(List[bool]{nil, true})
  boolList.Add(List[bool]{nil, true})
  boolList.Add(List[bool]{nil, false})

  boolList.Print()
}
