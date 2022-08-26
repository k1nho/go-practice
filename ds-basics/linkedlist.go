package DS


type Node struct {
	val int
	next *Node
}

type LinkedList struct { 
	head *Node
	size int
}

func (l *LinkedList) add(x int){
	newNode := new(Node)
	newNode.val = x

	if l.head == nil {
		l.head = newNode	
	}
	else{
		it := l.head
		for ; it.next != nil; it = it.next{
		}
		it.next = newNode
	}
}

func (l *LinkedList) get(x int) *Node{
	for it := l.head; it.next != nil; it = it.next{
		if it.val == x {
			return it
		}
	}	
	return nil
}

func (l LinkedList) remove(x int){

}


func (l LinkedList) String() string{
	sb := strings.Builder{}

	for it := l.head; it.next != nil; it = it.next{
		sb.WriteString(fmt.Sprintf("%d ", it.val))
	}

	return sb.String()

}


func main() {
	LL := LinkedList{}
	LL.add(1)
	LL.add(2)
	LL.add(3)
	LL.add(4)

}
