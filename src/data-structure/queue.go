package data_structure

import "reflect"

type node struct {
	element interface{}
	prev    *node
	next    *node
}

type Queue struct {
	size int
	head *node
	tail *node
}

func (this *Queue) Peek() interface{} {
	if this.head == nil {
		panic("Empty queue.")
	}
	return this.head.element
}

func (this *Queue) Remove() {
	if this.head == nil {
		panic("Empty queue.")
	}
	this.head = this.head.next
	this.size --
}

func (this *Queue) Add(element interface{}) {
	if element == nil || reflect.ValueOf(element).IsNil() {
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

func (this *Queue) Size() int {
	return this.size
}
