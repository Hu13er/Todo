package Todo

type Node struct {
	next *Node
	prev *Node

	Value interface{}
}

func NewNode(value interface{}) *Node {
	return &Node{next: nil, prev: nil, Value:value}
}


func (this *Node) Next() *Node{
	return this.next
}

func (this *Node) Prev() *Node{
	return this.prev
}

func (this *Node) Last() *Node{
	now := this
	for now.Next() != nil {
		now = now.Next()
	}
	return now
}

func (this *Node) First() *Node{
	now := this
	for now.Prev() != nil {
		now = now.Prev()
	}
	return now
}


func (this *Node) Push(node *Node) *Node{
	last := this.Last()
	last.next = node
	node.prev = last

	return last
}

func (this *Node) Pop() (oldFirst *Node, newFirst *Node) {
	oldFirst = this.First()

	newFirst = oldFirst.Next()
	oldFirst.next = nil
	newFirst.prev = nil

	return
}

func (this *Node) Size() int {
	now := this.First()

	outp := 1
	for now.Next() != nil {
		outp++
		now = now.Next()
	}

	return outp
}

type mapFunc func(interface{}) interface{}

func (this *Node) Map(f mapFunc) *Node{

	now := this.First()

	for {
		if now == nil {
			break
		}

		now.Value = f(now.Value)
		now = now.Next()
	}

	return this
}


