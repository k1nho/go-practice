package main

import (
	"fmt"
	"math"
)

type Node struct {
	val  int
	key  int
	prev *Node
	next *Node
}

type LRUCache struct {
	mru       *Node
	lru       *Node
	capacity  int
	keyToNode map[int]*Node
}

func (l *LRUCache) addNode(node *Node) {
	prev := l.mru.prev
	// fix left side
	prev.next = node
	node.prev = prev
	// fix right side
	l.mru.prev = node
	node.next = l.mru
}

func (l *LRUCache) removeNode(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func New(capacity int) *LRUCache {
	mru := &Node{
		val:  0,
		key:  math.MaxInt,
		prev: nil,
		next: nil,
	}
	lru := &Node{
		val:  0,
		key:  math.MaxInt,
		prev: nil,
		next: nil,
	}
	mru.prev = lru
	lru.next = mru

	return &LRUCache{
		mru:       mru,
		lru:       lru,
		capacity:  capacity,
		keyToNode: make(map[int]*Node),
	}
}

// Get: Returns a Node value if it exists otherwise it return -1
func (l *LRUCache) Get(key int) int {
	if _, exist := l.keyToNode[key]; exist {
		node := l.keyToNode[key]
		l.removeNode(l.keyToNode[key])
		l.addNode(node)
		return node.val
	}
	return -1

}

// Put: Updates the value of key if it exists, otherwise it adds the (key, val) pair into the cache
// If it exceeds the capacity it evicts the least recently used key
func (l *LRUCache) Put(key, val int) {

	if _, exists := l.keyToNode[key]; exists {
		l.removeNode(l.keyToNode[key])
	}

	node := &Node{val: val, key: key, prev: nil, next: nil}
	l.keyToNode[key] = node
	l.addNode(node)

	if len(l.keyToNode) > l.capacity {
		lruNode := l.lru.next
		l.removeNode(lruNode)
		delete(l.keyToNode, key)
	}
}

func (l *LRUCache) Display() {
	start := l.lru.next

	for start.next != nil {
		fmt.Printf("%d ", start.val)
		start = start.next
	}
	fmt.Println()
}

func main() {

	lru := New(3)
	lru.Put(1, 10)
	lru.Put(2, 150)
	lru.Put(3, 20)
	lru.Display()
	lru.Get(2)
	lru.Display()
	lru.Get(1)
	lru.Display()
	lru.Put(5, 300)
	lru.Display()

}
