package Todo

type node struct {
	next *node
	prev *node

	value interface{}
}

func newNode(value interface{}) *node {
	return &node{next: nil, prev: nil, value: value}
}

func (this *node) nextNode() *node {
	return this.next
}

func (this *node) prevNode() *node {
	return this.prev
}

func (this *node) lastNode() *node {
	now := this
	for now.nextNode() != nil {
		now = now.nextNode()
	}
	return now
}

func (this *node) firstNode() *node {
	now := this
	for now.prevNode() != nil {
		now = now.prevNode()
	}
	return now
}

func (this *node) push(node *node) *node {
	last := this.lastNode()
	last.next = node
	node.prev = last
	return last
}

func (this *node) pop() (*node, *node) {
	oldFirst := this.firstNode()
	newFirst := oldFirst.nextNode()
	if newFirst != nil {
		newFirst.prev = nil
	}
	oldFirst.next = nil
	return oldFirst, newFirst
}

func (this *node) size() int {
	now := this.firstNode()
	outp := 1
	for now.nextNode() != nil {
		outp++
		now = now.nextNode()
	}
	return outp
}
