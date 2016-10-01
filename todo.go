package Todo

import (
	"sync"
)

type nothing struct{}

var pulse = nothing{}

type Todo struct {
	firstNode *Node
	lastNode  *Node

	mutex *sync.Mutex
	pulse chan nothing
	stop  chan nothing

	do func(interface{})
}

func NewTodo(f func(interface{})) *Todo {
	return &Todo{firstNode: nil, lastNode: nil, mutex: &sync.Mutex{}, pulse: make(chan nothing), stop: make(chan nothing), do: f}
}

func (this *Todo) IsEmpty() bool {
	if this.firstNode == nil && this.lastNode == nil {
		return true
	} else {
		return false
	}
}

func (this *Todo) Push(something interface{}) {

	defer this.beat()
	defer this.mutex.Unlock()
	this.mutex.Lock()

	node := newNode(something)
	if this.firstNode == nil && this.lastNode == nil {
		this.firstNode = node
		this.lastNode = node
		return
	}
	this.lastNode = this.lastNode.push(node)
}

func (this *Todo) beat() {
	select {
	case this.pulse <- pulse:
	default:
	}
}

func (this *Todo) Run() {
	locked := true
	this.mutex.Lock()
	if !this.IsEmpty() {
		go func() {
			this.pulse <- pulse
		}()
	}
	go func() {
	outer:
		for {
			if locked {
				this.mutex.Unlock()
				locked = false
			}
			select {
			case <-this.pulse:
				for !this.IsEmpty() {
					select {
					case <-this.stop:
						break outer
					default:
					}
					now, newFirst := this.firstNode.pop()
					this.firstNode = newFirst
					if this.firstNode == nil {
						this.lastNode = nil
					}
					this.do(now.Value)
				}
			case <-this.stop:
				break
			}
		}
	}()
}

func (this *Todo) Stop() {
	this.stop <- pulse
}
