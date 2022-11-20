package main

import "fmt"


//type Constraint interface {
  //int | bool
//}

type List[T any] struct {
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

}
