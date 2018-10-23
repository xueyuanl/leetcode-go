package stack

import (
	"strconv"
	. "data-structure"
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