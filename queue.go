package Todo

type node struct {
	next *node
	prev *node

	value interface{}
}

func newNode(value interface{}) *Node {
	return &node{next: nil, prev: nil, value: value}
}

func (this *node) next() *node {
	return this.next
}

func (this *node) prev() *node {
	return this.prev
}

func (this *node) last() *node {
	now := this
	for now.next() != nil {
		now = now.next()
	}
	return now
}

func (this *node) first() *node {
	now := this
	for now.prev() != nil {
		now = now.prev()
	}
	return now
}

func (this *node) push(node *node) *node {
	last := this.last()
	last.next = node
	node.prev = last
	return last
}

func (this *node) pop() (*node, *node) {
	oldFirst := this.first()
	newFirst := oldFirst.next()
	if newFirst != nil {
		newFirst.prev = nil
	}
	oldFirst.next = nil
	return oldFirst, newFirst
}

func (this *node) size() int {
	now := this.first()
	outp := 1
	for now.next() != nil {
		outp++
		now = now.next()
	}
	return outp
}
