package tree

import (
	. "basic"
	"reflect"
)

func averageOfLevels(root *TreeNode) []float64 {

	var average []float64
	que := new(queue)
	que.add(root)

	var sum int
	var num int
	next := new(queue)
	for que.size != 0 {
		peek:= que.peek().(*TreeNode)
		sum += peek.Val
		num ++
		next.add(peek.Left)
		next.add(peek.Right)
		que.remove()

		if que.size == 0 {
			var levelSum float64 = float64(sum) / float64(num)
			average = append(average, levelSum)
			if next.size != 0 {
				que = next
				next = new(queue)
				sum = 0
				num = 0
			}
		}
	}
	return average
}

//implement a queue structure
type node struct {
	element interface{}
	prev    *node
	next    *node
}

type queue struct {
	size int
	head *node
	tail *node
}

func (this *queue) peek() interface{} {
	if this.head == nil {
		panic("Empty queue.")
	}
	return this.head.element
}

func (this *queue) remove() {
	if this.head == nil {
		panic("Empty queue.")
	}
	this.head = this.head.next
	this.size --
}

func (this *queue) add(element interface{}) {
	if element == nil || reflect.ValueOf(element).IsNil()  {
		return
	}
	newNode := &node{element, this.tail, nil}

	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	}else {
		this.tail.next = newNode
		this.tail = newNode
	}

	this.size ++
}

func (this *queue) getSize() int {
	return this.size
}
