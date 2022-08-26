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

}

func (l LinkedList) remove(x int){

}


func (l LinkedList) String() string{


}


func main() [

]