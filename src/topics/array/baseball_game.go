package array

import (
	"strconv"
)

func calPoints(ops []string) int {

	stack := new(Stack)

	for _,v := range ops {
		switch v {
			case "C":
				stack.Pop()
			case "D":
				value := stack.Top().(*Element).Value.(int) * 2
				stack.Push(value)
			case "+":
				value := stack.Top().(*Element).Value.(int) + stack.Top().(*Element).Before.Value.(int)
				stack.Push(value)
			default:
				i,err := strconv.Atoi(v)
				if err == nil {
					stack.Push(i)
				}
		}
	}
	res := 0
	for stack.Len() > 0 {
		res += stack.Pop().(int)
	}
	return res
}

type Stack struct {
	top *Element

	size int
}

type Element struct {
	Value  interface{}
	Before *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Top() interface{} {
	if s.size > 0 {
		return s.top
	}
	return nil
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size ++
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.Value
		s.top = s.top.Before
		s.size --
		return value
	}
	return nil
}