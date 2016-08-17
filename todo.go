package Todo

type nothing struct {}
var pulse = nothing{}

type Todo struct {
	firstNode*Node
	lastNode *Node

	pulse chan nothing
	stop chan nothing

	do func(interface{})
}


func NewTodo(f func(interface{})) *Todo{
	return &Todo{firstNode: nil, lastNode: nil, pulse: make(chan nothing), stop: make(chan nothing), do: f}
}

func (this *Todo) IsEmpty() bool {
	if this.firstNode == nil && this.lastNode == nil {
		return true
	} else {
		return false
	}
}

func (this *Todo) Push(something interface{}){

	defer this.beat()

	node := NewNode(something)
	if this.firstNode == nil && this.lastNode == nil {
		this.firstNode = node
		this.lastNode = node
		return
	}

	this.lastNode = this.lastNode.Push(node)
}

func (this *Todo) beat(){
	select {
	case this.pulse <- pulse:
		println("beated")
	default:
		println("default")
	}
}

func (this *Todo) Run(){
	if !this.IsEmpty() {
		go func(){
			this.pulse <- pulse
		}()
	}
	go func() {
		outer: for {
			println("w8ing")
			select {
			case <- this.pulse:
				for !this.IsEmpty() {
					select {
					case <- this.stop:
						break outer
					default:
					}

					now, newFirst := this.firstNode.Pop()
					this.firstNode = newFirst
					if this.firstNode == nil {
						this.lastNode = nil
						// this.Size == 0
					}

					this.do(now.Value)
				}
			case <- this.stop:
				break
			}

		}
	}()
}

func (this *Todo) Stop(){
	this.stop <- pulse
}