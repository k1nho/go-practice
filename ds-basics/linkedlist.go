package DS


type Node struct {
	val int
	next *Node
}

type LinkedList struct { 
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Add(x int){
	newNode := new(Node)
	newNode.val = x
	l.size++

	if l.head == nil {
		l.head, l.tail = newNode	
	}
	else{
		l.tail.next = newNode
	}
}

func (l LinkedList) Get(x int) *Node{
	for it := l.head; it != nil; it = it.next{
		if it.val == x {
			return it
		}
	}	
	return nil
}

func (l *LinkedList) Remove(x int){
	var prev *Node
	
	for cur := l.head; cur != nil; cur = cur.next {

		if cur.val == x {
			if l.head == cur {
				l.head = cur.next
				if l.size == 0 {
					l.tail = cur.next
				}
			} else if l.tail == cur{
				l.tail = prev
				prev.next = cur.next
			} else{
				prev.next = cur.next
			}
			return
		}
		prev = cur
	}
	l.size--
}


func (l LinkedList) Size() int {
	return l.size
}

func (l LinkedList) String() string{
	sb := strings.Builder{}

	for it := l.head; it != nil; it = it.next{
		sb.WriteString(fmt.Sprintf("%d ", it.val))
	}

	return sb.String()

}
