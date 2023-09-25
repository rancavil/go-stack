package stack

type Stack []interface{}

func (stack *Stack) Pop() {
	*stack = (*stack)[:len(*stack)-1]
}

func (stack *Stack) Push(elem interface{}) {
	*stack = append(*stack, elem)
}
