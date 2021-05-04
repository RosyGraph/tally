// Package stack provides a simple string stack for parsing variables and operators
package stack

type StringStack struct {
	data string
	next *StringStack
}

func (stack StringStack) Add(s string) StringStack {
	return StringStack{s, &stack}
}

func (stack StringStack) Pop() (StringStack, string) {
	return *stack.next, stack.data
}

type FloatStack struct {
	data float32
	next *FloatStack
}

func (stack FloatStack) Add(n float32) FloatStack {
	return FloatStack{n, &stack}
}

func (stack FloatStack) Pop() (FloatStack, float32) {
	return *stack.next, stack.data
}
