package lru

import (
	"github.com/zdq0394/algorithm/others/cache"
)

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	cap  int
	head *Node
	tail *Node
	data map[int]*Node
}

var _ cache.Cache = &LRUCache{}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		head: nil,
		tail: nil,
		data: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.Hit(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.value = value
		this.Hit(node)
	} else {
		if this.IsFull() {
			this.Eliminate()
		}
		this.AddNode(&Node{key, value, nil, nil})
	}
}

func (this *LRUCache) IsFull() bool {
	return len(this.data) == this.cap
}

func (this *LRUCache) Eliminate() {
	if len(this.data) == 0 {
		return
	}
	if len(this.data) == 1 {
		this.head = nil
		this.tail = nil
		this.data = make(map[int]*Node, this.cap)
		return
	}
	curNode := this.tail
	this.tail = curNode.pre
	this.tail.next = nil
	curNode.pre = nil
	delete(this.data, curNode.key)
}

func (this *LRUCache) AddNode(node *Node) {
	node.next = this.head
	if this.head != nil {
		this.head.pre = node
	} else {
		this.tail = node
	}
	this.head = node
	this.data[node.key] = node
}

func (this *LRUCache) Hit(node *Node) {
	if this.head == node {
		// the visited node is head, no need to move.
		return
	}
	node1 := node.pre
	node2 := node.next
	node1.next = node2
	if node2 != nil {
		node2.pre = node1
	} else {
		this.tail = node1
	}
	node.next = this.head
	this.head.pre = node
	node.pre = nil
	this.head = node
}
