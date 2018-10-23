package stack


func dailyTemperatures(T []int) []int {

	stack := new(Stack)

	res := make([]int, len(T))

	for i,v := range T {

		for stack.Len()	> 0 && T[stack.Top().(*Element).Value.(int)] < v {
			res[stack.Top().(*Element).Value.(int)] = i - stack.Top().(*Element).Value.(int)
			stack.Pop()
		}
		stack.Push(i)
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
		ele := s.top
		s.top = s.top.Before
		s.size --
		return ele
	}
	return nil
}