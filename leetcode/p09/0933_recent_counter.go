package p09

type Node struct {
	t    int
	next *Node
}
type RecentCounter struct {
	head *Node
	tail *Node
	size int
}

func Constructor() RecentCounter {
	return RecentCounter{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.size++
	n := &Node{
		t:    t,
		next: nil,
	}
	if this.tail == nil {
		this.tail = n
	} else {
		this.tail.next = n
		this.tail = n
	}
	if this.head == nil {
		this.head = n
	} else {
		p := this.head
		for p != this.tail {
			if p.t+3000 < t {
				p = p.next
				this.size--
			} else {
				break
			}
		}
		this.head = p
	}
	return this.size
}
